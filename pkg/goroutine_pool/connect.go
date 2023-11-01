package goroutine_pool

import (
	"github.com/panjf2000/ants/v2"
	"goImPro-service/src/config"
)

var AntsPool *ants.Pool

func ConnectPool() *ants.Pool {
	//设置数量
	AntsPool, _ = ants.NewPool(config.Conf.Server.CoroutinePoll)
	return AntsPool
}
