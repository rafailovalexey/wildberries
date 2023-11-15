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
	return func(_context context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		_metadata, isExist := metadata.FromIncomingContext(_context)

		if !isExist {
			return nil, status.Errorf(codes.Internal, "Failed to read metadata")
		}

		tracecode := _metadata["tracecode"][0]

		log.Printf("Incoming GRPC request: %s (%s)", info.FullMethod, tracecode)

		response, _error := handler(_context, request)

		if _error != nil {
			log.Printf("Error in GRPC request: %s (%s) \n %v", info.FullMethod, tracecode, _error)
		}

		if _error == nil {
			log.Printf("Outgoing GRPC response: %s (%s)", info.FullMethod, tracecode)
		}

		return response, _error
	}
}
