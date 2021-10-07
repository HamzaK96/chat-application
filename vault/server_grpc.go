package vault

import (
	"github.com/HamzaK96/cloudplex/training/code/vault/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	hash     grpctransport.Handler
	validate grpctransport.Handler
}

func (s *grpcServer) Hash(ctx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.HashResponse), nil
}

func (s *grpcServer) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.ValidateResponse), nil
}

func EncodeGRPCHashRequest(ctx context.Context,
	r interface{}) (interface{}, error) {
	req := r.(hashRequest)
	return &pb.HashRequest{Password: req.Password}, nil
}
