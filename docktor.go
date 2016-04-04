package docktor

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/soprasteria/godocktor-api/common/logs"
	"github.com/soprasteria/godocktor-api/daemons"
	"github.com/soprasteria/godocktor-api/groups"
	"github.com/soprasteria/godocktor-api/services"
	"github.com/soprasteria/godocktor-api/sites"
	"github.com/soprasteria/godocktor-api/users"
)

func init() {
	initLoggers()
}

func initLoggers() {
	var debugLogger io.Writer
	var flag int
	debugLogger = ioutil.Discard
	flag = log.Ldate | log.Ltime

	logs.InitLog(debugLogger, os.Stdout, os.Stdout, os.Stderr, flag)
}

// Open the connexion to docktor API
func Open(docktorMongoHost string) (*Docktor, error) {

	session, err := mgo.Dial(docktorMongoHost)
	if err != nil {
		return &Docktor{}, err
	}
	session.SetMode(mgo.Monotonic, true)
	context := appContext{session.DB("docktor")}
	services := &services.Repo{Coll: context.db.C("services")}
	groups := &groups.Repo{Coll: context.db.C("groups")}
	daemons := &daemons.Repo{Coll: context.db.C("daemons")}
	users := &users.Repo{Coll: context.db.C("users")}
	sites := &sites.Repo{Coll: context.db.C("sites")}

	return &Docktor{
		services: services,
		groups:   groups,
		daemons:  daemons,
		users:    users,
		sites:    sites,
		session:  session,
	}, nil
}

// Close the connexion to docktor API
func (dock *Docktor) Close() {
	dock.session.Close()
}

// Services is the entrypoint for Services API
func (dock *Docktor) Services() services.RepoServices {
	return dock.services
}

// Groups is the entrypoint for Groups API
func (dock *Docktor) Groups() groups.RepoGroups {
	return dock.groups
}

// Daemons is the entrypoint for Daemons API
func (dock *Docktor) Daemons() daemons.RepoDaemons {
	return dock.daemons
}

// Users is the entrypoint for Users API
func (dock *Docktor) Users() users.RepoUsers {
	return dock.users
}

// Sites is the entrypoint for Sites API
func (dock *Docktor) Sites() sites.RepoSites {
	return dock.sites
}
