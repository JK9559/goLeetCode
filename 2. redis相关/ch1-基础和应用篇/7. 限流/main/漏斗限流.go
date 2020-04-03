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
		totCap:   cap,
		rate:     rate,
		cap:      cap,
		leakTime: time.Now().UnixNano() / 1e6,
	}
}

func (f *Funnel) makeCap() {
	nowTime := time.Now().UnixNano() / 1e6
	durTime := nowTime - f.leakTime
	newCap := float64(durTime) / 1000 * f.rate
	//fmt.Println(durTime)
	if newCap < 1 {
		return
	}
	f.cap += newCap
	f.leakTime = nowTime
	if f.cap > f.totCap {
		f.cap = f.totCap
	}
}

func (f *Funnel) watering(unit float64) bool {
	f.makeCap()
	if f.cap >= unit {
		f.cap -= unit
		return true
	}
	return false
}

func isAllow(usrId, action string, f *Funnel) bool {
	//key := fmt.Sprintf("key:%s:%s", usrId, action)
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
