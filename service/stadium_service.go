package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	pb "stadion/genproto/stadium"
	"stadion/storage"
	"stadion/storage/postgres"
)

type StadiumService struct {
	pb.UnimplementedStadiumServiceServer
	Stadium storage.IStadiumStorage
	Log     *slog.Logger
}

func NewStadiumService(db *sql.DB, log *slog.Logger) *StadiumService {
	return &StadiumService{
		Stadium: postgres.NewStadiumRepository(db),
		Log:     log,
	}
}

func (s *StadiumService) CreateStadium(ctx context.Context, req *pb.CreateStadiumRequest) (*pb.CreateStadiumResponse, error) {
	res, err := s.Stadium.CreateStadium(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error creating stadium service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) UpdateStadium(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	res, err := s.Stadium.UpdateStadium(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error updating stadium service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) GetStadium(ctx context.Context, req *pb.GetStadiumRequest) (*pb.GetStadiumResponse, error) {
	res, err := s.Stadium.GetStadium(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting stadium service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) GetStadiums(ctx context.Context, req *pb.GetStadiumsRequest) (*pb.GetStadiumsResponse, error) {	
	res, err := s.Stadium.GetStadiums(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting stadiums service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) DeleteStadium(ctx context.Context, req *pb.DeleteStadiumRequest) (*pb.DeleteStadiumResponse, error) {	
	res, err := s.Stadium.DeleteStadium(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error deleting stadium service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) CreateOrderStadium(ctx context.Context, req *pb.CreateOrderStadiumRequest) (*pb.CreateOrderStadiumResponse, error) {
	res, err := s.Stadium.CreateOrderStadium(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error creating order service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) GetOrderStadiums(ctx context.Context, req *pb.GetOrderStadiumsRequest) (*pb.GetOrderStadiumsResponse, error) {
	res, err := s.Stadium.GetOrderStadiums(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting orders service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) GetOrderStadium(ctx context.Context, req *pb.GetOrderStadiumRequest) (*pb.GetOrderStadiumResponse, error) {
	res, err := s.Stadium.GetOrderStadium(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting order service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) UpdateOrderStadium(ctx context.Context, req *pb.UpdateOrderStadiumRequest) (*pb.UpdateOrderStadiumResponse, error) {
	res, err := s.Stadium.UpdateOrderStadium(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error updating order service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) DeleteOrderStadium(ctx context.Context, req *pb.DeleteOrderStadiumRequest) (*pb.DeleteOrderStadiumResponse, error) {
	res, err := s.Stadium.DeleteOrderStadium(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error deleting order service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (s *StadiumService) GetDeletedOrderStadiums(ctx context.Context, req *pb.GetDeletedOrderStadiumsRequest) (*pb.GetDeletedOrderStadiumsResponse, error) {
	res, err := s.Stadium.GetDeletedOrderStadiums(ctx, req)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error getting deleted orders service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}