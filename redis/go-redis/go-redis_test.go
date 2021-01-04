package go_redis

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	goRedis, err := New("127.0.0.1:6379", "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(goRedis.Set("myname", "derek", time.Second*1).Err())
}
