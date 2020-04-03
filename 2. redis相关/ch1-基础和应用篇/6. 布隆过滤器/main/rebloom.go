/*
 * Copyright © 2020 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
布隆过滤器可能会出现误判，
如果某个元素 bloom说存在 那不一定存在
如果某个元素 bloom说不存在 那一定不存在

bloom原理就是: 创建一个很大的位数组 和几个不一样的无偏hash函数(就是可以使元素分布均匀的hash函数)
每个值进来的时候 都会用这几个hash函数运算 得到的位置 在位数组上标记为1
当检查某个值是否存在的时候 同样用这几个hash函数晕眩 得到的位置上 如果都为1 返回存在 如果有一个为0 返回不存在
但是为1的位 可能是别的值得出的，所以判定存在时不一定准确 但是判定不存在时 一定准确
*/

func main() {

	conn, err := redis.Dial("tcp", "39.97.162.199:6379")
	if nil != err {
		fmt.Println("ERROR:", err)
	}
	defer conn.Close()

	conn.Do("del", "bloom1")
	conn.Do("del", "bloom2")
	conn.Do("del", "bloom3")

	//通过参数对bloom过滤器进行优化 查看误判率(即不存在的数据 判定为存在)
	//conn.Do("BF.RESERVE","bloom1","error_rate","0.01","initial_size","100")
	var false1 = 0
	for i := 0; i < 500; i++ {
		conn.Do("bf.add", "bloom1", fmt.Sprintf("usr%d", i))
	}
	for i := 0; i < 500; i++ {
		ret, _ := redis.Int(conn.Do("bf.exists", "bloom1", fmt.Sprintf("usr%d", 500+i)))
		if ret == 1 {
			//fmt.Println(i)
			false1++
		}
	}

	fmt.Printf("bloom1 error times %d, rate %f\n", false1, float64(false1)/500)

	str, _ := redis.String(conn.Do("BF.RESERVE", "bloom2", "0.001", "500"))
	fmt.Println(str)
	var false2 = 0
	for i := 0; i < 500; i++ {
		conn.Do("bf.add", "bloom2", fmt.Sprintf("usr%d", i))
	}
	for i := 0; i < 500; i++ {
		ret, _ := redis.Int(conn.Do("bf.exists", "bloom2", fmt.Sprintf("usr%d", 500+i)))
		if ret == 1 {
			//fmt.Println(i)
			false2++
		}
	}
	fmt.Printf("bloom2 error times %d, rate %f\n", false2, float64(false2)/500)

	str1, _ := redis.String(conn.Do("BF.RESERVE", "bloom3", "0.0001", "5"))
	fmt.Println(str1)
	var false3 = 0
	for i := 0; i < 500; i++ {
		conn.Do("bf.add", "bloom3", fmt.Sprintf("usr%d", i))
	}
	for i := 0; i < 500; i++ {
		ret, _ := redis.Int(conn.Do("bf.exists", "bloom3", fmt.Sprintf("usr%d", 500+i)))
		if ret == 1 {
			//fmt.Println(i)
			false3++
		}
	}
	fmt.Printf("bloom3 error times %d, rate %f\n", false3, float64(false3)/500)
}
