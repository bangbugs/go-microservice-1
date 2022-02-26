package daos

import (
	"github.com/bangbugs/go-microservice-1/movie-service/databases"
	"github.com/bangbugs/go-microservice-1/movie-service/models"
	"gopkg.in/mgo.v2/bson"
)

type Movie struct{}

const COLLECTION = "movies"

// GetAll -> get list of movies
func (m *Movie) GetAll() ([]models.Movie, error) {
	// Creating the session
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	// fetching collection
	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	var movies []models.Movie

	err := collection.Find(bson.M{}).All(&movies)

	return movies, err
}

// GetByID -> finds a movies based on id
func (m *Movie) GetByID(id string) (models.Movie, error) {
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	var movie models.Movie

	err := collection.FindId(bson.ObjectIdHex(id)).One(&movie)

	return movie, err
}

// Insert -> adds a new movie
func (m *Movie) Insert(movie models.Movie) error {
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Clone()

	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	err := collection.Insert(&movie)

	return err
}

// Delete -> removes a movie
func (m *Movie) Delete(movie models.Movie) error {
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	err := collection.Remove(&movie)

	return err
}

// Update -> modifies an existing movie
func (m *Movie) Update(movie models.Movie) error {
	sessionCopy := databases.Database.MongoDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	err := collection.UpdateId(movie.ID, &movie)

	return err
}
