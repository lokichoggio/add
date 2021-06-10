package logic

import (
	"context"

	"github.com/lokichoggio/add/add"
	"github.com/lokichoggio/add/internal/common/errorx"
	"github.com/lokichoggio/add/internal/svc"
	"github.com/lokichoggio/add/model"

	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/codes"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	_, err := l.svcCtx.Model.Insert(model.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, errorx.CodeMsgErrorWithStack(codes.Internal, "mysql error", err)
	}

	return &add.AddResp{
		Ok: true,
	}, nil
}
