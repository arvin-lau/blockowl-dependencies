package chainExplorer

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

const retryNum = 3

func chainExplorerRetrySleep(retryNum int) {
	log.Error().Msg(fmt.Sprintf("chainExplorer retry... retryNum: %v", retryNum))
	time.Sleep(time.Second * 3)
}
