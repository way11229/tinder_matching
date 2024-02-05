package test

import (
	"context"
	"log"
	"net"

	"github.com/way11229/tinder_matching/domain"
	handlerGrpc "github.com/way11229/tinder_matching/handler/grpc"
	pb "github.com/way11229/tinder_matching/pb"
	memDB "github.com/way11229/tinder_matching/repo/mem_db"
	"github.com/way11229/tinder_matching/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (pb.UserServiceClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterUserServiceServer(baseServer, newGrpcHandler())
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	return pb.NewUserServiceClient(conn), closer
}

func newGrpcHandler() *handlerGrpc.GrpcHandler {
	serviceManager := newServiceManager()

	return handlerGrpc.NewGrpcHandler(
		serviceManager.UserService,
	)
}

func newServiceManager() *domain.ServiceManager {
	db := memDB.NewMemDB()
	usersMaleMemDB := memDB.NewUsersMaleMemDB(db)
	usersFemaleMemDB := memDB.NewUsersFemaleMemDB(db)

	return &domain.ServiceManager{
		UserService: service.NewUserService(
			usersMaleMemDB,
			usersFemaleMemDB,
		),
	}
}
