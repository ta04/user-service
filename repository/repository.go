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

package repository

import proto "github.com/ta04/user-service/model/proto"

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	GetAllByQuery(request *proto.GetAllUsersRequest) ([]*proto.User, error)
	GetAll(request *proto.GetAllUsersRequest) ([]*proto.User, error)
	GetOneByUsername(request *proto.GetOneUserRequest) (*proto.User, error)
	GetOneCredentialsByUsername(request *proto.GetOneUserRequest) (*proto.User, error)
	GetOneCredentials(request *proto.GetOneUserRequest) (*proto.User, error)
	GetOne(request *proto.GetOneUserRequest) (*proto.User, error)
	CreateOne(user *proto.User) (*proto.User, error)
	UpdateOne(user *proto.User) (*proto.User, error)
}
