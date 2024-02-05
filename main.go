package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/way11229/tinder_matching/domain"
	handlerGrpc "github.com/way11229/tinder_matching/handler/grpc"
	pb "github.com/way11229/tinder_matching/pb"
	memDB "github.com/way11229/tinder_matching/repo/mem_db"
	"github.com/way11229/tinder_matching/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	GRPC_PORT = 9000
	HTTP_PORT = 8080
)

func main() {
	serviceManager := newServiceManager()
	newServer(serviceManager)
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

func newServer(
	serviceManager *domain.ServiceManager,
) {
	grpcHandler := handlerGrpc.NewGrpcHandler(
		serviceManager.UserService,
	)

	go newGatewayServer(grpcHandler)

	newGrpcServer(grpcHandler)
}

func newGatewayServer(
	grpcHandler *handlerGrpc.GrpcHandler,
) {
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := pb.RegisterUserServiceHandlerServer(ctx, grpcMux, grpcHandler); err != nil {
		panic(err)
	}

	handler := grpcHandler.PanicHandler(grpcMux)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", HTTP_PORT))
	if err != nil {
		panic(err)
	}

	log.Printf("http server start to listen port: %d\n", HTTP_PORT)
	if err := http.Serve(lis, handler); err != nil {
		panic(err)
	}
}

func newGrpcServer(
	grpcHandler *handlerGrpc.GrpcHandler,
) {
	sigCh := make(chan os.Signal, 1)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", GRPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	pb.RegisterUserServiceServer(
		grpcServer,
		grpcHandler,
	)

	reflection.Register(grpcServer)

	log.Printf("grpc server start to listen port: %d\n", GRPC_PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	grpcServer.GracefulStop()

	log.Println("grpc server graceful shutdown End")
}
