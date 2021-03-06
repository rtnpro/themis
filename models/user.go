package models

import (
  "time"

	"gopkg.in/mgo.v2/bson"
  "github.com/manyminds/api2go/jsonapi"
)

// UserName stores the common type name.
const UserName = "identities"

// User is a user in the system.
type User struct {
    ID        bson.ObjectId		`bson:"_id,omitempty" json:"-"`
    FullName  string          `bson:"full_name" json:"fullName"`
    ImageURL  string          `bson:"image_url" json:"imageUrl"`
    Username	string          `bson:"username" json:"username"`
    CreatedAt time.Time 	    `bson:"created_at" json:"-"`
    UpdatedAt time.Time	      `bson:"updated_at" json:"-"`
}

// NewUser creates a new User instance.
func NewUser() (user *User) {
  user = new(User)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
  return user
}

// GetCollectionName returns the database collection name.
func (user User) GetCollectionName() string {
  return UserName
}

// GetID returns the ID for marshalling to json.
func (user User) GetID() string {
  return user.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (user *User) SetID(id string) error {
  user.ID = bson.ObjectIdHex(id)
  return nil
}

// GetName returns the entity type name for marshalling to json.
func (user User) GetName() string {
  return user.GetCollectionName()
}

// GetCustomLinks returns the custom links, namely the self link.
func (user User) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
	}
	return links
}
