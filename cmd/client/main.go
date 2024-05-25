package main

import (
	"context"
	"log"
	"net/http"
	"os"

	greetv1 "connect-go-example/gen/greet/v1"
	"connect-go-example/gen/greet/v1/greetv1connect"

	"connectrpc.com/connect"
)

func main() {
	var url string
	if len(os.Args) < 2 {
		url = "http://localhost:8080"
	} else {
		url = os.Args[1]
	}

	opts := []connect.ClientOption{}
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "--grpc":
			opts = append(opts, connect.WithGRPC())
		case "--grpc-web":
			opts = append(opts, connect.WithGRPCWeb())
		}
	}

	client := greetv1connect.NewGreetServiceClient(http.DefaultClient, url, opts...)
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
