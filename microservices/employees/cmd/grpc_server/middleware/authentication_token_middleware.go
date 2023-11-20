package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

func AuthenticationTokenMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, isExist := metadata.FromIncomingContext(ctx)

		if !isExist {
			return nil, status.Errorf(codes.Unauthenticated, "authentication token not found")
		}

		header := os.Getenv("AUTHENTICATION_TOKEN_HEADER")

		if header == "" {
			log.Fatalf("not found authentication token header in environment")
		}

		list := md[header]

		if len(list) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authentication token not found")
		}

		key := list[0]

		token := os.Getenv("AUTHENTICATION_TOKEN")

		if token == "" {
			log.Fatalf("not found authentication token in environment")
		}

		if token != key {
			log.Printf("invalid authentication token: %s", key)

			return nil, status.Errorf(codes.PermissionDenied, "invalid authentication token")
		}

		return handler(ctx, request)
	}
}
