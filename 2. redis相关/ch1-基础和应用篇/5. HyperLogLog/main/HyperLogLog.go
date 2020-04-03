/*
 * Copyright © 2020 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
hyperLogLog 有去重的功能 将同一个key下的值进行了去重 只不过可能结果会有些偏差
会占用12KB的空间，所以不适合用来统计单个用户的信息

HLL原理理解：
首先是概念：伯努利试验，一次试验为硬币最终朝上，记k为未成功次数。k_max为总次数为n的伯努利试验中最大的k
类比该算法，可以计一个数字从右到左首次出现1的位置，可以通过这个数字来估计本组的试验总量。

HLL有很多个桶，每个桶最终会得到一个k_max，并且以此来估计这个桶内的数量，
最后求得调和平均数, m/(1/a1 + 1/a2 +...+ 1/am) 最终的个数为 const * m * 平均值


*/

func main() {

	conn, err := redis.Dial("tcp", "39.97.162.199:6379")
	if nil != err {
		fmt.Println("ERROR: ", err)
		return
	}

	defer conn.Close()

	_ = conn.Send("multi")
	for i := 0; i < 100; i++ {
		redis.String(conn.Do("pfadd", "myUser3", fmt.Sprintf("usr%d", i)))
	}
	conn.Do("exec")

	res, err := redis.Int64(conn.Do("pfcount", "myUser3"))
	if nil != err {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(res)

}
