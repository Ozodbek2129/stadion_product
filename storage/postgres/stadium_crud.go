package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	"stadion/pkg/logger"
	"stadion/storage"
	pb "stadion/genproto/stadium"
)

type StadiumRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewStadiumRepository(db *sql.DB) storage.IStadiumStorage {
	return &StadiumRepository{Db: db, Log: logger.NewLogger()}
}

func (s *StadiumRepository) CreateStadium(ctx context.Context, req *pb.CreateStadiumRequest) (*pb.CreateStadiumResponse, error) {
	return nil, nil
}

func (s *StadiumRepository) UpdateStadium(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return nil, nil
}

func (s *StadiumRepository) GetStadium(ctx context.Context, req *pb.GetStadiumRequest) (*pb.GetStadiumResponse, error) {
	return nil, nil
}

func (s *StadiumRepository) GetStadiums(ctx context.Context, req *pb.GetStadiumsRequest) (*pb.GetStadiumsResponse, error) {
	return nil, nil
}

func (s *StadiumRepository) DeleteStadium(ctx context.Context, req *pb.DeleteStadiumRequest) (*pb.DeleteStadiumResponse, error) {
	return nil, nil
}

func (s *StadiumRepository) OrderStadium(ctx context.Context, req *pb.OrderStadiumRequest) (*pb.OrderStadiumResponse, error) {
	return nil, nil
}
