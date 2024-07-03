package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"unsia/controllers"
	"unsia/pb/cities"
	"unsia/pkg/database"

	// Import driver PostgreSQL
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {
	log := log.New(os.Stdout, "CRUD-go : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	  // Start Database

	  db, err := database.OpenDB()
	  if err != nil {
		  log.Fatalf("error: connecting to db: %s", err)
	  }
	  defer db.Close()

	grpcServer := grpc.NewServer()

	cityServer := controllers.City{DB: db, Log: log}
	cities.RegisterCitiesServiceServer(grpcServer, &cityServer)

	fmt.Println("running server grpc")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %s", err)
		return
	}
}
