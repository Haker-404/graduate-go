package utils

import (
	"context"
	"google.golang.org/grpc"
)

func Interceptor(interceptor, next grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		return interceptor(ctx, req, info,
			func(nextCtx context.Context, nextReq interface{}) (interface{}, error) {
				return next(nextCtx, nextReq, info, handler)
			})
	}
}
