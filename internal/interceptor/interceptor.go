package interceptor

import (
	"context"

	"github.com/lokichoggio/add/internal/common/errorx"

	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)

	code := status.Convert(err).Code()
	msg := status.Convert(err).Message()

	if code != errorx.SuccessCode {
		logx.WithContext(ctx).Errorf("business_code: %d, err_detail: %+v", code, msg)
		err = errorx.CodeError(code)
	}

	return resp, err
}
