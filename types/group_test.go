package types

import (
	"fmt"
	"testing"
)

func ExampleGroupVolumeFormat() {
	vc := VolumeContainer{
		External: "/brace/yourselves",
		Internal: "/winter/is/coming",
		Rights:   "ro",
	}
	fmt.Println(vc.Format())

	vc = VolumeContainer{
		External: "/brace/yourselves",
		Internal: "/winter/is/coming",
	}
	fmt.Println(vc.Format())

	//Output:
	// /brace/yourselves:/winter/is/coming:ro
	// /brace/yourselves:/winter/is/coming:rw
}

func TestPortsRangeIsNegativeShouldFail(t *testing.T) {
	portRange := PortsRange{0, 0}
	if _, err := portRange.Check(); err == nil {
		t.Fatalf("Port range is negative (%+v). It's not possible", portRange)
	}
}

func TestPortsRangeIsInversedShouldFail(t *testing.T) {
	portRange := PortsRange{100, 50}
	if _, err := portRange.Check(); err == nil {
		t.Fatalf("Max is less than min (%+v). It's not possible", portRange)
	}
}

func TestPortsRangeDefaultShouldFail(t *testing.T) {
	portRange := PortsRange{}
	if _, err := portRange.Check(); err == nil {
		t.Fatalf("Default port range %v is not valid", portRange)
	}
}

func TestPortsRangeConsecutiveAndExclusiveShouldBeOK(t *testing.T) {
	portRange := PortsRange{100, 200}
	if _, err := portRange.Check(); err != nil {
		t.Fatalf("Port range %+v should be valid", portRange)
	}
}

func TestGetsPortRange(t *testing.T) {
	g := Group{
		PortMinRange: 10000,
		PortMaxRange: 11000,
	}
	expected := PortsRange{
		Min: 10000,
		Max: 11000,
	}
	r, err := g.GetPortsRange()
	if err != nil {
		t.Fatal(err)
	}

	if r.Min != expected.Min || r.Max != expected.Max {
		t.Fatalf("expected %+v, obtained %+v", expected, r)
	}
}

func TestGetsPortRangeOutOfRange(t *testing.T) {
	g := Group{
		PortMinRange: 10000,
		PortMaxRange: 66666,
	}
	r, err := g.GetPortsRange()
	if err == nil {
		t.Fatalf("Port max range is out of 65535. Should be KO. Obtainer %v", r)
	}
}

func TestPortsRangeMerge(t *testing.T) {
	portRanges := []PortsRange{
		{10, 20},
		{15, 18},
		{30, 40},
		{5, 22},
		{21, 25},
	}

	obtainedPortsRange := mergeRanges(portRanges)
	expectedPortsRange := []PortsRange{
		{5, 25},
		{30, 40},
	}

	for i, v := range obtainedPortsRange {
		if v != expectedPortsRange[i] {
			t.Fatalf("expected %+v, obtained %+v", expectedPortsRange[i], v)
		}
	}

}

func TestPortsAvailableShouldFindARangeAtBegining(t *testing.T) {
	usedPortRange := []PortsRange{
		{12000, 20000},
		{15000, 18000},
		{30000, 40000},
		{11000, 22000},
	}

	availablePortRange, err := FindAvailablePortRange(1000, usedPortRange)
	if err != nil {
		t.Fatal(err)
	}
	expectedPortRange := PortsRange{9000, 9999}
	if availablePortRange.Min != expectedPortRange.Min || availablePortRange.Max != expectedPortRange.Max {
		t.Fatalf("expected %+v, obtained %+v", expectedPortRange, availablePortRange)
	}
}

func TestPortsAvailableShouldFindARangeInBetween(t *testing.T) {
	usedPortRange := []PortsRange{
		{8000, 20000},
		{15000, 18000},
		{30000, 40000},
		{11000, 22000},
	}

	availablePortRange, err := FindAvailablePortRange(1000, usedPortRange)
	if err != nil {
		t.Fatal(err)
	}
	expectedPortRange := PortsRange{22001, 23000}
	if availablePortRange.Min != expectedPortRange.Min || availablePortRange.Max != expectedPortRange.Max {
		t.Fatalf("expected %+v, obtained %+v", expectedPortRange, availablePortRange)
	}
}

func TestPortsAvailableShouldFindARangeAtEnd(t *testing.T) {
	usedPortRange := []PortsRange{
		{8000, 20000},
		{15000, 18000},
		{20000, 40000},
		{11000, 22000},
	}

	availablePortRange, err := FindAvailablePortRange(1000, usedPortRange)
	if err != nil {
		t.Fatal(err)
	}
	expectedPortRange := PortsRange{40001, 41000}
	if availablePortRange.Min != expectedPortRange.Min || availablePortRange.Max != expectedPortRange.Max {
		t.Fatalf("expected %+v, obtained %+v", expectedPortRange, availablePortRange)
	}
}

func TestPortsAvailableShouldFindARangeAtEndBecauseNotEnougSpaceInBetween(t *testing.T) {
	usedPortRange := []PortsRange{
		{8000, 20000},
		{15000, 18000},
		{22500, 40000},
		{11000, 22000},
	}

	availablePortRange, err := FindAvailablePortRange(1000, usedPortRange)
	if err != nil {
		t.Fatal(err)
	}
	expectedPortRange := PortsRange{40001, 41000}
	if availablePortRange.Min != expectedPortRange.Min || availablePortRange.Max != expectedPortRange.Max {
		t.Fatalf("expected %+v, obtained %+v", expectedPortRange, availablePortRange)
	}
}

func TestPortsAvailableShouldNotFindAnyRangeBecauseNotEnoughSpace(t *testing.T) {
	usedPortRange := []PortsRange{
		{8000, 20000},
		{22000, 65535},
	}

	a, err := FindAvailablePortRange(2001, usedPortRange)
	if err == nil {
		t.Fatalf("Sould not found an available port range but found %v", a)
	}

}

func TestPortsAvailableShouldFindARangeOfOneAtTheEnd(t *testing.T) {
	usedPortRange := []PortsRange{{8000, 65534}}

	availablePortRange, err := FindAvailablePortRange(1, usedPortRange)
	if err != nil {
		t.Fatal(err)
	}
	expectedPortRange := PortsRange{65535, 65535}
	if availablePortRange.Min != expectedPortRange.Min || availablePortRange.Max != expectedPortRange.Max {
		t.Fatalf("expected %+v, obtained %+v", expectedPortRange, availablePortRange)
	}
}
