package repository

import (
	userPB "github.com/G0tYou/user-service/proto"
)

type Repository interface{
	Store(*userPB.User) (*userPB.User, error)
	Index() ([]*userPB.User, error)
	Show(*userPB.User) (*userPB.User, error)
	Update(*userPB.User) (*userPB.User, error)
	Destroy(*userPB.User) (*userPB.User, error)
}
