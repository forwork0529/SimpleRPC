package main

import (
	"log"
	"net"
	"net/rpc"
	"os"
	"server/files"
	"server/packages/numbers"
)

var(
	network = "tcp4"
	addr = "127.0.0.1:12345"
)


func main(){

	output, err := os.Create(files.Pwd + "output.txt")
	if err != nil{
		log.Fatal(err)
	}
	numbers.Logs = log.New(output, "INFO :", log.LstdFlags )
	n := new(numbers.Number)
	err = rpc.Register(n)
	if err != nil{
		log.Fatal(err)
	}
	l, err := net.Listen(network, addr)
	if err != nil{
		log.Fatal(err)
	}
	rpc.Accept(l)
}

