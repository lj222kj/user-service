package main

import (
	"fmt"
	"log"
	"net"
	"user-service/database"
	"user-service/global"
	"user-service/proto"
	"user-service/users"
	v2 "user-service/users/v2"
	"user-service/web"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := global.New()
	if err != nil {
		fmt.Errorf("failed to read env config")
	}
	db := database.New()
	userSvc := users.New(db)
	userSvcV2 := v2.New(db) // ignore package name, it's a quick hack to get both up.
	api := web.New(cfg, userSvc)

	grpcConn, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		log.Fatalf("failed to setup grpc at port: %s", cfg.GrpcPort)
	}
	defer grpcConn.Close()
	srv := grpc.NewServer()
	proto.RegisterUserSummaryServer(srv, userSvcV2)
	go func() {
		fmt.Printf("rest api running on: %s\n", cfg.Port)

		err = api.ListenAndServe(); if err != nil {
			fmt.Errorf("failed to listen to port", cfg.Port)
		}
	}()

	fmt.Printf("grpc server running on: %s\n", cfg.GrpcPort)
	if err = srv.Serve(grpcConn); err != nil {
		log.Fatalf("failed to startup grpc server at port %s", cfg.GrpcPort)
	}
}