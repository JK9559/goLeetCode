/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "39.97.162.199:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	_ = conn.Send("multi")

	_ = conn.Send("set", "pipeline-key1", "val1")
	_ = conn.Send("set", "pipeline-key2", "val2")
	_ = conn.Send("get", "pipeline-key3")
	_ = conn.Send("del", "pipeline-key1")

	_ = conn.Flush()

	reply, err := redis.Values(conn.Do("exec"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	/********************************************/
	for _, r := range reply {
		switch r.(type) {
		case []uint8:
			fmt.Println(string(r.([]uint8)))
		case nil:
			fmt.Println(nil)
		default:
			fmt.Println(r)
		}
	}
	/********************************************/
}
