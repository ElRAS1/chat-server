package config

import (
	"context"
	"github.com/ELRAS1/chat-server/pkg/chatServer"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitHTTP(ctx context.Context, grpcPort, httpPort string) *http.Server {
	mux := runtime.NewServeMux()
	if err := chatServer.RegisterChatServerHandlerFromEndpoint(ctx, mux, grpcPort, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}); err != nil {
		log.Fatalln(err)
	}

	return &http.Server{
		Handler: corsMiddleware(mux),
		Addr:    httpPort,
	}
}

func InitSwagger() *http.ServeMux {
	swaggerHTTP := http.NewServeMux()
	swaggerHTTP.HandleFunc("/api.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/swagger/api.swagger.json")
	})

	swaggerHTTP.Handle("/", http.FileServer(http.Dir("./swagger-ui/")))

	return swaggerHTTP
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, UPDATE, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
