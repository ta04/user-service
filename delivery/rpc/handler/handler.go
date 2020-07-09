/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package handler

import (
	"context"
	"errors"

	proto "github.com/ta04/user-service/model/proto"
	"github.com/ta04/user-service/usecase"
)

// Handler is the handler of user service
type Handler struct {
	Usecase usecase.Usecase
}

// NewHandler will create a new user handler
func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

// GetAllUsers will get all users
func (handler *Handler) GetAllUsers(ctx context.Context, req *proto.GetAllUsersRequest, res *proto.Response) error {
	users, err := handler.Usecase.GetAll(req)
	if err != nil {
		res.Users = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Users = users
	res.Error = nil

	return nil
}

// GetOneUser will get a user
func (handler *Handler) GetOneUser(ctx context.Context, req *proto.GetOneUserRequest, res *proto.Response) error {
	user, err := handler.Usecase.GetOne(req)
	if err != nil {
		res.User = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.User = user
	res.Error = nil

	return nil
}

// CreateOneUser will create a new user
func (handler *Handler) CreateOneUser(ctx context.Context, req *proto.User, res *proto.Response) error {
	user, err := handler.Usecase.CreateOne(req)
	if err != nil {
		res.User = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.User = user
	res.Error = nil

	return nil
}

// UpdateOneUser will update a user
func (handler *Handler) UpdateOneUser(ctx context.Context, req *proto.User, res *proto.Response) error {
	user, err := handler.Usecase.UpdateOne(req)
	if err != nil {
		res.User = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.User = user
	res.Error = nil

	return nil
}
