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
	CreateOrderStadium (context.Context, *pb.CreateOrderStadiumRequest) (*pb.CreateOrderStadiumResponse, error)		
	GetOrderStadiums (context.Context, *pb.GetOrderStadiumsRequest) (*pb.GetOrderStadiumsResponse, error)
	GetOrderStadium (context.Context, *pb.GetOrderStadiumRequest) (*pb.GetOrderStadiumResponse, error)	
	UpdateOrderStadium (context.Context, *pb.UpdateOrderStadiumRequest) (*pb.UpdateOrderStadiumResponse, error)
	DeleteOrderStadium (context.Context, *pb.DeleteOrderStadiumRequest) (*pb.DeleteOrderStadiumResponse, error)
	GetDeletedOrderStadiums (context.Context, *pb.GetDeletedOrderStadiumsRequest) (*pb.GetDeletedOrderStadiumsResponse, error)	
	GetAllStadium (context.Context, *pb.GetAllStadiumRequest) (*pb.GetAllStadiumResponse, error)							
}
