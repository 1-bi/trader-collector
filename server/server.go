package server

import (
	di "github.com/1-bi/fire-di"
	"github.com/1-bi/log-api"
)

// Main  Application program entry
func mainEntry() {

	logapi.GetLogger("collctor.server").Info("Start Trader-Collector-Server ...", nil)

	diConf := new(di.Configuration)

	/**
	 * custom  di config
	 */
	di.Config(diConf)

}

//  startMaster start master server with engine
func startMaster() {

}
