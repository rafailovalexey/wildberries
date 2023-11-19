package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, isExist := metadata.FromIncomingContext(ctx)

		if !isExist {
			return nil, status.Errorf(codes.Internal, "Failed to read metadata")
		}

		tracecode := md["tracecode"][0]

		log.Printf("Incoming GRPC request: %s (%s)", info.FullMethod, tracecode)

		response, err := handler(ctx, request)

		if err != nil {
			log.Printf("Error in GRPC request: %s (%s) \n %v", info.FullMethod, tracecode, err)
		}

		if err == nil {
			log.Printf("Outgoing GRPC response: %s (%s)", info.FullMethod, tracecode)
		}

		return response, err
	}
}
