package main

import "chatgpt_x/bootstrap"

// init Service.
func init() {
	bootstrap.Setup()
}

// main 入口函数.
func main() {
	Run(bootstrap.NewServe())
}
