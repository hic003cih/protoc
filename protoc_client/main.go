package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"protoc/echo"
)

func main()  {
	ctx :=context.Background()

	conn,err :=grpc.Dial("localhost:8080",grpc.WithInsecure())
	if err != nil{
		panic(err)
	}

	defer conn.Close()

	e :=echo.NewEchoServerClient(conn)
	resp,err :=e.Echo(ctx,&echo.EchoRequest{
		Message: "Hello World!",
	})
	if err !=nil{
		panic(err)
	}

	fmt.Println("Got from Serve:",resp.Response)
}