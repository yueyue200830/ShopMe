package main

import (
	"backend-go/api"
	"backend-go/dao"
)

func main() {
	api.InitServer()
	defer dao.Close()
}
