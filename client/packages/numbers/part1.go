package numbers

import (
	"fmt"
	"log"
)

type Number int

type Result struct{
	Quotient int
	Remain int
}

var (
	Logs *log.Logger
)

func (n *Number)Multiply(val int, res *int)error{
	LogOut(fmt.Sprintf("got value: %v", val))
	*res = val * 2
	LogOut(fmt.Sprintf("result value: %v", *res))
	return nil
}

func (n * Number) Division(val int , res *Result )error{
	LogOut(fmt.Sprintf("got value: %v", val))
	res.Quotient =  val / 2
	res.Remain = val % 2
	LogOut(fmt.Sprintf("resultStruct.Quotient, resultStruct.Remain : %v, %v", res.Quotient, res.Remain))
	return nil
}

func LogOut(mes string){
	if Logs != nil{
		Logs.Println(mes)
	}
}

