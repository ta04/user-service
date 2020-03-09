package main

import (
	"fmt"
	"log"

	"github.com/G0tYou/user-service/database"
	"github.com/G0tYou/user-service/handler"
	userPB "github.com/G0tYou/user-service/proto"
	"github.com/G0tYou/user-service/repository/postgres"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
)

func main() {
	// Setup the micro instance
	s := micro.NewService(
		micro.Name("com.ta04.srv.user"),
	)
	// Initialize the service
	s.Init()

	// Connect to Postgres
	db, err := database.OpenPostgresConnection()
	if err != nil {
		log.Fatalf("failed to connect to postgress: #{err}")
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create a new handler
	h := handler.NewHandler(&postgres.Repository{
		DB: db,
	})

	// Register handler
	userPB.RegisterUserServiceServer(s.Server(), h)

	// Run the server
	err = s.Run()
	if err != nil {
		fmt.Println(err)
	}
}
