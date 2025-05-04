package main

import "api-golang/configs"


func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}