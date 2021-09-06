package main

import (
	"flag"
	"github.com/1-bi/log-api"
	"github.com/1-bi/log-zap"
	"github.com/1-bi/log-zap/appender"
	zaplayout "github.com/1-bi/log-zap/layout"
	"github.com/1-bi/trader-collector/server"
	"log"
)

/**
 * init progfile config
 */
func init() {

	// --- find mode
	//var node = flag.String("node", "client", "Use --node <nodemode> , supported node  : [master] [workder] [client]")
	//var profile = flag.String("profile", "dev", "Use --profile <profile> , supported mode : [dev] [prod] [uat]")
	flag.Parse()

	//var rProfile = splider.GetRuntimeProfile()
	// assign value to runtime inst
	//rProfile["node"] = *node
	//rProfile["profile"] = *profile

}

// prepareLogSetting 预设置日志访问配置格式
func prepareLogSetting() {

	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	//jsonLayout.SetTimeFormat("2006-01-02 15:04:05")
	jsonLayout.SetTimeFormat("2006-01-02 15:04:05 +0800 CST")
	//fmt.Println( time.Now().Location() )

	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var mainOpt = logzap.NewLoggerOption()
	mainOpt.SetLevel("warn")
	mainOpt.AddAppender(consoleAppender)

	var agentOpt = logzap.NewLoggerOption()
	agentOpt.SetLoggerPattern("trader-collector")
	agentOpt.SetLevel("debug")
	agentOpt.AddAppender(consoleAppender)

	var implReg = new(logzap.ZapFactoryRegister)

	_, err := logapi.RegisterLoggerFactory(implReg, mainOpt, agentOpt)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	prepareLogSetting()

	//server.Main()

	nodeWorker := server.NewNodeWorker("master")

	println(nodeWorker)

}
