package models

import (
	"errors"

	"github.com/bangbugs/go-microservice-1/user-service/common"
	"gopkg.in/mgo.v2/bson"
)

// User Info
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Name     string        `bson:"name" json:"name" example:"thanos"`
	Password string        `bson:"password" json:"password" example:"tahnos@123"`
}

// AddUser Info
type AddUser struct {
	Name     string `json:"name" example:"User Name"`
	Password string `json:"password" example:"user password"`
}

// Validate User
func (a AddUser) Validate() error {
	switch {
	case len(a.Name) == 0:
		return errors.New(common.ErrNameEmpty)
	case len(a.Password) == 0:
		return errors.New(common.ErrPasswordEmpty)
	default:
		return nil
	}
}
