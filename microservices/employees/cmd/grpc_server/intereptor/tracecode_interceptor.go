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
	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, isExist := metadata.FromIncomingContext(ctx)

		if !isExist {
			log.Printf("TracecodeInterceptor: Metadata not found in the request context")

			return nil, status.Errorf(codes.Internal, "Failed to read metadata")
		}

		if len(md["tracecode"]) != 0 {
			return handler(ctx, request)
		}

		tracecode, err := GenerateTracecode()

		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to generate tracecode")
		}

		md = metadata.Join(md, metadata.New(map[string]string{"tracecode": tracecode}))
		ctx = metadata.NewIncomingContext(ctx, md)

		return handler(ctx, request)
	}
}

func GenerateTracecode() (string, error) {
	tracecode := make([]byte, 16)

	if _, err := rand.Read(tracecode); err != nil {
		return "", err
	}

	return hex.EncodeToString(tracecode), nil
}
