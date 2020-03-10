// user-service/repository/postgres/repository.go

package postgres

import (
	"fmt"

	userPB "github.com/SleepingNext/user-service/proto"
)

func (repo *Repository) Store(user *userPB.User) (*userPB.User, error) {
	query := fmt.Sprintf("INSERT INTO users (first_name, last_name, username, password, email_address,"+
		" phone_number, date_of_birth, address, role, status) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %s)",
		user.FirstName, user.LastName, user.Username, user.Password, user.EmailAddress, user.PhoneNumber,
		user.DateOfBirth, user.Address, user.Role, user.Status)
	_, err := repo.DB.Exec(query)

	return user, err
}

func (repo *Repository) Update(user *userPB.User) (*userPB.User, error) {
	query := fmt.Sprintf("UPDATE users SET first_name = '%s', last_name = '%s', username = '%s', password = '%s',"+
		" email_address = '%s', phone_number = '%s', date_of_birth = '%s', address = '%s', role = '%s',"+
		" credit_card_number = %s, credit_card_type = %s, credit_card_expired_month = %s, credit_card_expired_year = %s,"+
		" credit_card_cvv = %s, status = %s WHERE id = %d", user.FirstName, user.LastName, user.Username,
		user.Password, user.EmailAddress, user.PhoneNumber, user.DateOfBirth, user.Address, user.Role,
		user.CreditCardNumber, user.CreditCardType, user.CreditCardExpiredMonth, user.CreditCardExpiredYear,
		user.CreditCardCvv, user.Status, user.Id)
	_, err := repo.DB.Exec(query)

	return user, err
}

func (repo *Repository) Destroy(user *userPB.User) (*userPB.User, error) {
	query := fmt.Sprintf("DELETE FROM users WHERE id=%d", user.Id)
	_, err := repo.DB.Exec(query)

	return user, err
}
