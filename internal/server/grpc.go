package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"rate-limiter/configs"
)

//import "github.com/novikov-ai/practice-misc/"
//  pb "github.com/novikov-ai/practice-misc/hw12_13_14_15_calendar/internal/server/pb"

type Service struct {
	//  pb.UnimplementedCalendarServer
}

func Start(ctx context.Context, conf configs.Config) error {
	lsn, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	//server := grpc.NewServer(
	//	grpc.ChainUnaryInterceptor(UnaryServerRequestValidatorInterceptor(ValidateReq),
	//		UnaryServerLoggingInterceptor(logger)),
	//)

	service := new(Service)
	fmt.Println(service)

	// pb.RegisterCalendarServer(server, service)

	if err = server.Serve(lsn); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

//func (s *Service) DetectBruteforce(ctx context.Context, request *pb.)  (*pb.AddEventResponse, error) {
