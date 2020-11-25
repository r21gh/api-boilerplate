package mongo

import (
	"api-boilerplate/pkg"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type userModel struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Password string
}

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key:	[]string{"username"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
}

func newUserModel(u *root.User) *userModel {
	return &userModel{
		Username: u.Username,
		Password: u.Password,
	}
}

func (u *userModel) toRootUser() *root.User {
	return &root.User{
		ID: u.ID.Hex(),
		Username: u.Username,
		Password: u.Password,
	}
}

