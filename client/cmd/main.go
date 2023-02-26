package main

import (
	"client/packages/numbers"
	"fmt"
	"log"
	"net/rpc"
)

var(
	network = "tcp4"
	addr = "127.0.0.1:12345"
)

func main(){
	cl, err := rpc.Dial(network, addr)
	if err != nil{
		log.Fatal(err.Error())
	}
	var res numbers.Result
	err = cl.Call("Number.Division", 5, &res)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("result of RPC : %v, %v\n", res.Quotient, res.Remain)


	call := cl.Go("Number.Division", 5, &res, nil )
	<- call.Done
	if err = call.Error; err != nil{
		fmt.Println(err)

	}else{
		fmt.Printf("result of RPC : %v, %v\n", res.Quotient, res.Remain)
	}
}
