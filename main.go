package main

import(
	"context"
	"database/sql"
	"fmt"
	pb "github.com/G0tYou/user-service/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

const(
	port = ":50051";
)

type repository interface{
	Create(*pb.User) (*pb.User, error)
	GetALL() ([]*pb.User, error)
	GetOne(*pb.User) (*pb.User, error)
	Update(*pb.User) (*pb.User, error)
	Delete(*pb.User) (*pb.User, error)
}

type Repository struct {
	mu sync.RWMutex
	db *sql.DB
}

func (repo *Repository) Create(user *pb.User) (*pb.User, error){
	repo.mu.Lock()
	query := fmt.Sprintf("INSERT INTO users (first_name, last_name, username, password, email_address, phone_number, date_of_birth, address, role, credit_card_number, credit_card_type, credit_card_expired_month, credit_card_expired_year, credit_card_cvv, status)"+
		"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%t')", user.FirstName, user.LastName, user.Username, user.Password, user.EmailAddress, user.PhoneNumber, user.DateOfBirth, user.Address, user.Role, user.CreditCardNumber, user.CreditCardType, user.CreditCardExpiredMonth, user.CreditCardExpiredYear, user.CreditCardCvv, user.Status)
		_, err := repo.db.Exec(query)
		repo.mu.Unlock()

		return user, err
}

func (repo *Repository) GetAll() (users []*pb.User, err error){
	var id int32
	var first_name, last_name, username, password, email_address, phone_number, address, role, credit_card_number, credit_card_type, credit_card_expired_month, credit_card_expired_year,credit_card_cvv string
	var status bool
	query := "SELECT * FROM users"
	rows, err := repo.db.Query(query)

	for rows.Next(){
		err := rows.Scan(&id, &first_name, &last_name, &username, &password, &email_address, &phone_number, &address, &role, &credit_card_number, &credit_card_type, &credit_card_expired_month,		err := rows.Scan(&id, &first_name, &last_name, &username, &password, &email_address, &phone_number, &address, &role, &credit_card_number, &credit_card_type, &credit_card_expired_month,		err := rows.Scan(&id, &first_name, &last_name, &username, &password, &email_address, &phone_number, &address, &role, &credit_card_number, &credit_card_type, &credit_card_expired_month		err := rows.Scan(&id, &first_name, &last_name, &username, &password, &email_address, &phone_number, &address, &role, &credit_card_number, &credit_card_type, &credit_card_expired_month, &credit_card_expired_year, &credit_card_ccv, &status)
		if err != nil{
			return nil,err
		}

		user := pb.User{
			Id : id,
			FirstName:first_name,
			LastName:last_name,
			Username:username,
			Password:password,
			EmailAddress:email_address,
			PhoneNumber:phone_number,
			Address:address,
			Role:role,
			CreditCardNumber:credit_card_number,
			CreditCardType:credit_card_type,
			CreditCardExpiredMonth:credit_card_expired_month,
			CreditCardExpiredYear:credit_card_expired_year,
			CreditCardCvv:credit_card_cvv,
			Status:status,
		}
		users = append(users, &user)
	}

	return users, err
}