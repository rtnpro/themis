package resources

import (
	"errors"
	"net/http"

	"github.com/manyminds/api2go"
	"gopkg.in/mgo.v2/bson"

	"themis/utils"
	"themis/models"
	"themis/database"
)

// IterationResource for api2go routes.
type IterationResource struct {
	IterationStorage database.IIterationStorage
	WorkItemStorage database.IWorkItemStorage
}

func (c IterationResource) getFilterFromRequest(r api2go.Request) (bson.M, *utils.NestedEntityError) {
	var filter bson.M
	// Getting reference context
	sourceContext, sourceContextID, thisContext := utils.ParseContext(r)
	switch sourceContext {
		case models.IterationName:
			entity, err := c.IterationStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, &utils.NestedEntityError { InnerError: err, Code: 0 }
			}
			if thisContext == "parent" {
				if entity.ParentIterationID.Hex()=="" {
					// this is the root iteration
					return nil, &utils.NestedEntityError { InnerError: nil, Code: 42 }
				}
				filter = bson.M{"_id": entity.ParentIterationID}
			}
		case models.WorkItemName:
			entity, err := c.WorkItemStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, &utils.NestedEntityError { InnerError: err, Code: 0 }
			}
			if thisContext == "iteration" {
				filter = bson.M{"_id": entity.IterationID}
			}
		default:
			// build standard filter expression
			filter = utils.BuildDbFilterFromRequest(r)
	}
	return filter, nil
}

// FindAll Iterations.
func (c IterationResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	// build filter expression
	filter, err := c.getFilterFromRequest(r)
	if err != nil && err.Code==42 {
		// this is the root iteration
		var empty []models.Iteration
		return &api2go.Response{Res: empty}, nil
	}
	if err != nil {
		return &api2go.Response{}, err.InnerError
	}
	
	iterations, _ := c.IterationStorage.GetAll(filter)
	return &api2go.Response{Res: iterations}, nil
}

// PaginatedFindAll can be used to load users in chunks.
// Possible success status code 200.
func (c IterationResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {

	// build filter expression
	filter, nestedErr := c.getFilterFromRequest(r)
	if nestedErr != nil && nestedErr.Code==42 {
		// this is the root iteration
		var empty []models.Iteration
		return 0, &api2go.Response{Res: empty}, nil
	}
	if nestedErr != nil {
		return 0, &api2go.Response{}, nestedErr.InnerError
	}

	// parse out offset and limit
	queryOffset, queryLimit, err := utils.ParsePaging(r)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get the paged data from storage
	result, err := c.IterationStorage.GetAllPaged(filter, queryOffset, queryLimit)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get total count for paging
	allCount, err := c.IterationStorage.GetAllCount(filter)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// return everything
	return uint(allCount), &api2go.Response{Res: result}, nil
}

// FindOne Iteration.
// Possible success status code 200
func (c IterationResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	utils.DebugLog.Printf("Received FindOne with ID %s.", id)
	res, err := c.IterationStorage.GetOne(bson.ObjectIdHex(id))
	return &api2go.Response{Res: res}, err
}

// Create a new Iteration.
// Possible status codes are:
// - 201 Created: Resource was created and needs to be returned
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
//   the server
func (c IterationResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	iteration, ok := obj.(models.Iteration)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	id, _ := c.IterationStorage.Insert(iteration)
	iteration.ID = id
	return &api2go.Response{Res: iteration, Code: http.StatusCreated}, nil
}

// Delete a Iteration.
// Possible status codes are:
// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Deletion was successful, return nothing
func (c IterationResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.IterationStorage.Delete(bson.ObjectIdHex(id))
	return &api2go.Response{Code: http.StatusOK}, err
}

// Update a Iteration.
// Possible status codes are:
// - 200 OK: Update successful, however some field(s) were changed, returns updates source
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
func (c IterationResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	iteration, ok := obj.(models.Iteration)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	err := c.IterationStorage.Update(iteration)
	return &api2go.Response{Res: iteration, Code: http.StatusNoContent}, err
}
