package base

import (
	"github.com/zsai001/leaf_cluster/chanrpc"
	"github.com/zsai001/leaf_cluster/module"
	"github.com/zsai001/leaf_cluster/demos/server/conf"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
