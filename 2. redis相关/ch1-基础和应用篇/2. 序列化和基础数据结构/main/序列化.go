/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    uint   `json:"age"`
}

func main() {
	var usr = &User{
		Age:    15,
		Gender: "boy",
		Name:   "Bob",
		Id:     "109920",
	}

	str, err := json.Marshal(usr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str))

	conn, err := redis.Dial("tcp", "39.97.162.199:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()

	rsp, err := redis.String(conn.Do("set", usr.Id, string(str)))
	fmt.Println(rsp)

	rspData, err := redis.String(conn.Do("get", usr.Id))
	var rspUsr User
	err = json.Unmarshal([]byte(rspData), &rspUsr)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println(rspUsr.Name)

}
