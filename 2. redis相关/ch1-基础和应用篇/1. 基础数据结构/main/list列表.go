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

	keys, err := redis.Strings(conn.Do("keys", "list*"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	var cmd []interface{}
	for k := range keys {
		cmd = append(cmd, keys[k])
	}
	fmt.Println(cmd)
	_, err = conn.Do("del", cmd...)

	// 队列
	retCode, err = redis.Int(conn.Do("rpush", "list-key1", "node1", "node2", "node3"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	for i := 0; i < 2; i++ {
		retString, err = redis.String(conn.Do("lpop", "list-key1"))
		if retString == "" {
			break
		}
		fmt.Println(retString)
	}

	// 栈
	retCode, err = redis.Int(conn.Do("rpush", "list-key2", "node1", "node2", "node3"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	for i := 0; i < 2; i++ {
		retString, err = redis.String(conn.Do("rpop", "list-key2"))
		if err != nil {
			fmt.Println("Err: ", err)
		}
		fmt.Println(retString)
	}

	retCode, err = redis.Int(conn.Do("llen", "list-key2"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("length of list is: ", retCode)

	// 禁用的慢操作
	// index 可以为负数 表示倒数第几个元素
	retCode, err = redis.Int(conn.Do("rpush", "list-key3", "node1", "node2", "node3"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return code:", retCode)

	retString, err = redis.String(conn.Do("lindex", "list-key3", "-1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return string:", retString)

	// range 获取所有元素
	retStrings, err = redis.Strings(conn.Do("lrange", "list-key3", "0", "1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return strings:", retStrings)

	// trim 保留区间内的元素
	retString, err = redis.String(conn.Do("ltrim", "list-key3", "0", "1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	retStrings, err = redis.Strings(conn.Do("lrange", "list-key3", "0", "-1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("return strings:", retStrings)

}
