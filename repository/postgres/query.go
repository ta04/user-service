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
	"database/sql"
	"fmt"

	userPB "github.com/ta04/user-service/proto"
)

// Postgres is the implementor of Postgres interface
type Postgres struct {
	DB *sql.DB
}

// Index indexes all active user
func (repo *Postgres) Index(indexUsersRequest *userPB.IndexUsersRequest) (users []*userPB.User, err error) {
	var id int32
	var firstName, lastName, username, emailAddress, phoneNumber, dateOfBirth, address, role, status string

	query := "SELECT id, first_name, last_name, username, email_address, phone_number, date_of_birth," +
		" address, role, status FROM users WHERE status = 'active'"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &username, &emailAddress, &phoneNumber, &dateOfBirth,
			&address, &role, &status)
		if err != nil {
			return nil, err
		}

		user := &userPB.User{
			Id:           id,
			FirstName:    firstName,
			LastName:     lastName,
			Username:     username,
			EmailAddress: emailAddress,
			PhoneNumber:  phoneNumber,
			DateOfBirth:  dateOfBirth,
			Address:      address,
			Role:         role,
			Status:       status,
		}
		users = append(users, user)
	}

	return users, err
}

// Show shows an active user by id
func (repo *Postgres) Show(user *userPB.User) (*userPB.User, error) {
	var id int32
	var firstName, lastName, username, emailAddress, phoneNumber, dateOfBirth, address, role, status string

	query := fmt.Sprintf("SELECT id, first_name, last_name, username, email_address, phone_number, date_of_birth,"+
		" address, role, status FROM users WHERE id = %d AND status = 'active'", user.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &firstName, &lastName, &username, &emailAddress, &phoneNumber,
		&dateOfBirth, &address, &role, &status)
	if err != nil {
		return nil, err
	}

	return &userPB.User{
		Id:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Username:     username,
		EmailAddress: emailAddress,
		PhoneNumber:  phoneNumber,
		DateOfBirth:  dateOfBirth,
		Address:      address,
		Role:         role,
		Status:       status,
	}, err
}

// ShowByUsername shows an active user by username
func (repo *Postgres) ShowByUsername(user *userPB.User) (*userPB.User, error) {
	var id int32
	var firstName, lastName, username, emailAddress, phoneNumber, dateOfBirth, address, role, status string

	query := fmt.Sprintf("SELECT id, first_name, last_name, username, email_address, phone_number, date_of_birth,"+
		" address, role, status FROM users WHERE username = '%s' AND status = 'active'", user.Username)
	err := repo.DB.QueryRow(query).Scan(&id, &firstName, &lastName, &username, &emailAddress, &phoneNumber,
		&dateOfBirth, &address, &role, &status)
	if err != nil {
		return nil, err
	}

	return &userPB.User{
		Id:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Username:     username,
		EmailAddress: emailAddress,
		PhoneNumber:  phoneNumber,
		DateOfBirth:  dateOfBirth,
		Address:      address,
		Role:         role,
		Status:       status,
	}, err
}

// ShowCredentialsByUsername shows credentials of an active user by username
func (repo *Postgres) ShowCredentialsByUsername(user *userPB.User) (*userPB.User, error) {
	var password, primeNumber, generatorValue int32

	query := fmt.Sprintf("SELECT password, prime_number, generator_value FROM users WHERE"+
		" username = '%s' AND status = 'active'", user.Username)
	err := repo.DB.QueryRow(query).Scan(&password, &primeNumber, &generatorValue)
	if err != nil {
		return nil, err
	}

	return &userPB.User{
		Password:       password,
		PrimeNumber:    primeNumber,
		GeneratorValue: generatorValue,
	}, err

}
