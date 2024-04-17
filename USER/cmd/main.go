package main

import (
	"context"
	"fmt"
	"user/api"
	"user/config"
	"user/pkg/logger"
	"user/service"
	"user/storage/postgres"
	"user/storage/redis"

)

func main() {
	cfg := config.Load()
	
	newRedis := redis.New(cfg)

	log := logger.New(cfg.ServiceName)
	store, err := postgres.New(context.Background(), cfg,newRedis)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	services := service.New(store,log,newRedis)
	c := api.New(services,log)

	fmt.Println("programm is running on localhost:8080...")
	c.Run(":8080")
}

