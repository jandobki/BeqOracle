package server

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/jandobki/beqoracle/server/internal/model"
)

type beqOracleServer struct {
	pb.UnimplementedBeqOracleServer
}

func NewServer() *beqOracleServer {
	return &beqOracleServer{}
}

func (s *beqOracleServer) CreateAnswer(context.Context, *pb.CreateAnswerRequest) (*pb.Answer, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *beqOracleServer) UpdateAnswer(context.Context, *pb.UpdateAnswerRequest) (*pb.Answer, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *beqOracleServer) GetAnswer(context.Context, *pb.GetAnswerRequest) (*pb.Answer, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *beqOracleServer) DeleteAnswer(context.Context, *pb.DeleteAnswerRequest) (*empty.Empty, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *beqOracleServer) GetAnswerHistory(context.Context, *pb.GetAnswerHistoryRequest) (*pb.EventList, error) {
	return nil, fmt.Errorf("not implemented")
}
