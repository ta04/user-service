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

	proto "github.com/ta04/user-service/model/proto"
)

// CreateOne will create a new user
func (postgres *Postgres) CreateOne(user *proto.User) (*proto.User, error) {
	query := fmt.Sprintf("INSERT INTO users (first_name, last_name, username, password,"+
		" prime_number, generator_value, email_address, phone_number, date_of_birth, address,"+
		" role, status) VALUES ('%s', '%s', '%s', %d, %d, %d, '%s', '%s', '%s', '%s', '%s', 'active')",
		user.FirstName, user.LastName, user.Username, user.Password, user.PrimeNumber, user.GeneratorValue,
		user.EmailAddress, user.PhoneNumber, user.DateOfBirth, user.Address, user.Role)
	_, err := postgres.DB.Exec(query)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateOne will update a user
func (postgres *Postgres) UpdateOne(user *proto.User) (*proto.User, error) {
	query := fmt.Sprintf("UPDATE users SET first_name = '%s', last_name = '%s', username = '%s',"+
		" password = %d, prime_number = %d, generator_value = %d, email_address = '%s', phone_number = '%s',"+
		" date_of_birth = '%s', address = '%s', role = '%s', status = '%s' WHERE id = %d",
		user.FirstName, user.LastName, user.Username, user.Password, user.PrimeNumber, user.GeneratorValue,
		user.EmailAddress, user.PhoneNumber, user.DateOfBirth, user.Address, user.Role, user.Status, user.Id)
	_, err := postgres.DB.Exec(query)
	if err != nil {
		return nil, err
	}

	return user, nil
}
