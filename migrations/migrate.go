package main

import (
	"context"
	"flag"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/ZooLearn/zoo/config"
	"github.com/ZooLearn/zoo/ent"
	"github.com/ZooLearn/zoo/pkg/log"
)

func main() {
	flag.Parse()

	c, err := config.NewConfig("./config.yaml")
	if err != nil {
		log.Panicf("config.NewConfig error: %v", err)
	}
	db := ent.NewClient(
		ent.Log(log.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	if err := db.Schema.Create(context.Background(), schema.WithForeignKeys(false), schema.WithDropColumn(true),
		schema.WithDropIndex(true)); err != nil {
		panic(err)
	}
}
