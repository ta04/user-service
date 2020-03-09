// user-service/repository/posgres/query.go

package postgres

import(
	"database/sql"
	"fmt"

	userPB "github.com/G0tYou/user-service/proto"
)

type Repository struct{
	DB *sql.DB
}

func (repo *Repository) Index() (users []*userPB.User, err error)  {
	var int32 id 
	var first_name, last_name, username, password, email_address, phone_number, date_of_birth, address, role, credit_card_number, credit_card_type, credit_card_expired_month, credit_card_expired_year, credit_card_cvv string 
	var status bool
	
	query := "SELCET * FROM users"
	rows, err := repo.DB.Query(query)
	if err != nil{
		return nil, err
	}
	
	for rows.Next(){
		err := rows.Scan(&id, &first_name, &last_name, &username, &password, &email_address, &phone_number, &date_of_birth, &address, &role, &credit_card_number, &credit_card_type, &credit_card_expired_month, &credit_card_expired_year, &credit_card_cvv, &status)
		if err != nil{
			return nil, err
		}

		user:= &userPB.User{
			Id: : id,
			FirstName : first_name,
			LastName : last_name,
			Username : username,
			Password : password,
			EmailAddress : email_address,
			PhoneNumber : phone_number,
			DataOfBirth : date_of_birth,
			Address : address,
			Role : role,
			CreditCardNumber : credit_card_number,
			CreditCardExpiredMonth : credit_card_expired_month,
			CreditCardExpiredYear : credit_card_expired_year,
			CreditCardCvv : credit_card_cvv,
			Status : status
		}
		users = append(users, user)
	}

	return users, err
}

func (repo *Repository) Show(user *userPB.User) (*userPB.User, error)  {
	var int32 id 
	var first_name, last_name, username, password, email_address, phone_number, date_of_birth, address, role, credit_card_number, credit_card_type, credit_card_expired_month, credit_card_expired_year, credit_card_cvv string 
	var status bool

	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", user.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &first_name, &last_name, &username, &password, &email_address, &phone_number, &date_of_birth, &address, &role, &credit_card_number, &credit_card_type, &credit_card_expired_month, &credit_card_expired_year, &credit_card_cvv, &status)
	iff err != nil{
		return nil, err
	}

	return &userPB.User{
			Id: : id,
			FirstName : first_name,
			LastName : last_name,
			Username : username,
			Password : password,
			EmailAddress : email_address,
			PhoneNumber : phone_number,
			DataOfBirth : date_of_birth,
			Address : address,
			Role : role,
			CreditCardNumber : credit_card_number,
			CreditCardType : credit_card_type,
			CreditCardExpiredMonth : credit_card_expired_month,
			CreditCardExpiredYear : credit_card_expired_year,
			CreditCardCvv : credit_card_cvv,
			Status : status
	}
}