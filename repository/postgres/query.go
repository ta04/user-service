// user-service/repository/postgres/query.go

package postgres

import (
	"database/sql"
	"fmt"

	userPB "github.com/G0tYou/user-service/proto"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) Index() (users []*userPB.User, err error) {
	var id int32
	var firstName, lastName, username, password, emailAddress, phoneNumber, dateOfBirth, address, role, creditCardNumber, creditCardType, creditCardExpiredMonth, creditCardExpiredYear, creditCardCvv string
	var status bool

	query := "SELCET * FROM users"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &username, &password, &emailAddress, &phoneNumber, &dateOfBirth, &address, &role, &creditCardNumber, &creditCardType, &creditCardExpiredMonth, &creditCardExpiredYear, &creditCardCvv, &status)
		if err != nil {
			return nil, err
		}

		user := &userPB.User{
			Id: id,
			FirstName: firstName,
			LastName: lastName,
			Username: username,
			Password: password,
			EmailAddress: emailAddress,
			PhoneNumber: phoneNumber,
			DateOfBirth: dateOfBirth,
			Address: address,
			Role: role,
			CreditCardNumber: creditCardNumber,
			CreditCardType: creditCardType,
			CreditCardExpiredMonth: creditCardExpiredMonth,
			CreditCardExpiredYear: creditCardExpiredYear,
			CreditCardCvv: creditCardCvv,
			Status: status,
		}
		users = append(users, user)
	}

	return users, err
}

func (repo *Repository) Show(user *userPB.User) (*userPB.User, error) {
	var id int32
	var firstName, lastName, username, password, emailAddress, phoneNumber, dateOfBirth, address, role, creditCardNumber, creditCardType, creditCardExpiredMonth, creditCardExpiredYear, creditCardCvv string
	var status bool

	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", user.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &firstName, &lastName, &username, &password, &emailAddress, &phoneNumber, &dateOfBirth, &address, &role, &creditCardNumber, &creditCardType, &creditCardExpiredMonth, &creditCardExpiredYear, &creditCardCvv, &status)
	if err != nil {
		return nil, err
	}

	return &userPB.User{
		Id: id,
		FirstName: firstName,
		LastName: lastName,
		Username: username,
		Password: password,
		EmailAddress: emailAddress,
		PhoneNumber: phoneNumber,
		DateOfBirth: dateOfBirth,
		Address: address,
		Role: role,
		CreditCardNumber: creditCardNumber,
		CreditCardType: creditCardType,
		CreditCardExpiredMonth: creditCardExpiredMonth,
		CreditCardExpiredYear: creditCardExpiredYear,
		CreditCardCvv: creditCardCvv,
		Status: status,
	}, err
}

func (repo *Repository) ShowByUsername(user *userPB.User) (*userPB.User, error) {
	var id int32
	var firstName, lastName, username, password, emailAddress, phoneNumber, dateOfBirth, address, role, creditCardNumber, creditCardType, creditCardExpiredMonth, creditCardExpiredYear, creditCardCvv string
	var status bool

	query := fmt.Sprintf("SELECT * FROM users WHERE username = %s", user.Username)
	err := repo.DB.QueryRow(query).Scan(&id, &firstName, &lastName, &username, &password, &emailAddress, &phoneNumber, &dateOfBirth, &address, &role, &creditCardNumber, &creditCardType, &creditCardExpiredMonth, &creditCardExpiredYear, &creditCardCvv, &status)
	if err != nil {
		return nil, err
	}

	return &userPB.User{
		Id: id,
		FirstName: firstName,
		LastName: lastName,
		Username: username,
		Password: password,
		EmailAddress: emailAddress,
		PhoneNumber: phoneNumber,
		DateOfBirth: dateOfBirth,
		Address: address,
		Role: role,
		CreditCardNumber: creditCardNumber,
		CreditCardType: creditCardType,
		CreditCardExpiredMonth: creditCardExpiredMonth,
		CreditCardExpiredYear: creditCardExpiredYear,
		CreditCardCvv: creditCardCvv,
		Status: status,
	}, err
}
