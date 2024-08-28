package main

import (
	"log"

	"github.com/kodinggo/auth-service-gb1/pb" // <-- sesuaikan github user-service
	"github.com/kodinggo/service-kalian/middleware"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Initialize gRPC client (replace with actual client initialization code)
	var authClient pb.AuthServiceClient = newAuthClientGRPC()

	// Initialize the JWT middleware using DI
	jwtMiddleware := middleware.NewJWTMiddleware(authClient)

	// Set up Echo server
	e := echo.New()

	// Apply the JWT middleware
	e.Use(jwtMiddleware.ValidateJWT) // <-- pake disini

	// Define routes
	e.GET("/protected", protectedEndpointHandler)

	// Start the server
	e.Start(":3200")
}

func newAuthClientGRPC() pbAuth.AuthServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:3100", opts...)
	if err != nil {
		log.Fatal(err)
	}

	return pbAuth.NewAuthServiceClient(conn)
}
