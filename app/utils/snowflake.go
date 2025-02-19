package utils

import (
	"go.uber.org/zap"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func InitSnowflake(startTime string, machineID int64) (err error) {
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		zap.L().Error("time.Parse", zap.Error(err))
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
