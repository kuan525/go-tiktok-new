package redis

import (
	"github.com/go-redis/redis/v7"
	"strconv"
)

func add(c *redis.Client, k string, v int64) {
	tx := c.TxPipeline()
	// k这个set中添加一个kv对
	tx.SAdd(k, v)
	tx.Expire(k, expireTime)
	tx.Exec()
}

func del(c *redis.Client, k string, v int64) {
	tx := c.TxPipeline()
	// k这个set中删掉v
	tx.SRem(k, v)
	tx.Expire(k, expireTime)
	tx.Exec()
}

func check(c *redis.Client, k string) bool {
	// 判断key在不在，任何类型对象都可以判断
	if e, _ := c.Exists(k).Result(); e > 0 {
		return true
	}
	return false
}

func exist(c *redis.Client, k string, v int64) bool {
	// 判断集合k中是否有v
	if e, _ := c.SIsMember(k, v).Result(); e {
		c.Expire(k, expireTime)
		return true
	}
	return false
}

func count(c *redis.Client, k string) (sum int64, err error) {
	// 获取集合k对成员数量
	if sum, err = c.SCard(k).Result(); err == nil {
		c.Expire(k, expireTime)
		return sum, err
	}
	return sum, err
}

func get(c *redis.Client, k string) (vt []int64) {
	v, _ := c.SMembers(k).Result()
	c.Expire(k, expireTime)
	for _, vs := range v {
		// 64指，转换后最大适应的值，如果爆炸了会报错
		vI64, _ := strconv.ParseInt(vs, 10, 64)
		vt = append(vt, vI64)
	}
	return vt
}
