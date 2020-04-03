/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	conn, err := redis.Dial("tcp", "39.97.162.199:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()

	// string-del
	keys, err := redis.Strings(conn.Do("keys", "string-*"))
	fmt.Println(keys)
	var delCmd []interface{}
	for _, v := range keys {
		delCmd = append(delCmd, v)
	}
	k, err := redis.Int(conn.Do("del", delCmd...))
	fmt.Println(k)

	// string-set
	fmt.Println("=====>set<=====")
	_, err = conn.Do("set", "string-key1", "val1")
	if err != nil {
		fmt.Println("redis set error:", err)
	}

	// string-get
	fmt.Println("=====>get<=====")
	name, err := redis.String(conn.Do("get", "string-key1"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}

	// string-exists
	fmt.Println("=====>exists<=====")
	ret, err := redis.Int(conn.Do("exists", "string-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("Return Code:", ret)
	ret1, err := redis.Int(conn.Do("exists", "string-key2"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("Return Code:", ret1)

	// string-mset
	fmt.Println("=====>mset<=====")
	m := make(map[string]string, 5)
	m["string-mkey1"] = "mval1"
	m["string-mkey2"] = "mval2"
	m["string-mkey3"] = "mval3"
	var mCmd []interface{}
	for k, v := range m {
		mCmd = append(mCmd, k)
		mCmd = append(mCmd, v)
	}
	retS, err := redis.String(conn.Do("mset", mCmd...))
	fmt.Println(mCmd)
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println("MSET Return Code:", retS)

	// string-mget
	fmt.Println("=====>mget<=====")
	var sMget []interface{}
	for k := range m {
		sMget = append(sMget, k)
	}
	retSlice, err := redis.Strings(conn.Do("mget", sMget...))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(retSlice)

	// string-expire
	fmt.Println("=====>expire<=====")
	retS, err = redis.String(conn.Do("set", "string-expire-key1", "expire-val1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	retS, err = redis.String(conn.Do("get", "string-expire-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(retS)

	retS, err = redis.String(conn.Do("expire", "string-expire-key1", "1"))
	fmt.Println(retS)

	time.Sleep(1 * time.Second)
	fmt.Println("After a second")
	retS, err = redis.String(conn.Do("get", "string-expire-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(retS)

	// string-setex
	fmt.Println("=====>setex<=====")
	retS, err = redis.String(conn.Do("setex", "string-setex-key1", "1", "setex-val1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	retS, err = redis.String(conn.Do("get", "string-setex-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(retS)
	time.Sleep(1 * time.Second)
	retS, err = redis.String(conn.Do("get", "string-setex-key1"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(retS)

	// string-setnx
	fmt.Println("=====>setnx<=====")
	//_, err = conn.Do("del", "setnx-key1")
	retS, err = redis.String(conn.Do("get", "string-setnx-key1"))
	fmt.Println(retS)
	ret, err = redis.Int(conn.Do("setnx", "string-setnx-key1", "setnx-val1"))
	if err != nil {
		fmt.Println("Err before setnx :", err)
	}
	fmt.Println("ret: ", ret)
	ret, err = redis.Int(conn.Do("setnx", "string-setnx-key1", "setnx-val1"))
	if err != nil {
		fmt.Println("Err after setnx :", err)
	}
	fmt.Println("ret: ", ret)

	// string-incrby
	fmt.Println("=====>incr<=====")
	_, err = conn.Do("set", "string-age", "30")
	ret, err = redis.Int(conn.Do("incrby", "string-age", "5"))
	fmt.Println(ret)

	retS, err = redis.String(conn.Do("get", "string-age"))
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(retS)

}
