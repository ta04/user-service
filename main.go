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
	GetAll() ([]*pb.User, error)
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
		err := rows.Scan(&id, &first_name, &last_name, &username, &password, &email_address, &phone_number, &address, &role, &credit_card_number, &credit_card_type, &credit_card_expired_month, &credit_card_expired_year, &credit_card_cvv, &status)
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

func (repo *Repository) GetOne(user *pb.User)(*pb.User, error){
	var id int32
	var first_name, last_name, username, password, email_address, phone_number, address, role, credit_card_number, credit_card_type, credit_card_expired_month, credit_card_expired_year,credit_card_cvv string
	var status bool
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d",user.Id)

	err := repo.db.QueryRow(query).Scan(&id, &first_name, &last_name, &username, &password, &email_address, &phone_number, &address, &role, &credit_card_number, &credit_card_type, &credit_card_expired_month, &credit_card_expired_year, &credit_card_cvv, &status)
	if err != nil{
		return  nil, err
	}
	return &pb.User{
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
		}, err
}

func (repo *Repository) Update(user *pb.User)(*pb.User, error){
	repo.mu.Lock()
	query := fmt.Sprintf("UPDATE users SET first_name=%s, last_name=%s, username=%s, password=%s, email_address=%s, phone_number=%s, date_of_birth=%s, address=%s, role=%s, credit_card_number=%s, credit_card_type=%s, credit_card_expired_month=%s, credit_card_expired_year=%s, credit_card_cvv=%s, status=%t" +
		"WHERE id=%d", user.FirstName, user.LastName, user.Username, user.Password, user.EmailAddress, user.PhoneNumber, user.DateOfBirth, user.Address, user.Role, user.CreditCardNumber, user.CreditCardType, user.CreditCardExpiredMonth, user.CreditCardExpiredYear, user.CreditCardCvv, user.Status, user.Id)
	_, err:= repo.db.Exec(query)
	repo.mu.Unlock()

	return user, err
}

func (repo *Repository) Delete(user *pb.User) (*pb.User, error){
	repo.mu.Lock()
	query := fmt.Sprintf("DELETE FROM users WHERE id =#{users.Id}")
	_, err:=repo.db.Exec(query)
	repo.mu.Unlock()

	return user,err
}

type service struct{
	repo repository
}

func(s *service) CreateUser(ctx context.Context, req *pb.User) (*pb.Response, error){
	user, err := s.repo.Create(req)
	if err != nil{
		return nil, err
	}

	return &pb.Response{
		User:                 user,
		Error:                nil,
	},err
}

func(s *service) GetAllUser(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.Response, error){
	users, err := s.repo.GetAll()

	return &pb.Response{
		Users: users,
	}, err
}

func(s *service) GetOneUser(ctx context.Context, req *pb.User) (*pb.Response, error){
	user, err := s.repo.GetOne(req)

	return &pb.Response{
		User: user,
	}, err
}

func(s *service) UpdateUser(ctx context.Context, req *pb.User) (*pb.Response, error){
	user, err := s.repo.Update(req)
	if err != nil{
		return nil, err
	}

	return &pb.Response{
		User:                 user,
		Error:                nil,
	},err
}

func(s *service) DeleteUser(ctx context.Context, req *pb.User) (*pb.Response, error){
	user, err := s.repo.Delete(req)
	if err != nil{
		return nil, err
	}

	return &pb.Response{
		User:                 user,
		Error:                nil,
	},err
}

func connectPostgres() (*sql.DB, error) {
	connStr := "user =silly-wizard dbname=users sslmode=disable password=asyouwish"
	db, err:= sql.Open("postgres", connStr)

	return db, err
}

func main() {
	db, err:= connectPostgres()
	repo := &Repository{}
	repo.db = db

	if err != nil{
		log.Fatalf("failed to connect to postgress: #{err}")
	}

	listen, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("fatiled to listen: #{err}")
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &service{repo})

	reflection.Register(s)
	log.Println("Running on port", port)
	if err != s.Serve(listen) {
		log.Fatalf("failed to serve: #{err}")
	}
}