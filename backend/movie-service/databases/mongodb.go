package databases

import (
	"time"

	"github.com/bangbugs/go-microservice-1/movie-service/common"

	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

type MongoDB struct {
	MongoDbSession *mgo.Session
	Databasename   string
}

func (db *MongoDB) Init() error {
	db.Databasename = common.Config.MongoDbName

	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{common.Config.MongoAddrs},
		Timeout:  time.Second * 60,
		Database: db.Databasename,
		Username: common.Config.MongoDbUsername,
		Password: common.Config.MongoDbPassword,
	}

	// Create a session which maintains a pool of socket connections to the mongodb database
	var err error
	db.MongoDbSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Debug("Can't connect to mongo, go error: ", err)
		return err
	}

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MongoDbSession != nil {
		db.MongoDbSession.Close()
	}
}
