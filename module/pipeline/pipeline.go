package pipeline

import (
	"fmt"
	"time"

	"github.com/containerops/vessel/models"
)

func BootPipelineVersion(pipelineVersion models.PipelineVersion) {

	stageCount := 65535
	stageBootTag := map[string]string{}
	stageNeedBootChans := make(chan models.Stage, stageCount)
	stageTagBootChans := make(chan models.Stage, stageCount)
	stageStatusChans := make(chan models.Stage, stageCount)

	for {
		select {
		case stageStatus := <-stageStatusChans:
			fmt.Println(stageStatus)
			//if stageStatus is success then goroutine call PutNeedBootStage func
		case stageTagBoot := <-stageTagBootChans:
			if stageBootTag[stageTagBoot.Name] == "" {
				stageBootTag[stageTagBoot.Name] = stageTagBoot.Name
				stageNeedBootChans <- stageTagBoot
			}
		case stageNeedBoot := <-stageNeedBootChans:
			fmt.Println(stageNeedBoot)
			//goroutine call BootStage func
		case <-time.After(time.Duration(3) * time.Second): //设置超时时间为３ｓ，如果channel　3s钟没有响应，一直阻塞，则报告超时，进行超时处理．
			fmt.Println("timeout")
			return
		}
	}
}

func PutNeedBootStage(bootedStage models.Stage, stageTagBootChans chan models.Stage) {

}

func BootStage(needBootStage models.Stage, stageStatusChans chan models.Stage) {

}
