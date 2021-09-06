package server

import (
	"fmt"
	"github.com/1-bi/log-api"
	"time"
)

// NodeWorker 节点服务器，用来控制服务器应用管理
type NodeWorker struct {
	controlCh chan int

	// 该节点服务的角色
	role string

	logger logapi.Logger
}

func NewNodeWorker(role string) *NodeWorker {
	wm := new(NodeWorker)
	wm.role = role
	wm.logger = logapi.GetLogger("pisauridae.master.NodeWorker")
	return wm
}

func (myself *NodeWorker) Stop() {

	time.Sleep(6 * time.Second)

	myself.controlCh <- CMD_STOP
}

func (myself *NodeWorker) readAndCheckConfig() error {

	return nil
}

func (myself *NodeWorker) Start() error {
	var err error

	// --- check config
	err = myself.readAndCheckConfig()
	if err != nil {
		return err
	}

	myself.controlCh = make(chan int)

	// --- start  the run
	go func() {
		myself.controlCh <- CMD_RUN
	}()

	// --- close channel --
	defer close(myself.controlCh)

	var flagBreak = false
	for !flagBreak {
		select {
		case recSign := <-myself.controlCh:

			if CMD_RUN == recSign {

				go func() {
					// boot streap main logic
					mainEntry()
					err = nil
				}()

			} else if CMD_PAUSE == recSign {

				go func() {

				}()

			} else if CMD_STOP == recSign {

				fmt.Println("stop server ")

				flagBreak = true
			}

		}
	}

	return nil
}
