package models

import "gopkg.in/mgo.v2/bson"

// Movie information
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	URL         string        `bson:"url" json:"url"`
	CoverImage  string        `bson:"coerImage" json:"coverImage"`
	Description string        `bson:"decription" json:"description"`
}

// AddMovie information
type AddMovie struct {
	Name        string `json:"name" example:"Movie Name"`
	URL         string `json:"url" example:"movie url"`
	CoverImage  string `json:"coverImage" example:"Movie Cover Image"`
	Description string `json:"description" example:"Movie description"`
}
