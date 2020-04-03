/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, b)
	}
	return string(ba)
}

func main() {
	conn, err := redis.Dial("tcp", "39.97.162.199:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	var retCode int
	//var retString string
	var retStrings []string
	//var retInts []int

	keys, err := redis.Strings(conn.Do("keys", "zset*"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	var cmd []interface{}
	for k := range keys {
		cmd = append(cmd, keys[k])
	}
	fmt.Println(cmd)
	_, err = conn.Do("del", cmd...)

	// zadd
	retCode, err = redis.Int(conn.Do("zadd", "zset-key1", "3501.6", "node1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)
	retCode, err = redis.Int(conn.Do("zadd", "zset-key1", "3502", "node2"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)
	retCode, err = redis.Int(conn.Do("zadd", "zset-key1", "1000", "node3"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	// zrange
	retStrings, err = redis.Strings(conn.Do("zrange", "zset-key1", "0", "-1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return strings:", retStrings)

	// zrevrange
	retStrings, err = redis.Strings(conn.Do("zrevrange", "zset-key1", "0", "-1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return strings:", retStrings)

	// zcard
	retCode, err = redis.Int(conn.Do("zcard", "zset-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return int:", retCode)

	// zscore
	k, err := conn.Do("zscore", "zset-key1", "node1")
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return strings:", string(k.([]uint8)))

	// zrank
	retCode, err = redis.Int(conn.Do("zrank", "zset-key1", "node3"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	// zrangebyscore
	retStrings, err = redis.Strings(conn.Do("zrangebyscore", "zset-key1", "-inf", "9999", "withscores"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	for i := 0; 2*i < len(retStrings); i++ {
		fmt.Println("keys, vals:", retStrings[2*i], retStrings[2*i+1])
	}
	fmt.Println("return strings", retStrings)

	// zrem
	retCode, err = redis.Int(conn.Do("zrem", "zset-key1", "node1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	retCode, err = redis.Int(conn.Do("zcard", "zset-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)
}
