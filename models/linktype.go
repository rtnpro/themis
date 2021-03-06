package models

import (
	"time"

	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// LinkTypeName stores the common type name.
const LinkTypeName = "workitemlinktypes"

// LinkType is a type for a link.
type LinkType struct {
    ID                  	bson.ObjectId       `bson:"_id,omitempty" json:"-"`
    Name                	string              `bson:"name" json:"name"`
    Description         	string              `bson:"description" json:"description"`
    ForwardName         	string              `bson:"forward_name" json:"forward_name"`
    ReverseName         	string              `bson:"reverse_name" json:"reverse_name"`
    Topology            	string              `bson:"topology" json:"topology"`
    Version             	int                 `bson:"version" json:"version"`
		LinkCategoryID				bson.ObjectId				`bson:"link_category_id" json:"-"`
		SourceWorkItemTypeID	bson.ObjectId				`bson:"source_workitemtype_id" json:"-"`
		TargetWorkItemTypeID	bson.ObjectId				`bson:"target_workitemtype_id" json:"-"`
		SpaceID             	bson.ObjectId       `bson:"space_id" json:"-"`
    CreatedAt 	        	time.Time  		    	`bson:"created_at" json:"-"`
    UpdatedAt 	        	time.Time		    		`bson:"updated_at" json:"-"`
    CategoryRef           string              `bson:"-" json:"-"`
    SourceWorkItemTypeRef string              `bson:"-" json:"-"`
    TargetWorkItemTypeRef string              `bson:"-" json:"-"`
}

// NewLinkType creates a new LinkType instance.
func NewLinkType() (linkType *LinkType) {
  linkType = new(LinkType)
	linkType.CreatedAt = time.Now()
	linkType.UpdatedAt = time.Now()
  return linkType
}

// GetCollectionName returns the database collection name.
func (linkType LinkType) GetCollectionName() string {
  return LinkTypeName
}

// GetID returns the ID for marshalling to json.
func (linkType LinkType) GetID() string {
  return linkType.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (linkType *LinkType) SetID(id string) error {
  linkType.ID = bson.ObjectIdHex(id)
  return nil
}

// GetName returns the entity type name for marshalling to json.
func (linkType LinkType) GetName() string {
  return linkType.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (linkType LinkType) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "linkcategories",
			Name: "link_category",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type: "workitemtypes",
			Name: "source_type",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type: "workitemtypes",
			Name: "target_type",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type: "spaces",
			Name: "space",
			IsNotLoaded: false, // we want to have the data field
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (linkType LinkType) GetReferencedIDs() []jsonapi.ReferenceID {
  result := []jsonapi.ReferenceID {
			jsonapi.ReferenceID {
	    	ID:   linkType.LinkCategoryID.Hex(),
 	   		Type: "linkcategories",
 	   		Name: "link_category",
			},
			jsonapi.ReferenceID {
	    	ID:   linkType.SourceWorkItemTypeID.Hex(),
 	   		Type: "workitemtypes",
 	   		Name: "source_type",
			},
			jsonapi.ReferenceID {
	    	ID:   linkType.TargetWorkItemTypeID.Hex(),
 	   		Type: "workitemtypes",
 	   		Name: "target_type",
			},
			jsonapi.ReferenceID{
				ID:   linkType.SpaceID.Hex(),
				Type: "spaces",
				Name: "space",
			},
	}
	return result
}

// GetCustomLinks returns the custom links, namely the self link.
func (linkType LinkType) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
	}
	return links
}
