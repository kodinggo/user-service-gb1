package console

import (
	"log"
	"net"
	"sync"

	"github.com/kodinggo/user-service-gb1/db"
	"github.com/kodinggo/user-service-gb1/internal/config"
	grpcSvc "github.com/kodinggo/user-service-gb1/internal/delivery/grpc"
	handlerHttp "github.com/kodinggo/user-service-gb1/internal/delivery/http"
	"github.com/kodinggo/user-service-gb1/internal/middleware"
	"github.com/kodinggo/user-service-gb1/internal/repository"
	"github.com/kodinggo/user-service-gb1/internal/usecase"
	pb "github.com/kodinggo/user-service-gb1/pb/auth"
	"google.golang.org/grpc"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "httpsrv",
	Short: "Start the HTTP server",
	Run:   httpServer,
}

func httpServer(cmd *cobra.Command, args []string) {
	// Get env variables from .env file
	config.LoadWithViper()
	cfg := config.Cfg

	mysql := db.NewMysql()
	defer mysql.Close()

	// redis := db.NewRedis()

	userRepo := repository.NewUserRepository(mysql)
	roleRepo := repository.NewRoleRepository(mysql)

	userUsecase := usecase.NewUserUsecase(userRepo, roleRepo)
	authUsecase := usecase.NewAuthUsecase(cfg.JWT.Secret, userRepo, roleRepo)

	authGrpc := grpcSvc.NewJWTValidatorServer(cfg.JWT.Secret, authUsecase)

	grpcServer := grpc.NewServer()
	pb.RegisterJWTValidatorServer(grpcServer, authGrpc)

	JWTMiddleware := middleware.NewJWTMiddleware(authUsecase)

	// Create a new Echo instance
	e := echo.New()
	// e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Logger())

	routeGroup := e.Group("/api/v1")

	handlerHttp.NewUserHandler(routeGroup, userUsecase, JWTMiddleware.ValidateJWT)

	grpcPort := "4000" // todo make it env
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Panicf("failed create http listener, error: %v", err)
	}
	log.Println("Starting gRPC server at " + grpcPort)

	var wg sync.WaitGroup
	errCh := make(chan error, 2)
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := e.Start(":3200")
		if err != nil {
			errCh <- err
		}
	}()

	go func() {
		defer wg.Done()
		err := grpcServer.Serve(lis)
		if err != nil {
			errCh <- err
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			logrus.Error(err.Error())
		}
	}
}
