package types

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"time"

	maths "github.com/pkg/math"
	"gopkg.in/mgo.v2/bson"
)

// FileSystem is a filesystem watched by the group
type FileSystem struct {
	ID          bson.ObjectId `bson:"_id,omitempty" yaml:"id,omitempty"`
	Daemon      string        `bson:"daemon" yaml:"daemon,omitempty"`
	Partition   string        `bson:"partition,omitempty" yaml:"partition,omitempty"`
	Description string        `bson:"description" yaml:"description,omitempty"`
}

//FileSystems is a slice of FileSystem
type FileSystems []FileSystem

// Group is a entity (like a project) that gather services instances as containers
type Group struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Created      time.Time     `bson:"created"`
	Title        string        `bson:"title"`
	Description  string        `bson:"description"`
	PortMinRange int           `bson:"portminrange"`
	PortMaxRange int           `bson:"portmaxrange"`
	Daemon       bson.ObjectId `bson:"daemon"`
	FileSystems  FileSystems   `bson:"filesystems"`
	Containers   Containers    `bson:"containers"`
	User         bson.ObjectId `bson:"variables"`
}

// AddFileSystem adds a FileSystem to the Group
func (g *Group) AddFileSystem(f FileSystem) {
	g.FileSystems = append(g.FileSystems, f)
}

// AddContainer adds a Container to the Group
func (g *Group) AddContainer(c Container) {
	g.Containers = append(g.Containers, c)
}

// ContainerWithGroup is a entity which contains a container, linked to a group
type ContainerWithGroup struct {
	Group     Group
	Container Container
}

// ContainerWithGroupID is an entity which contains a container, linked to a group ID
type ContainerWithGroupID struct {
	Container Container     `bson:"container"`
	ID        bson.ObjectId `bson:"_id,omitempty"`
}

var reservedPortsRange = PortsRange{0, 8999}

// MinPort port is the minimum possible port
const MinPort int = 1

// PortsRange is a range of ports, from min to max, respectively between 1 and 65535
// Max is inclusive
type PortsRange struct {
	Min int `yaml:"min" bson:"min"`
	Max int `yaml:"max" bson:"max"`
}

// Check checks that the port range is a valid port range. Return false and an error if not valid, true and nil otherwise
// A port range is made of two int16 so : 1 <= min < max <= 65535
func (p PortsRange) Check() (bool, error) {
	if p.Min < MinPort || p.Min > p.Max {
		return false, fmt.Errorf("Expected port range is made of two number between 1 and %v. Actual: Min=%v, Max=%v", math.MaxUint16, p.Min, p.Max)
	}
	return true, nil
}

// GetPortsRange get the range of port given for a group
// Checks values because types are changed
func (g *Group) GetPortsRange() (PortsRange, error) {
	if g.PortMinRange < MinPort {
		return PortsRange{}, fmt.Errorf("A port is between %v and %v. Obtained: %v", MinPort, math.MaxUint16, g.PortMinRange)
	}
	if g.PortMaxRange > math.MaxUint16 {
		return PortsRange{}, fmt.Errorf("A port is between %v and %v. Obtained: %v", MinPort, math.MaxUint16, g.PortMaxRange)
	}
	return PortsRange{g.PortMinRange, g.PortMaxRange}, nil
}

// FindAvailablePortRange finds available range of consecutive <numberOfPorts> ports, not in usedPortsRange range.
func FindAvailablePortRange(numberOfPorts int, usedPortsRanges []PortsRange) (PortsRange, error) {

	if numberOfPorts <= 0 {
		return PortsRange{}, fmt.Errorf("<numberOfPorts> (%v) should be positive", numberOfPorts)
	}

	// Add the range of reserved ports for the system and usual programs
	usedPortsRanges = append(usedPortsRanges, reservedPortsRange)
	// Merge all ranges and sort them by min port
	usedPortsRanges = mergeRanges(usedPortsRanges)

	// Find first range of <numberOfPorts> between used port ranges
	if len(usedPortsRanges) == 0 {
		return PortsRange{}, errors.New("Used port range is empty and should not be")
	}

	precedent := usedPortsRanges[0]
	for i := 0; i < len(usedPortsRanges); i++ {
		current := usedPortsRanges[i]
		if precedent.Max+numberOfPorts < current.Min {
			return PortsRange{precedent.Max + 1, precedent.Max + numberOfPorts}, nil
		}
		precedent = current
	}

	if precedent.Max <= math.MaxUint16-numberOfPorts {
		// If there is room between the last used port and the max
		return PortsRange{precedent.Max + 1, precedent.Max + numberOfPorts}, nil
	}

	return PortsRange{}, fmt.Errorf("Can't find any port range of %v consecutive ports available ", numberOfPorts)
}

// ByMinIncreasing is the type resolving Golang Sorting interfaces
// Aim is to sort by the min port of the range, in the increasing order
type ByMinIncreasing []PortsRange

func (a ByMinIncreasing) Len() int           { return len(a) }
func (a ByMinIncreasing) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMinIncreasing) Less(i, j int) bool { return a[i].Min < a[j].Min }

func mergeRanges(portsRanges []PortsRange) []PortsRange {

	if len(portsRanges) == 0 || len(portsRanges) == 1 {
		return portsRanges
	}

	var result []PortsRange

	sort.Sort(ByMinIncreasing(portsRanges))

	first := portsRanges[0]
	min := first.Min
	max := first.Max

	for i := 1; i < len(portsRanges); i++ {
		current := portsRanges[i]
		if current.Min <= max {
			max = maths.MaxInt(current.Max, max)
		} else {
			result = append(result, PortsRange{min, max})
			min = current.Min
			max = current.Max
		}
	}

	result = append(result, PortsRange{min, max})

	return result

}
