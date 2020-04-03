/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

/*
hash list当元素个数超512时，或者 key/value 等元素长度超过64，数据结构由ziplist->标准结构存储
zset 元素超过128时，或者元素长度大于64，数据结构由ziplist->标准结构存储
*/

func main() {
	conn, err := redis.Dial("tcp", "39.97.162.199:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	_, _ = conn.Do("del", "zipObjTest")

	// hash 超过512元素
	_ = conn.Send("multi")

	for i := 0; i < 512; i++ {
		_ = conn.Send("hset", "zipObjTest", strconv.Itoa(i), strconv.Itoa(i))
	}
	_ = conn.Flush()
	_, _ = conn.Do("exec")

	str, err := redis.String(conn.Do("object", "encoding", "zipObjTest"))

	fmt.Println(str)

	_, _ = conn.Do("hset", "zipObjTest", "512", "512")
	str, err = redis.String(conn.Do("object", "encoding", "zipObjTest"))

	fmt.Println(str)

	// set当存储元素均为数字 intset 有字符串时 为hashtable
	_, _ = conn.Do("del", "zipObjTest")

	_, _ = conn.Do("sadd", "zipObjTest", "1", "2", "3")
	str, err = redis.String(conn.Do("object", "encoding", "zipObjTest"))

	fmt.Println(str)
	_, _ = conn.Do("sadd", "zipObjTest", "1", "2", "hah")
	str, err = redis.String(conn.Do("object", "encoding", "zipObjTest"))

	fmt.Println(str)

	// zset

}
