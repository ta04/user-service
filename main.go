// user-service/main.go

package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	"github.com/G0tYou/user-service/database"
	"github.com/G0tYou/user-service/handler"
	userPB "github.com/G0tYou/user-service/proto"
	"github.com/G0tYou/user-service/repository/postgres"
	authPB "github.com/SleepingNext/auth-service/proto"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
)

var methodsWithoutAuth = map[string]bool{"UserService.ShowUser": true, "UserService.ShowUserByUsername": true, "UserService.StoreUser": true}

func main() {
	// Setup the micro instance
	s := micro.NewService(
		micro.Name("com.ta04.srv.user"),
		micro.WrapHandler(AuthWrapper),
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
	userPB.RegisterUserServiceHandler(s.Server(), h)

	// Run the service
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, res interface{}) error {
		method := req.Method()
		if _, ok := methodsWithoutAuth[method]; ok {
			return fn(ctx, req, res)
		}

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in the request")
		}

		token := meta["Token"]
		log.Println("authenticating with token: ", token)

		// Validate the token
		authClient := authPB.NewAuthServiceClient("com.ta04.srv.auth", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &authPB.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, res)
		return err
	}
}
