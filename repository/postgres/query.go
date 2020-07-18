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

	proto "github.com/ta04/user-service/model/proto"
)

// Postgres is the implementor of Repository interface
type Postgres struct {
	DB *sql.DB
}

// NewPostgres will create a new postgres instance
func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		DB: db,
	}
}

// GetAllByQuery will get all user by query
func (postgres *Postgres) GetAllByQuery(request *proto.GetAllUsersRequest) (*[]*proto.User, error) {
	var id int32
	var firstName, lastName, username, emailAddress, phoneNumber, dateOfBirth, address, role, status string
	var users []*proto.User

	query := fmt.Sprintf("SELECT id, first_name, last_name, username, email_address, phone_number,"+
		" date_of_birth, address, role, status FROM users WHERE (LOWER(first_name) LIKE '%%%s%%' OR LOWER(last_name)"+
		" LIKE '%%%s%%' OR LOWER(email) LIKE '%%%s%%') AND role = '%s' AND status = '%s'", request.Query, request.Query,
		request.Query, request.Role, request.Status)
	rows, err := postgres.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &username, &emailAddress, &phoneNumber, &dateOfBirth,
			&address, &role, &status)
		if err != nil {
			return nil, err
		}

		user := &proto.User{
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

	return &users, err
}

// GetAll will get all user
func (postgres *Postgres) GetAll(request *proto.GetAllUsersRequest) (*[]*proto.User, error) {
	var id int32
	var firstName, lastName, username, emailAddress, phoneNumber, dateOfBirth, address, role, status string
	var users []*proto.User

	query := fmt.Sprintf("SELECT id, first_name, last_name, username, email_address, phone_number, "+
		"date_of_birth, address, role, status FROM users WHERE role = '%s' AND status = '%s'", request.Role, request.Status)
	rows, err := postgres.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &username, &emailAddress, &phoneNumber, &dateOfBirth,
			&address, &role, &status)
		if err != nil {
			return nil, err
		}

		user := &proto.User{
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

	return &users, err
}

// GetOneByUsername will get a user by username
func (postgres *Postgres) GetOneByUsername(request *proto.GetOneUserRequest) (*proto.User, error) {
	var id int32
	var firstName, lastName, username, emailAddress, phoneNumber, dateOfBirth, address, role, status string

	query := fmt.Sprintf("SELECT id, first_name, last_name, username, email_address, phone_number, date_of_birth,"+
		" address, role, status FROM users WHERE username = '%s'", request.Username)
	err := postgres.DB.QueryRow(query).Scan(&id, &firstName, &lastName, &username, &emailAddress, &phoneNumber,
		&dateOfBirth, &address, &role, &status)
	if err != nil {
		return nil, err
	}

	return &proto.User{
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

// GetOneCredentialsByUsername will get a user's credentials by username
func (postgres *Postgres) GetOneCredentialsByUsername(request *proto.GetOneUserRequest) (*proto.User, error) {
	var password, primeNumber, generatorValue int32

	query := fmt.Sprintf("SELECT password, prime_number, generator_value FROM users WHERE username = '%s'",
		request.Username)
	err := postgres.DB.QueryRow(query).Scan(&password, &primeNumber, &generatorValue)
	if err != nil {
		return nil, err
	}

	return &proto.User{
		Password:       password,
		PrimeNumber:    primeNumber,
		GeneratorValue: generatorValue,
	}, err
}

// GetOneCredentialsByUsername will get a user's credentials by id
func (postgres *Postgres) GetOneCredentials(request *proto.GetOneUserRequest) (*proto.User, error) {
	var password, primeNumber, generatorValue int32

	query := fmt.Sprintf("SELECT password, prime_number, generator_value FROM users WHERE id = %d",
		request.Id)
	err := postgres.DB.QueryRow(query).Scan(&password, &primeNumber, &generatorValue)
	if err != nil {
		return nil, err
	}

	return &proto.User{
		Password:       password,
		PrimeNumber:    primeNumber,
		GeneratorValue: generatorValue,
	}, err
}

// GetOne will get a user by id
func (postgres *Postgres) GetOne(request *proto.GetOneUserRequest) (*proto.User, error) {
	var id int32
	var firstName, lastName, username, emailAddress, phoneNumber, dateOfBirth, address, role, status string

	query := fmt.Sprintf("SELECT id, first_name, last_name, username, email_address, phone_number, date_of_birth,"+
		" address, role, status FROM users WHERE id = %d", request.Id)
	err := postgres.DB.QueryRow(query).Scan(&id, &firstName, &lastName, &username, &emailAddress, &phoneNumber,
		&dateOfBirth, &address, &role, &status)
	if err != nil {
		return nil, err
	}

	return &proto.User{
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
