package main

import(
	"context"
	"database/sql"
	"fmt"
	pb "github.com/G0tYouuser-service/proto"
	"google.golang.org/grpc"
	"google.golang.org.grpc/reflection"
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
	repo.mu.lock()
	query := fmt.Sprintf("INSERT INTO users (first_name, last_name, username, password, email_address, phone_number, date_of_birth, address, role, credit_card_number, credit_card_type, credit_card_expired_month, credit_card_expired_year, credit_card_cvv, status)"+
		"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%t')", user.First_name, user.Last_name, user.Username, user.Password, user.Email_address, user.Phone_number, user.Date_of_birth, user.Address, user.Role, user.Credit_card_number, user.Credit_card_type, user.Credit_card_expired_month, user.Credit_card_expired_year, user.Credit_card_cvv, user.Status)
		_, err := repo.db.Exec(query)
		repo.mu.Unlock()

		return user, err
}