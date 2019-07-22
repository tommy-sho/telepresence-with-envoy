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

type Request struct {
	Name string `json:"name" form:"name" query:"name"`
}

type Responce struct {
	Message  string `json:"message"`
	DateTime int64  `json:"datetime"`
}

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

func Greeting(client proto.BackendServerClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(Request)
		if err := c.Bind(r); err != nil {
			return err
		}

		ctx := context.Background()
		req := &proto.MessageRequest{
			Name: r.Name,
		}
		m, err := client.Message(ctx, req)
		if err != nil {
			return err
		}

		res := Responce{
			Message:  m.Message,
			DateTime: m.Datetime,
		}

		return c.JSON(http.StatusOK, res)
	}
}
