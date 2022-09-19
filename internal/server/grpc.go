package server

import (
	"context"
	pb "github.com/novikov-ai/rate-limiter/api/pb"
	"github.com/novikov-ai/rate-limiter/configs"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	pb.UnimplementedLimiterServer
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

	pb.RegisterLimiterServer(server, service)

	if err = server.Serve(lsn); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Service) DetectBruteforce(ctx context.Context, request *pb.DetectBruteforceRequest) (*pb.DetectBruteforceResponse, error) {
	return &pb.DetectBruteforceResponse{Detected: false}, nil
}

func (s *Service) ResetBucket(ctx context.Context, request *pb.ResetBucketRequest) (*pb.ResetBucketResponse, error) {
	return &pb.ResetBucketResponse{Status: 0}, nil
}

func (s *Service) WhiteListAddIP(ctx context.Context, request *pb.ManageAddressRequest) (*pb.ManageAddressResponse, error) {
	return &pb.ManageAddressResponse{Status: 0}, nil
}

func (s *Service) WhiteListRemoveIP(ctx context.Context, request *pb.ManageAddressRequest) (*pb.ManageAddressResponse, error) {
	return &pb.ManageAddressResponse{Status: 0}, nil
}

func (s *Service) BlackListAddIP(ctx context.Context, request *pb.ManageAddressRequest) (*pb.ManageAddressResponse, error) {
	return &pb.ManageAddressResponse{Status: 0}, nil
}

func (s *Service) BlackListRemoveIP(ctx context.Context, request *pb.ManageAddressRequest) (*pb.ManageAddressResponse, error) {
	return &pb.ManageAddressResponse{Status: 0}, nil
}
