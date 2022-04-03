package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc_go_blog/services/user/constant"
	"grpc_go_blog/services/user/handler/proto"
)

func RunServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", constant.Port))
	if err != nil {
		log.Printf("failed to run server: %v", err.Error())
		return err
	}

	grpcServer := grpc.NewServer()
	apiServer, err := InjectAPIServer()
	if err != nil {
		log.Println("failed to inject api server")
		return err
	}

	proto.RegisterUserServiceServer(grpcServer, apiServer.UserHandler)

	log.Printf("Listening on %v:", constant.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to run server: %v", err.Error())
		return err
	}
	
	return nil
}