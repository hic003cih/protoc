package main

import (
	"context"
	"fmt"
	"net"
	"protoc/echo"
	"google.golang.org/grpc"
)

//使用空的struct,不用站記憶體空間,效率更高
//佔空間的用法如下
//type EchoServer int
//srv := EchoServer(0)
type EchoServer struct{}

func (e *EchoServer) Echo(ctx context.Context,req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Response:"My Echo:" + req.Message,
	},nil
}
func main(){
	lst,err :=net.Listen("tcp",":8080")
	if err !=nil{
		panic(err)
	}

	s :=grpc.NewServer()

	srv := &EchoServer{}

	echo.RegisterEchoServerServer(s,srv)

	fmt.Println("Now serving at port 8080")
	err = s.Serve(lst)
	if err !=nil{
		panic(err)
	}
}