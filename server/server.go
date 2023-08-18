package server

import (
	"context"
	pb "markplatts.org/scrud/protos"
)



type ScrudServer struct {
	pb.UnimplementedScrudServer
}

func (s *ScrudServer) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateReply, error) {
	println("Create RPC called")
	return &pb.CreateReply{}, nil
}

func (s *ScrudServer) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadReply, error) {
	println("Read RPC called")
	return &pb.ReadReply{}, nil
}

func (s *ScrudServer) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateReply, error) {
	println("Update RPC called")
	return &pb.UpdateReply{}, nil
}

func (s *ScrudServer) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteReply, error) {
	println("Delete RPC called")
	return &pb.DeleteReply{}, nil
}

func (s *ScrudServer) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	println("Send RPC called")
	return &pb.SendReply{}, nil
}

func (s *ScrudServer) SendAll(in *pb.SendAllRequest, serv pb.Scrud_SendAllServer) error {
	println("SendAll RPC called")
	return nil
}