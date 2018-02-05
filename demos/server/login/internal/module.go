package internal

import (
	"github.com/zsai001/leaf_cluster/module"
	"github.com/zsai001/leaf_cluster/demos/server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
