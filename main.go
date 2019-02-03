package main

import (
	"github.com/KarloB/kfb/internal/firebase"
	"github.com/KarloB/kfb/internal/model/conf"
	"github.com/KarloB/kfb/internal/worker"
)

var (
	confPath = "conf.json"
)

func main() {
	conf, err := conf.Get(confPath)
	if err != nil {
		panic(err)
	}

	client, err := firebase.New(conf)
	if err != nil {
		panic(err)
	}

	w := worker.New(client, nil)
	err = w.Watch()
	if err != nil {
		panic(err)
	}
}
