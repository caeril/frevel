package app

import (
	"github.com/caeril/frevel/revel"
)

func init() {
	revel.OnAppStart(func() {
		revel.AppLog.Info("Go to /@tests to run the tests.")
	})
}
