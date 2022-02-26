package daos

import (
	"github.com/bangbugs/go-microservice-1/user-service/common"
	"github.com/bangbugs/go-microservice-1/user-service/databases"
	"github.com/bangbugs/go-microservice-1/user-service/models"
	"github.com/bangbugs/go-microservice-1/user-service/utils"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	utils *utils.Utils
}

// GetAll -> get the list of users
func (u *User) GetAll() ([]models.User, error) {
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute query
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	var users []models.User
	err := collection.Find(bson.M{}).All(&users)

	return users, err
}

// GetByID -> finds a user by id
func (u *User) GetByID(id string) (models.User, error) {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return models.User{}, err
	}

	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute query
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	var user models.User
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)

	return user, err
}

// DeleteByID -> finds and deletes a user by id
func (u *User) DeleteByID(id string) error {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return err
	}

	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Login -> login user
func (u *User) Login(name, password string) (models.User, error) {
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	var user models.User
	err := collection.Find(bson.M{"$and": []bson.M{{"name": name}, {"password": password}}}).One(&user)
	return user, err
}

// Insert -> insert new user into database
func (u *User) Insert(user models.User) error {
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err := collection.Insert(&user)
	return err
}

// Delete -> delete new user into database
func (u *User) Delete(user models.User) error {
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err := collection.Remove(&user)
	return err
}

// Update -> modifies exisitng user into database
func (u *User) Update(user models.User) error {
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err := collection.UpdateId(user.ID, &user)
	return err
}
