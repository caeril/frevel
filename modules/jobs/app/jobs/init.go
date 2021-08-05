package jobs

import (
	"github.com/caeril/frevel/revel"
)

var jobLog = revel.AppLog

func init() {
	revel.RegisterModuleInit(func(m *revel.Module) {
		jobLog = m.Log
	})
}
