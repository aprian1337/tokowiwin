package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"time"
	"tokowiwin/constants"
	"tokowiwin/repositories/db"
)

//var viperClient string

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	timeout := viper.GetInt(constants.CONTEXT_TIMEOUT)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	z := db.NewDatabaseRepository(ctx)
	email, err := z.GetUserByEmail(ctx, "z")
	if err != nil {
		fmt.Println("ERR", err)
		return
	}
	em, _ := json.Marshal(email)
	fmt.Println("EMAIL", string(em))
}
