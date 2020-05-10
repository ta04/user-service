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

	userPB "github.com/ta04/user-service/proto"
	userRepo "github.com/ta04/user-service/repository"
)

// Handler is the handler of user service
type Handler struct {
	repository userRepo.Repository
}

// NewHandler creates a new user service handler
func NewHandler(repo userRepo.Repository) *Handler {
	return &Handler{
		repository: repo,
	}
}

// IndexUsers indexes the users
func (h *Handler) IndexUsers(ctx context.Context, req *userPB.IndexUsersRequest, res *userPB.Response) error {
	users, err := h.repository.Index(req)
	if err != nil {
		return err
	}

	res.Users = users
	res.Error = nil

	return err
}

// ShowUser shows a user by ID
func (h *Handler) ShowUser(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.Show(req)
	if err != nil {
		return err
	}

	res.User = user
	res.Error = nil

	return err
}

// ShowUserByUsername shows a user by username
func (h *Handler) ShowUserByUsername(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.ShowByUsername(req)
	if err != nil {
		return err
	}

	res.User = user
	res.Error = nil

	return err
}

// ShowUserCredentialsByUsername shows credentials of a user by username
func (h *Handler) ShowUserCredentialsByUsername(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.ShowCredentialsByUsername(req)
	if err != nil {
		return err
	}

	res.User = user
	res.Error = nil

	return err
}

// StoreUser stores a new user
func (h *Handler) StoreUser(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.User = user
	res.Error = nil

	return err
}

// UpdateUser updates a user
func (h *Handler) UpdateUser(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.Update(req)
	if err != nil {
		return err
	}

	res.User = user
	res.Error = nil

	return err
}

// DestroyUser destroys a user
func (h *Handler) DestroyUser(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.Destroy(req)
	if err != nil {
		return err
	}

	res.User = user
	res.Error = nil

	return err
}
