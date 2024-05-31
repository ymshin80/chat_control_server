package main

import (
	"chat_golang_control/cmd/app"
	"chat_golang_control/config"
	"flag"
)

var pathFlg = flag.String("config", "./config.toml", "config set")

func main(){
	flag.Parse()
	c := config.NewConfig(*pathFlg)
	//TODO: app 객체를 사용해서 구동.
	a := app.NewApp(c)
	a.Start()
	
}