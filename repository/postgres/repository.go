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

package postgres

import (
	"fmt"

	userPB "github.com/ta04/user-service/proto"
)

// Store stores a new user
func (repo *Postgres) Store(user *userPB.User) (*userPB.User, error) {
	query := fmt.Sprintf("INSERT INTO users (first_name, last_name, username, password,"+
		" prime_number, generator_value, email_address, phone_number, date_of_birth, address,"+
		" role, status) VALUES ('%s', '%s', '%s', '%d', '%d', '%d', '%s', '%s', '%s', '%s', '%s')",
		user.FirstName, user.LastName, user.Username, user.Password, user.PrimeNumber, user.GeneratorValue,
		user.EmailAddress, user.PhoneNumber, user.DateOfBirth, user.Address, user.Role)
	_, err := repo.DB.Exec(query)

	return user, err
}

// Update updates a user
func (repo *Postgres) Update(user *userPB.User) (*userPB.User, error) {
	query := fmt.Sprintf("UPDATE users SET first_name = '%s', last_name = '%s', username = '%s',"+
		" password = '%d', prime_number = %d, generator_value = %d, email_address = '%s', phone_number = '%s',"+
		" date_of_birth = '%s', address = '%s', role = '%s', status = 'active' WHERE id = %d",
		user.FirstName, user.LastName, user.Username, user.Password, user.PrimeNumber, user.GeneratorValue,
		user.EmailAddress, user.PhoneNumber, user.DateOfBirth, user.Address, user.Role, user.Id)
	_, err := repo.DB.Exec(query)

	return user, err
}

// Destroy updates a user's status to inactive
func (repo *Postgres) Destroy(user *userPB.User) (*userPB.User, error) {
	query := fmt.Sprintf("UPDATE users SET status = 'inactive' WHERE id = %d", user.Id)
	_, err := repo.DB.Exec(query)

	return user, err
}
