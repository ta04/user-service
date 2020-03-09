package handler

import (
	"context"

	userPB "github.com/G0tYou/user-service/proto"
	userRepo "github.com/G0tYou/user-service/repository"
)

type handler struct {
	repository userRepo.Repository
}

func NewHandler(repo userRepo.Repository) *handler {
	return &handler{
		repository: repo,
	}
}

func (h *handler) StoreUser(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.User = user
	res.Error = nil

	return err
}

func (h *handler) IndexUser(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	users, err := h.repository.Index()

	res.Users = users
	res.Error = nil

	return err
}

func (h *handler) ShowUser(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.Show(req)

	res.User = user
	res.Error = nil

	return err
}

func (h *handler) UpdateUser(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.Update(req)
	if err != nil {
		return err
	}

	res.User = user
	res.Error = nil

	return err
}

func (h *handler) DestroyUser(ctx context.Context, req *userPB.User, res *userPB.Response) error {
	user, err := h.repository.Destroy(req)
	if err != nil {
		return err
	}

	res.User = user
	res.Error = nill

	return err
}
