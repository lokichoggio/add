package errorx

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	SuccessCode codes.Code = 0
)

var errMap = map[codes.Code]string{
	SuccessCode: "成功",
}

func CodeError(code codes.Code) error {
	return status.Error(code, errMap[code])
}

func CodeErrorWithStack(code codes.Code, err error) error {
	err = errors.WithStack(err)
	return status.Errorf(code, "%v: %+v", errMap[code], err)
}

func CodeMsgErrorWithStack(code codes.Code, msg string, err error) error {
	err = errors.WithStack(err)
	return status.Errorf(code, "%v: %+v", msg, err)
}
