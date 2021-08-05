package controllers

import (
	"github.com/newrelic/go-agent"
	"github.com/caeril/frevel/modules/server-engine/newrelic"
	"github.com/caeril/frevel/revel"
)

type RelicController struct {
	*revel.Controller
}

func (r *RelicController) GetRelicApplication() newrelic.Application {
	if app, ok := revel.CurrentEngine.(*revelnewrelic.ServerNewRelic); ok {
		return app.NewRelicApp
	}
	return nil
}
