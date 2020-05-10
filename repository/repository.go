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

import (
	userPB "github.com/ta04/user-service/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	Index(*userPB.IndexUsersRequest) ([]*userPB.User, error)
	Show(*userPB.User) (*userPB.User, error)
	ShowByUsername(*userPB.User) (*userPB.User, error)
	ShowCredentialsByUsername(*userPB.User) (*userPB.User, error)
	Store(*userPB.User) (*userPB.User, error)
	Update(*userPB.User) (*userPB.User, error)
	Destroy(*userPB.User) (*userPB.User, error)
}
