package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/gurodrigues-dev/account-gRPC/config"
	"github.com/gurodrigues-dev/account-gRPC/internal/controllers"
	"github.com/gurodrigues-dev/account-gRPC/internal/repository"
	"github.com/gurodrigues-dev/account-gRPC/internal/service"
	"github.com/gurodrigues-dev/account-gRPC/pb"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = migrate(db, config.Database.Schema)
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	_ = controllers.NewAccountController(service.NewAccountService(repository.NewAccountRepository((db))))

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterAccountServer(s, &controllers.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("initing service: %s", config.Name)

}

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func migrate(db *sql.DB, filepath string) error {
	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}
