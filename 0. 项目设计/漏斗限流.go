/*
 * Copyright © 2020 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type Funnel struct {
	totCap   float64
	rate     float64
	cap      float64
	leakTime int64
}

func newFunnel(cap, rate float64) *Funnel {
	return &Funnel{
		// 漏斗总容量
		totCap: cap,
		// 漏水速录 每秒钟几个容量
		rate: rate,
		// 当前剩余容量
		cap: cap,
		// 上次漏水时间 毫秒
		leakTime: time.Now().UnixNano() / 1e6,
	}
}

func (f *Funnel) makeCap() {
	// 获得当前时间
	nowTime := time.Now().UnixNano() / 1e6
	// 距离上次漏水 过去了多久
	durTime := nowTime - f.leakTime
	// 这次可漏掉多少水
	newCap := float64(durTime) / 1000 * f.rate
	//fmt.Println(durTime)
	if newCap < 1 {
		return
	}
	// 更新当前剩余容量 和 漏水时间
	f.cap += newCap
	f.leakTime = nowTime
	// 最大剩余容量 为最大容器容量
	if f.cap > f.totCap {
		f.cap = f.totCap
	}
}

func (f *Funnel) watering(unit float64) bool {
	// 为了注入unit的水 需要先看看经历了这段时间 漏斗腾出了多少空间
	f.makeCap()
	// 判断腾出的空间是否够此次注入的
	if f.cap >= unit {
		f.cap -= unit
		return true
	}
	return false
}

func isAllow(usrId, action string, f *Funnel) bool {
	//key := fmt.Sprintf("key:%s:%s", usrId, action)
	// 每次请求都是向漏斗里注入1个cap的水
	return f.watering(1)
}

func main() {
	// 同一个用户 同一个动作 开始可以直接访问 15/0.5 次 之后每两秒能访问一次
	fun := newFunnel(1, 0.5)
	for i := 0; i < 20; i++ {
		fmt.Println(isAllow("jk9559", "reply", fun))
		time.Sleep(time.Second)
	}
	fmt.Println(">=====<")
	conn, err := redis.Dial("tcp", "39.97.162.199:6379")
	if nil != err {
		fmt.Println("Redis ERROR:", err)
		return
	}
	defer conn.Close()

	conn.Do("del", "key:jk9559:ans")
	for i := 0; i < 20; i++ {
		res, err := redis.Int64s(conn.Do("cl.throttle", "key:jk9559:ans", 0, 30, 60, 1))
		if nil != err {
			fmt.Println("ERROR is : ", err)
			break
		}
		fmt.Println(res)
		fmt.Println(res[0] == 0)
		time.Sleep(time.Second)
	}
}
