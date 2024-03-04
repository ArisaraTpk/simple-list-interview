package main

import (
	"simple-list-interview/https"
	"simple-list-interview/infra"
)

func init() {
	infra.InitConfig()
	infra.InitDb()
}
func main() {
	https.InitRoutes()
}
