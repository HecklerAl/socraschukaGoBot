package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	lm "github.com/HecklerAl/socraschukaGoBot/internal/link_modifier"
	nsql "github.com/HecklerAl/socraschukaGoBot/internal/nosql"
	pb "github.com/HecklerAl/socraschukaGoBot/pkg"
)

const port = ":8082"

type server struct {
	pb.UnimplementedLinkServiceServer
}

func (s *server) ShortenLink(ctx context.Context, req *pb.ModifyLinkRequest) (*pb.ModifyLinkResponse, error) {
	log.Printf("ShortenLink: %+v", req)

	short, err := lm.ShortenLink(req.ActualLink)
	if err != nil {
		errStr := err.Error()
		return &pb.ModifyLinkResponse{ModifiedLink: "", Error: &errStr}, nil
	}

	link := lm.ConvertURL(*short)
	nsql.Add(*link, req.ActualLink, "short")
	return &pb.ModifyLinkResponse{ModifiedLink: *link}, nil
}

func (s *server) LengthenLink(ctx context.Context, req *pb.ModifyLinkRequest) (*pb.ModifyLinkResponse, error) {
	log.Printf("LengthenLink: %+v", req)

	long, err := lm.LengthenLink(req.ActualLink)
	if err != nil {
		errStr := err.Error()
		return &pb.ModifyLinkResponse{ModifiedLink: "", Error: &errStr}, nil
	}

	link := lm.ConvertURL(*long)
	nsql.Add(*link, req.ActualLink, "long")
	return &pb.ModifyLinkResponse{ModifiedLink: *link}, nil
}

func main() {
	if err := nsql.LoadFromDB(); err != nil {
		log.Fatalf("Ошибка загрузки db.json: %v", err)
	}
	defer func() {
		if err := nsql.SaveData(); err != nil {
			log.Fatalf("Ошибка сохранения db.json: %v", err)
		}
	}()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLinkServiceServer(grpcServer, &server{})
	log.Printf("gRPC-сервер запущен на %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка при запуске gRPC: %v", err)
	}
}
