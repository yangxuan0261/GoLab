package test_redis

import (
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
)

// 参考: https://blog.csdn.net/wangshubo1989/article/details/75050024
// 详细参考: https://github.com/go-redis/redis/blob/master/example_test.go

func Test_001(t *testing.T) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	// TODO: 暂时不知道怎么通过redis做管道
	res, err := c.Do("PUBLISH", "hello", "test send2")
	if err == nil {
		fmt.Println("res:", res)
	}

	// SUBSCRIBE
}

func Test_subscribe(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_ = client
}