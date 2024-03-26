package config_test

import (
	"fmt"
	"panduputra/miniproject3/config"
	"testing"

	"github.com/joho/godotenv"
)

func Init(){
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Print(err)
	}
}

func TestConnection(t *testing.T){
	Init()
	config.OpenDB()
}