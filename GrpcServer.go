package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"nrxen.com/dbdriver/protogo"
	"os"
)

var server *grpc.Server = nil

//Server is the configuration for the grpc server
type GrpcConfig struct {
	host string
	port int
}

//An abstraction of the server, allowing to attach the grpc call functions
type GrpcDBServer struct {
}

//override the interface functions (in server .pb.go)
func (s *GrpcDBServer) Get(ctx context.Context, in *dbdriver.GetCmdIn) (*dbdriver.GetCmdOut, error) {
	fmt.Printf("Get Key : %s \n", in.Key)
	str, err := callGet(in.Key)
	if err != nil {
		fmt.Println(err)
		return &dbdriver.GetCmdOut{Key: in.Key, Value: ""}, fmt.Errorf("not found")
	}
	return &dbdriver.GetCmdOut{Key: in.Key, Value: str}, nil
}

//override the interface functions (in server.pb.go)
func (s *GrpcDBServer) Set(ctx context.Context, in *dbdriver.SetCmdIn) (*dbdriver.SetCmdOut, error) {
	fmt.Printf("Set Key: %s , Value : %s \n", in.Key, in.Value)
	return &dbdriver.SetCmdOut{Err: false}, nil
}

func initGrpcServer(config *Configuration) {
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", config.grpc.port))
	if err != nil {
		fmt.Printf("Failed to create a grpc server with port %s ", fmt.Sprintf(":%d", config.grpc.port))
		os.Exit(-1)
	}

	// create a dummy of the abstracted server struct
	s := GrpcDBServer{}

	//create a grpc server, that will listen
	server := grpc.NewServer()

	//register service with the server
	dbdriver.RegisterGrpcDBServer(server, &s)

	//start listening
	if err := server.Serve(list); err != nil {
		fmt.Println("Failed to start grpc server")
		os.Exit(-1)
	}
}

func shutdownGrpcServer() {
	if server != nil {
		server.Stop()
	}
}
