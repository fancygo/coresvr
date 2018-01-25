package msgproc

import (
	"frame/logger"
	"time"
)

func ProcMsgData(data []byte) {
	start := time.Now()
	dataLen := len(data)
	logger.Info("proc msg data = %d\n", dataLen)
	secs := time.Since(start).Seconds()
	logger.Info("Msg proc last secs = %0.5f", secs)
}
