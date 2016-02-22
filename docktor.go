package docktor

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"gitlab.cdk.corp.sopra/cdk/godocktor-api/common/logs"
	"gitlab.cdk.corp.sopra/cdk/godocktor-api/daemons"
	"gitlab.cdk.corp.sopra/cdk/godocktor-api/groups"
	"gitlab.cdk.corp.sopra/cdk/godocktor-api/services"
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

	return &Docktor{
		services: services,
		groups:   groups,
		daemons:  daemons,
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
