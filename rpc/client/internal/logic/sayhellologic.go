package logic

import (
	"context"

	"client/hello"
	"client/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *hello.HelloRequest) (*hello.HelloReply, error) {
	// todo: add your logic here and delete this line

	return &hello.HelloReply{}, nil
}
