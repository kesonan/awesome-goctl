// Code generated by goctl. DO NOT EDIT!
// Source: greet.proto

package server

import (
	"context"

	"greet/rpc/internal/logic"
	"greet/rpc/internal/svc"
	"greet/rpc/pb"
)

type GreetServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedGreetServer
}

func NewGreetServer(svcCtx *svc.ServiceContext) *GreetServer {
	return &GreetServer{
		svcCtx: svcCtx,
	}
}

func (s *GreetServer) Ping(ctx context.Context, in *pb.Placeholder) (*pb.Placeholder, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
