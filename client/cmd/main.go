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
	var resMultiple int
	err = cl.Call("Number.Multiply", 10, &resMultiple)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("result from RPC(Call Multiply) : %v\n", resMultiple)


	var resDivision numbers.Result
	err = cl.Call("Number.Division", 5, &resDivision)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("result from RPC(Call Division) : %v, %v\n", resDivision.Quotient, resDivision.Remain)


	call := cl.Go("Number.Division", 5, &resDivision, nil )
	<- call.Done
	if err = call.Error; err != nil{
		fmt.Println(err)

	}else{
		fmt.Printf("result of RPC(Go Division) : %v, %v\n", resDivision.Quotient, resDivision.Remain)
	}
}
