package server

import (
	"context"
	pb "github.com/novikov-ai/rate-limiter/api/pb"
	"github.com/novikov-ai/rate-limiter/configs"
	"github.com/novikov-ai/rate-limiter/internal/storage"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	pb.UnimplementedLimiterServer
	storage storage.KeyValue
}

func Start(ctx context.Context, st storage.KeyValue, conf configs.Config) error {
	lsn, err := net.Listen("tcp", conf.Server.Host+":"+conf.Server.Port)
	if err != nil {
		return err
	}

	if err = st.Connect(ctx); err != nil {
		return err
	}
	defer st.Close()

	server := grpc.NewServer()

	//server := grpc.NewServer(
	//	grpc.ChainUnaryInterceptor(UnaryServerRequestValidatorInterceptor(ValidateReq),
	//		UnaryServerLoggingInterceptor(logger)),
	//)

	service := &Service{storage: st}

	pb.RegisterLimiterServer(server, service)

	if err = server.Serve(lsn); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Service) DetectBruteforce(ctx context.Context, request *pb.DetectBruteforceRequest) (*pb.DetectBruteforceResponse, error) {
	response := &pb.DetectBruteforceResponse{Detected: false}

	select {
	case <-ctx.Done():
		break
	default:
		ip := request.Ip

		foundAtWhiteList, err := s.storage.FindAtWhiteList(ip)
		if err != nil {
			return nil, err
		}
		if foundAtWhiteList {
			break
		}

		foundAtBlackList, err := s.storage.FindAtBlackList(ip)
		if err != nil {
			return nil, err
		}
		if foundAtBlackList {
			response.Detected = true
			break
		}

		overflowedByLogins, err := s.storage.OverflowAttemptsLogin(request.Login)
		if err != nil {
			return nil, err
		}

		overflowedByPasswords, err := s.storage.OverflowAttemptsPasswords(request.Password)
		if err != nil {
			return nil, err
		}

		tooManyRequests := overflowedByLogins || overflowedByPasswords

		if tooManyRequests {
			response.Detected = true
			break
		}
	}

	return response, nil
}

func (s *Service) ResetBucket(ctx context.Context, request *pb.ResetBucketRequest) (*pb.ResetBucketResponse, error) {
	response := &pb.ResetBucketResponse{Status: 1}

	select {
	case <-ctx.Done():
		break
	default:
		err := s.storage.RemoveAllLoginsAttempts(request.Logins)
		if err != nil {
			return response, err
		}

		err = s.storage.RemoveAllAddressesAttempts(request.Ips)
		if err != nil {
			return response, err
		}

		response.Status = 0
	}

	return response, nil
}

func (s *Service) WhiteListAddIP(ctx context.Context, request *pb.ManageAddressRequest) (*pb.ManageAddressResponse, error) {
	response := &pb.ManageAddressResponse{Status: 1}

	select {
	case <-ctx.Done():
		break
	default:
		// todo: ip-mask
		err := s.storage.Add(storage.WhiteList, request.Ip, request.Mask)
		if err != nil {
			return response, err
		}

		response.Status = 0
	}

	return response, nil
}

func (s *Service) WhiteListRemoveIP(ctx context.Context, request *pb.ManageAddressRequest) (*pb.ManageAddressResponse, error) {
	response := &pb.ManageAddressResponse{Status: 1}

	select {
	case <-ctx.Done():
		break
	default:
		// todo: ip-mask
		err := s.storage.Remove(storage.WhiteList, request.Ip+request.Mask)
		if err != nil {
			return response, err
		}

		response.Status = 0
	}

	return response, nil
}

func (s *Service) BlackListAddIP(ctx context.Context, request *pb.ManageAddressRequest) (*pb.ManageAddressResponse, error) {
	response := &pb.ManageAddressResponse{Status: 1}

	select {
	case <-ctx.Done():
		break
	default:
		// todo: ip-mask
		err := s.storage.Add(storage.BlackList, request.Ip, request.Mask)
		if err != nil {
			return response, err
		}

		response.Status = 0
	}

	return response, nil
}

func (s *Service) BlackListRemoveIP(ctx context.Context, request *pb.ManageAddressRequest) (*pb.ManageAddressResponse, error) {
	response := &pb.ManageAddressResponse{Status: 1}

	select {
	case <-ctx.Done():
		break
	default:
		// todo: ip-mask
		err := s.storage.Remove(storage.BlackList, request.Ip+request.Mask)
		if err != nil {
			return response, err
		}

		response.Status = 0
	}

	return response, nil
}
