package main

import (
	"github.com/zsai001/leaf_cluster"
	lconf "github.com/zsai001/leaf_cluster/conf"
	"github.com/zsai001/leaf_cluster/demos/server/conf"
	"github.com/zsai001/leaf_cluster/demos/server/game"
	"github.com/zsai001/leaf_cluster/demos/server/gate"
	"github.com/zsai001/leaf_cluster/demos/server/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
