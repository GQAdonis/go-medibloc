package rpc

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	rpcpb "github.com/medibloc/go-medibloc/rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func httpServerRun() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err1 := rpcpb.RegisterApiServiceHandlerFromEndpoint(ctx, mux, "localhost:10000", opts)
	err2 := rpcpb.RegisterAdminServiceHandlerFromEndpoint(ctx, mux, "localhost:10000", opts)
	if err1 != nil {
		log.Printf("Somethins is wrong in httpServerRun : %v", err1)
	}
	if err2 != nil {
		log.Printf("Somethins is wrong in httpServerRun : %v", err2)
	}
	if err := http.ListenAndServe("localhost:10002", mux); err != nil {
		log.Printf("Somethins is wrong in httpServerRun : %v", err)
		return err
	}
	return nil
}