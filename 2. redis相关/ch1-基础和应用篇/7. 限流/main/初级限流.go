/*
 * Copyright © 2020 wkang. All rights reserved.
 */

package main

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

const (
	nodeBits  uint8 = 10
	stepBits  uint8 = 12
	nodeMax   int64 = -1 ^ (-1 << nodeBits)
	stepMax   int64 = -1 ^ (-1 << stepBits)
	timeShift       = nodeBits + stepBits
	nodeShift       = stepBits
)

var Epoch int64 = 1288834974657

type Node struct {
	mu        sync.Mutex
	timestamp int64
	node      int64
	step      int64
}

func NewNode(node int64) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}
	return &Node{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}

func (node *Node) Generate() int64 {
	node.mu.Lock()
	defer node.mu.Unlock()

	now := time.Now().UnixNano()

	if node.timestamp == now {
		node.step++

		if node.step > stepMax {
			for now <= node.timestamp {
				now = time.Now().UnixNano()
			}
		}
	} else {
		node.step = 0
	}

	node.timestamp = now

	res := (now-Epoch)<<timeShift | node.node<<nodeShift | node.step
	return res
}

type redisStr struct {
	conn redis.Conn
	node Node
}

func newRedis() *redisStr {
	conn, err := redis.Dial("tcp", "39.97.162.199:6379")
	if nil != err {
		fmt.Println("Initial Redis ERROR: ", err)
		return nil
	}
	node, err := NewNode(1)
	if nil != err {
		fmt.Println("Initial Snowflake ERROR: ", err)
	}
	return &redisStr{conn: conn, node: *node}
}

func (red *redisStr) allow(usrId, action string, period, maxCount int64) bool {
	key := fmt.Sprintf("key:%s:%s", usrId, action)
	nowTime := time.Now().UnixNano() / 1e6
	red.conn.Send("multi")
	red.conn.Send("zadd", key, nowTime, fmt.Sprintf("%d", red.node.Generate()))
	red.conn.Send("zremrangebyscore", key, 0, nowTime-1000*period)
	red.conn.Send("zcard", key)
	red.conn.Send("expire", key, period+1)
	red.conn.Flush()
	reply, err := redis.Values(red.conn.Do("exec"))
	if nil != err {
		fmt.Println("Redis ERROR:", err)
	}
	fmt.Println(reply)
	return reply[2].(int64) <= maxCount
}

func main() {

	red := newRedis()
	defer red.conn.Close()

	for i := 0; i < 20; i++ {
		// 同一个用户reply的动作 每60秒可以访问5次
		fmt.Println(red.allow("jk9559", "reply", 60, 5))
	}

}
