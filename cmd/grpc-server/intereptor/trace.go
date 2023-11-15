package interceptor

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func TracecodeInterceptor() grpc.UnaryServerInterceptor {
	return func(_context context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		_metadata, isExist := metadata.FromIncomingContext(_context)

		if !isExist {
			log.Printf("TracecodeInterceptor: Metadata not found in the request context")

			return nil, status.Errorf(codes.Internal, "Failed to read metadata")
		}

		if len(_metadata["tracecode"]) != 0 {
			return handler(_context, request)
		}

		tracecode, _error := GenerateTracecode()

		if _error != nil {
			return nil, status.Errorf(codes.Internal, "Failed to generate tracecode")
		}

		_metadata = metadata.Join(_metadata, metadata.New(map[string]string{"tracecode": tracecode}))
		_context = metadata.NewIncomingContext(_context, _metadata)

		return handler(_context, request)
	}
}

func GenerateTracecode() (string, error) {
	tracecode := make([]byte, 16)

	if _, _error := rand.Read(tracecode); _error != nil {
		return "", _error
	}

	return hex.EncodeToString(tracecode), nil
}
