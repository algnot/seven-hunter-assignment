package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"seven-hunter-assignment/question-3/logic"
	"seven-hunter-assignment/question-3/proto/gen"
	"seven-hunter-assignment/question-3/util/config"
	"seven-hunter-assignment/question-3/util/request"
)

type server struct {
	proto.BeefServer
}

func (s *server) BeefSummary(ctx context.Context, req *proto.BeefSummaryRequest) (*proto.BeefSummaryResponse, error) {
	beefLogic := logic.GetBeefLogic(&request.BeefAPIImpl{})
	beefs, err := beefLogic.GetBeefSummary()
	if err != nil {
		return nil, err
	}

	var beefResponse []*proto.BeefItem

	for _, beef := range beefs {
		beefResponse = append(beefResponse, &proto.BeefItem{
			Name:  beef.Name,
			Count: beef.Count,
		})
	}

	return &proto.BeefSummaryResponse{
		Beefs: beefResponse,
	}, nil
}

func StartGrpcServer() error {
	appConfig := config.GetAppConfig()
	lis, err := net.Listen("tcp", ":"+appConfig.CommonConfig.GrpcPort)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	proto.RegisterBeefServer(s, &server{})
	fmt.Printf("Server listening on port %s", appConfig.CommonConfig.GrpcPort)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
