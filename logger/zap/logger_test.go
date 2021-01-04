package zap

import (
	"go.uber.org/zap"
	"testing"
)

func TestInfo(t *testing.T) {

	log.Info("这是一个logger测试", zap.String("level", LevelInfo))
}
