package v1

import (
	"net/http"

	proto "github.com/ta04/user-service/model/proto"
	"github.com/ta04/user-service/repository"
)

type Usecase struct {
	Repository repository.Repository
}

func NewUsecase(repository repository.Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}

func (usecase *Usecase) GetAll(request *proto.GetAllUsersRequest) (*[]*proto.User, *proto.Error) {
	if request == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	if request.Status == "" {
		request.Status = "active"
	}

	if request.Role == "" {
		request.Role = "customer"
	}

	var users *[]*proto.User
	var err error

	if request.Query != "" {
		users, err = usecase.Repository.GetAllByQuery(request)
	} else {
		users, err = usecase.Repository.GetAll(request)
	}

	if users == nil || len(*users) <= 0 {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: "no users found",
		}
	}
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}

	return users, nil
}

func (usecase *Usecase) GetOne(request *proto.GetOneUserRequest) (*proto.User, *proto.Error) {
	if request == nil || (request.Id == 0 && request.Username == "") {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	var user *proto.User
	var err error

	if request.Username != "" && request.WithCredentials {
		user, err = usecase.Repository.GetOneCredentialsByUsername(request)
		if err != nil {
			return nil, &proto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if request.Username != "" && !request.WithCredentials {
		user, err = usecase.Repository.GetOneByUsername(request)
		if err != nil {
			return nil, &proto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	if request.Id != 0 && request.WithCredentials {
		user, err = usecase.Repository.GetOneCredentials(request)
		if err != nil {
			return nil, &proto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else if request.Id != 0 && !request.WithCredentials {
		user, err = usecase.Repository.GetOne(request)
		if err != nil {
			return nil, &proto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	return user, nil
}

func (usecase *Usecase) CreateOne(user *proto.User) (*proto.User, *proto.Error) {
	if user == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	user, err := usecase.Repository.CreateOne(user)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return user, nil
}

func (usecase *Usecase) UpdateOne(user *proto.User) (*proto.User, *proto.Error) {
	if user == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	user, err := usecase.Repository.UpdateOne(user)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error,
		}
	}

	return user, nil
}
