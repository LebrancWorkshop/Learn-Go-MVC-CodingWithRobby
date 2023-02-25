package main

import (
	"fmt"
	"lebrancconvas/gomvc-robby/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
}

func main() {
	fmt.Println("Hello Go MVC Naja.")	
}