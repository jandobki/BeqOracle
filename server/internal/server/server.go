package server

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/jandobki/beqoracle/server/internal/model"
	"github.com/jandobki/beqoracle/server/internal/oracle"
)

type beqOracleServer struct {
	pb.UnimplementedBeqOracleServer
	service *oracle.Service
}

func NewServer() *beqOracleServer {
	return &beqOracleServer{
		service: oracle.NewService(),
	}
}

func (s *beqOracleServer) CreateAnswer(ctx context.Context, req *pb.CreateAnswerRequest) (*pb.Answer, error) {
	err := s.service.CreateAnswer(ctx, req.Key, req.Value)
	if err != nil {
		return nil, err
	}

	return &pb.Answer{
		Key:   req.Key,
		Value: req.Value,
	}, nil
}

func (s *beqOracleServer) UpdateAnswer(ctx context.Context, req *pb.UpdateAnswerRequest) (*pb.Answer, error) {
	err := s.service.UpdateAnswer(ctx, req.Key, req.Value)
	if err != nil {
		return nil, err
	}

	return &pb.Answer{
		Key:   req.Key,
		Value: req.Value,
	}, nil
}

func (s *beqOracleServer) GetAnswer(ctx context.Context, req *pb.GetAnswerRequest) (*pb.Answer, error) {
	a, err := s.service.GetAnswer(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	return &pb.Answer{
		Key:   req.Key,
		Value: a,
	}, nil
}

func (s *beqOracleServer) DeleteAnswer(ctx context.Context, req *pb.DeleteAnswerRequest) (*empty.Empty, error) {
	err := s.service.DeleteAnswer(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *beqOracleServer) GetAnswerHistory(context.Context, *pb.GetAnswerHistoryRequest) (*pb.EventList, error) {
	return nil, fmt.Errorf("not implemented")
}
