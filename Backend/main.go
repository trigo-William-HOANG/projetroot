// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/trigo-William-HOANG/projectroot/backend/invoicer"
)

// server is used to implement the gRPC service.
type server struct {
	pb.UnimplementedInvoicerServer
	db *pgx.Conn
}

// App represents the application details.
type App struct {
	IdApp          string `json:"idApp"`
	NameApp        string `json:"nameApp"`
	DescriptionApp string `json:"descriptionApp"`
	AppUrl         string `json:"appUrl"`
	ImageAppUrl    string `json:"imageAppUrl"`
}

// Connect to PostgreSQL database
func connectToDB() (*pgx.Conn, error) {
	connStr := "postgres://username:password@localhost:5432/mydb?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return conn, nil
}

// InvoiceResponseJSON is a custom struct to be returned as JSON.
type InvoiceResponseJSON struct {
	Id       string  `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Paid     bool    `json:"paid"`
}

// GetInvoice method will return the invoice in JSON format.
func (s *server) GetInvoice(ctx context.Context, in *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	// Query to get apps data
	rows, err := s.db.Query(ctx, "SELECT id_app, name_app, description_app, app_url, image_app_url FROM apps LIMIT 3")
	if err != nil {
		log.Printf("Error fetching data from database: %v", err)
		return nil, fmt.Errorf("error fetching data from database: %v", err)
	}
	defer rows.Close()

	// Create slice to hold app data
	var apps []*pb.App
	for rows.Next() {
		var app pb.App
		err := rows.Scan(&app.IdApp, &app.NameApp, &app.DescriptionApp, &app.AppUrl, &app.ImageAppUrl)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		apps = append(apps, &app)
	}

	// Log the apps data that will be sent
	log.Printf("Apps fetched: %v", apps)

	// Return the structured response with multiple apps
	return &pb.InvoiceResponse{
		Apps: apps,
	}, nil
}

func main() {
	// Connect to PostgreSQL
	db, err := connectToDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close(context.Background())

	// Listen on TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterInvoicerServer(s, &server{db: db})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
