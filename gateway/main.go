package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"
	"github.com/tommy-sho/telepresence-with-envoy/proto"
	"google.golang.org/grpc"
)

const (
	port = "50002"
)

func main() {
	ctx := context.Background()

	bConn, err := grpc.DialContext(ctx, os.Getenv("BACKEND_PORT"), grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("failed to connect with backend server error : %v ", err))
	}

	bClient := proto.NewBackendServerClient(bConn)

	e := echo.New()
	e.GET("/greeting", Greeting(bClient))

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	go func() {
		<-stopChan
		if err := e.Close(); err != nil {
			log.Print("Failed to stop server")
		}
	}()

	errors := make(chan error)
	go func() {
		errors <- e.Start(fmt.Sprintf(":%s", port))
	}()

	if err := <-errors; err != nil {
		log.Fatal("Failed to server gRPC server", err)
	}
}

func Greeting(clinet proto.BackendServerClient) echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "Hello World")
	}
}
