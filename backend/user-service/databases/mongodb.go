package databases

import (
	"time"

	"github.com/bangbugs/go-microservice-1/user-service/common"
	"github.com/bangbugs/go-microservice-1/user-service/models"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	MongoDbSession *mgo.Session
	Databasename   string
}

func (db *MongoDB) Init() error {
	db.Databasename = common.Config.MongoDbName

	// Dialinfo holds options for establishing a session with a MongoDB cluster
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{common.Config.MongoAddrs},
		Timeout:  time.Second * 60,
		Database: db.Databasename,
		Username: common.Config.MongoDbUsername,
		Password: common.Config.MongoDbPassword,
	}

	// Create a session which maintains a pool of socket conncections to the mongodb
	var err error
	db.MongoDbSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Debug("Can't connect to mongo, go error: ", err)
		return err
	}

	return db.initData()
}

func (db *MongoDB) initData() error {
	var err error
	var count int

	// Check if user colletion hat at least one document
	sessionCopy := db.MongoDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against
	collection := sessionCopy.DB(db.Databasename).C(common.ColUsers)
	count, err = collection.Find(bson.M{}).Count()

	if count < 1 {
		// Create admin/admin account
		var user models.User
		user = models.User{bson.NewObjectId(), "admin", "admin"}
		err = collection.Insert(&user)
	}

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MongoDbSession != nil {
		db.MongoDbSession.Close()
	}
}
