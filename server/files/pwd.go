package files

import (
	"log"
	"os"
)

var (
	Pwd string
	err error
)


func init(){
	pwd ,err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	_ = pwd
}