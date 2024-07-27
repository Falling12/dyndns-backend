package db

import (
	"context"
	"fmt"
)

var DB *PrismaClient
var Ctx = context.Background()

func Connect() {
	DB = NewClient()

	err := DB.Connect()

	if err != nil {
		println(err)
	}

	fmt.Println("[database] Database sucessfully connected")
}

func Disconnect() {
	err := DB.Disconnect()

	if err != nil {
		panic(err)
	}
}
