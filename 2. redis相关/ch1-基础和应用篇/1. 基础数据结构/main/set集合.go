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

	var retCode int
	var retString string
	var retStrings []string

	keys, err := redis.Strings(conn.Do("keys", "set*"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	var cmd []interface{}
	for k := range keys {
		cmd = append(cmd, keys[k])
	}
	fmt.Println(cmd)
	_, err = conn.Do("del", cmd...)

	// sadd
	retCode, err = redis.Int(conn.Do("sadd", "set-key1", "node1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	retCode, err = redis.Int(conn.Do("sadd", "set-key1", "node2"))
	retCode, err = redis.Int(conn.Do("sadd", "set-key1", "node2"))
	retCode, err = redis.Int(conn.Do("sadd", "set-key1", "node2"))
	retCode, err = redis.Int(conn.Do("sadd", "set-key1", "node3"))

	fmt.Println("return code:", retCode)

	// scard
	retCode, err = redis.Int(conn.Do("scard", "set-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	// smembers
	retStrings, err = redis.Strings(conn.Do("smembers", "set-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return strings:", retStrings)

	// sismember
	retCode, err = redis.Int(conn.Do("sismember", "set-key1", "node1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)
	retCode, err = redis.Int(conn.Do("sismember", "set-key1", "node9"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	// spop
	retString, err = redis.String(conn.Do("spop", "set-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return string:", retString)
}
