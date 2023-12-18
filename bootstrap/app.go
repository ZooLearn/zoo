package bootstrap

import (
	"github.com/ZooLearn/zoo/config"
	"github.com/ZooLearn/zoo/ent"
	"github.com/ZooLearn/zoo/internal/log"
)

type Application struct {
	Env config.EnvConf
	Ent *ent.Client
}

func App() Application {
	app := &Application{}
	cfgs, err := config.NewConfig("../config.yaml")
	if err != nil {
		log.Panicf("config.NewConfig error: %v", err)
	}
	db := ent.NewClient(
		ent.Log(log.Info), // logger
		ent.Driver(cfgs.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	app.Env = cfgs.EnvConf
	app.Ent = db
	return *app
}
