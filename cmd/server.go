package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/resin-io/adapter-base/adapter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the adapter-base server",
	Run:   serverCmd,
}

func init() {
	ServerCmd.Flags().IntVarP(&apiPort, "apiPort", "a", 8080, "REST API port")
	ServerCmd.Flags().IntVarP(&rpcPort, "rpcPort", "r", 8081, "RPC API port")
	ServerCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Max concurrent workers")
}

func serverCmd(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	wg.Add(2)

	rpcServer := grpc.NewServer()

	server := adapter.CreateServer(concurrency, verbose)
	adapter.RegisterScanServer(rpcServer, server)
	adapter.RegisterUpdateServer(rpcServer, server)

	rpcAddr := fmt.Sprintf("localhost:%v", rpcPort)
	conn, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		log.WithFields(log.Fields{
			"Address": rpcAddr,
		}).Fatal("Failed to open connection")
	}

	go func() {
		defer wg.Done()
		if err := rpcServer.Serve(conn); err != nil {
			log.WithFields(log.Fields{
				"Address": rpcAddr,
			}).Fatal("Failed to start RPC server")
		}
	}()

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := adapter.RegisterUpdateHandlerFromEndpoint(ctx, mux, rpcAddr, opts); err != nil {
		log.Fatal("Failed to register update handler")
	}
	if err := adapter.RegisterScanHandlerFromEndpoint(ctx, mux, rpcAddr, opts); err != nil {
		log.Fatal("Failed to register scan handler")
	}

	apiAddr := fmt.Sprintf("localhost:%v", apiPort)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(apiAddr, mux); err != nil {
			log.WithFields(log.Fields{
				"Address": apiAddr,
			}).Fatal("Failed to start REST server")
		}
	}()

	log.WithFields(log.Fields{
		"API": apiAddr,
		"RPC": rpcAddr,
	}).Info("Started adapter-base module servers")

	wg.Wait()
}
