package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

func AuthenticationTokenInterceptor() grpc.UnaryServerInterceptor {
	return func(_context context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		_metadata, isExist := metadata.FromIncomingContext(_context)

		if !isExist {
			return nil, status.Errorf(codes.Unauthenticated, "Authentication token not found")
		}

		header := os.Getenv("AUTHENTICATION_TOKEN_HEADER")

		if header == "" {
			log.Fatalf("Not found authentication token header in environment")
		}

		list := _metadata[header]

		if len(list) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "Authentication token not found")
		}

		key := list[0]

		token := os.Getenv("AUTHENTICATION_TOKEN")

		if token == "" {
			log.Fatalf("Not found authentication token in environment")
		}

		if token != key {
			log.Printf("Invalid authentication token: %s", key)

			return nil, status.Errorf(codes.PermissionDenied, "Invalid authentication token")
		}

		return handler(_context, request)
	}
}
