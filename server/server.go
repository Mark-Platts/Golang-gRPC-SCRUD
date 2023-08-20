package server

import (
	"context"
	"log"

	"markplatts.org/scrud/models"
	pb "markplatts.org/scrud/protos"
	"markplatts.org/scrud/storage"
)

//set up a struct that will implement the methods defined in the proto for the service, then implement methods below
type ScrudServer struct {
	pb.UnimplementedScrudServer
}

func (s *ScrudServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	log.Printf("Create RPC called with request: %v", req)

	//use the server request to build a storage message and send it to storage layer
	//the storage layer sends back an id to give to the server reply
	ms := models.NewMessageStore(req.GetMessage())
	id, err := storage.StorageMechanism.Insert(ms)
	if err != nil {
		return nil, err
	}

	log.Printf("Local message store state: %v", storage.StorageMechanism.GetMessages()) //for testing/debugging TODO: remove when implementing mySQL
	return &pb.CreateReply{Id: id}, nil
}

func (s *ScrudServer) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadReply, error) {
	log.Printf("Create RPC called with request: %v", req)

	//use the service request's id to retrieve the stored message from the storage layer
	ms, err := storage.StorageMechanism.Retrieve(req.GetId())
	if err != nil {
		return nil, err
	}	

	return &pb.ReadReply{Id: req.GetId(), TimeCreated: ms.GetTimeCreated(), Message: ms.GetMessage()}, nil
}

func (s *ScrudServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateReply, error) {
	log.Printf("Update RPC called with request: %v", req)
	return &pb.UpdateReply{}, nil
}

func (s *ScrudServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	log.Printf("Delete RPC called with request: %v", req)
	return &pb.DeleteReply{}, nil
}

func (s *ScrudServer) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendReply, error) {
	log.Printf("Send RPC called with request: %v", req)
	return &pb.SendReply{}, nil
}

func (s *ScrudServer) SendAll(req *pb.SendAllRequest, serv pb.Scrud_SendAllServer) error {
	log.Printf("SendAll RPC called with request: %v", req)
	return nil
}