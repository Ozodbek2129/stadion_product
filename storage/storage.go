package storage

import (
	"context"
	pb "stadion/genproto/stadium"
)

type IStorage interface {
	Stadium() IStadiumStorage
	Close()
}

type IStadiumStorage interface {
	CreateStadium(context.Context, *pb.CreateStadiumRequest) (*pb.CreateStadiumResponse, error)
	UpdateStadium (context.Context, *pb.UpdateRequest) (*pb.UpdateResponse, error)
	GetStadium (context.Context, *pb.GetStadiumRequest) (*pb.GetStadiumResponse, error)
	GetStadiums (context.Context, *pb.GetStadiumsRequest) (*pb.GetStadiumsResponse, error)
	DeleteStadium (context.Context, *pb.DeleteStadiumRequest) (*pb.DeleteStadiumResponse, error)
	OrderStadium (context.Context, *pb.OrderStadiumRequest) (*pb.OrderStadiumResponse, error)											
}
