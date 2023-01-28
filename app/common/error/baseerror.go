package error

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

const defaultCode = 1001

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
func HandleGrpcErrorToHttp(err error) (response *CodeError) {
	//将grpc的code转换成http的状态码

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				response = &CodeError{
					Code: http.StatusNotFound,
					Msg:  e.Message(),
				}
			case codes.Internal:
				response = &CodeError{
					Code: http.StatusInternalServerError,
					Msg:  "内部错误",
				}
			case codes.InvalidArgument:

				response = &CodeError{
					Code: http.StatusBadRequest,
					Msg:  "参数错误",
				}
			case codes.Unavailable:

				response = &CodeError{
					Code: http.StatusInternalServerError,
					Msg:  "用户服务不可用",
				}
			default:
				response = &CodeError{
					Code: http.StatusInternalServerError,
					Msg:  e.Code().String(),
				}
			}
			return response
		}
	}
	return &CodeError{
		Code: defaultCode,
		Msg:  err.Error(),
	}
}
