package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/jandobki/beqoracle/server/internal/model"
	"github.com/jandobki/beqoracle/server/internal/oracle"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type beqOracleServer struct {
	pb.UnimplementedBeqOracleServer
	service *oracle.Service
}

func NewServer(ctx context.Context) *beqOracleServer {
	return &beqOracleServer{
		service: oracle.NewService(),
	}
}

func (s *beqOracleServer) CreateAnswer(ctx context.Context, req *pb.CreateAnswerRequest) (*pb.Answer, error) {
	err := s.service.CreateAnswer(ctx, req.Key, req.Value)
	if err != nil {
		return nil, mapError(err)
	}

	return &pb.Answer{
		Key:   req.Key,
		Value: req.Value,
	}, nil
}

func (s *beqOracleServer) UpdateAnswer(ctx context.Context, req *pb.UpdateAnswerRequest) (*pb.Answer, error) {
	err := s.service.UpdateAnswer(ctx, req.Key, req.Value)
	if err != nil {
		return nil, mapError(err)
	}

	return &pb.Answer{
		Key:   req.Key,
		Value: req.Value,
	}, nil
}

func (s *beqOracleServer) GetAnswer(ctx context.Context, req *pb.GetAnswerRequest) (*pb.Answer, error) {
	a, err := s.service.GetAnswer(ctx, req.Key)
	if err != nil {
		return nil, mapError(err)
	}

	return &pb.Answer{
		Key:   req.Key,
		Value: a,
	}, nil
}

func (s *beqOracleServer) DeleteAnswer(ctx context.Context, req *pb.DeleteAnswerRequest) (*empty.Empty, error) {
	err := s.service.DeleteAnswer(ctx, req.Key)
	if err != nil {
		return nil, mapError(err)
	}

	return &empty.Empty{}, nil
}

func (s *beqOracleServer) ListEvents(ctx context.Context, req *pb.ListEventsRequest) (*pb.EventList, error) {
	evs, to, err := s.service.ListEvents(ctx, req.Key, int(req.PageToken), int(req.PageSize))
	if err != nil {
		return nil, mapError(err)
	}

	res := make([]*pb.Event, len(evs))
	for i, e := range evs {
		res[i] = &pb.Event{
			Event: e.Event,
			Data:  &pb.Answer{Key: req.Key, Value: e.Value},
		}
	}

	return &pb.EventList{
		Events:        res,
		NextPageToken: int32(to),
	}, nil
}

func mapError(err error) error {
	switch err {
	case nil:
		return status.Error(codes.OK, codes.OK.String())
	case oracle.ErrAlreadyExists:
		return status.Error(codes.AlreadyExists, err.Error())
	case oracle.ErrNotFound:
		return status.Error(codes.NotFound, err.Error())
	default:
		return status.Error(codes.Unavailable, err.Error())
	}
}
