// user-service/repository/postgres/repository.go

package postgres

import (
	"fmt"

	userPB "github.com/G0tYou/user-service/proto"
)

// Store will store a new user
func (repo *Repository) Store(user *userPB.User) (*userPB.User, error) {
	query := fmt.Sprintf("INSERT INTO users (first_name, last_name, username, password, email_address,"+
		" phone_number, date_of_birth, address, role, credit_card_number, credit_card_type, credit_card_expired_month,"+
		" credit_card_expired_year, credit_card_cvv, status) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s',"+
		" '%s', '%s', '%s', '%s', '%s', 'active')",
		user.FirstName, user.LastName, user.Username, user.Password, user.EmailAddress, user.PhoneNumber,
		user.DateOfBirth, user.Address, user.Role, user.CreditCardNumber, user.CreditCardType, user.CreditCardExpiredMonth,
		user.CreditCardExpiredYear, user.CreditCardCvv)
	_, err := repo.DB.Exec(query)

	return user, err
}

// Update will update an existing user's data
func (repo *Repository) Update(user *userPB.User) (*userPB.User, error) {
	query := fmt.Sprintf("UPDATE users SET first_name = '%s', last_name = '%s', username = '%s', password = '%s',"+
		" email_address = '%s', phone_number = '%s', date_of_birth = '%s', address = '%s', role = '%s',"+
		" credit_card_number = '%s', credit_card_type = '%s', credit_card_expired_month = '%s', credit_card_expired_year = '%s',"+
		" credit_card_cvv = '%s', status = 'active' WHERE id = %d", user.FirstName, user.LastName, user.Username,
		user.Password, user.EmailAddress, user.PhoneNumber, user.DateOfBirth, user.Address, user.Role,
		user.CreditCardNumber, user.CreditCardType, user.CreditCardExpiredMonth, user.CreditCardExpiredYear,
		user.CreditCardCvv, user.Id)
	_, err := repo.DB.Exec(query)

	return user, err
}

// Destroy will update an existing user's status to inactive
func (repo *Repository) Destroy(user *userPB.User) (*userPB.User, error) {
	query := fmt.Sprintf("UPDATE users SET status = 'inactive' WHERE id = %d", user.Id)
	_, err := repo.DB.Exec(query)

	return user, err
}
