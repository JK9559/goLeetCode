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

	keys, err := redis.Strings(conn.Do("keys", "hash*"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	var cmd []interface{}
	for k := range keys {
		cmd = append(cmd, keys[k])
	}
	fmt.Println(cmd)
	_, err = conn.Do("del", cmd...)

	// hset
	retCode, err = redis.Int(conn.Do("hset", "hash-map1", "key1", "val1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	retCode, err = redis.Int(conn.Do("hset", "hash-map1", "key2", "val2"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	// hgetall
	retStrings, err = redis.Strings(conn.Do("hgetall", "hash-map1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return strings:", retStrings)

	// hget
	retString, err = redis.String(conn.Do("hget", "hash-map1", "key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return string:", retString)

	// hlen
	retCode, err = redis.Int(conn.Do("hlen", "hash-map1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	// hmset
	retString, err = redis.String(conn.Do("hmset", "hash-map1", "key3", "val3", "key4", "val4"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return string:", retString)

	retCode, err = redis.Int(conn.Do("hlen", "hash-map1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	// hincrby

}
