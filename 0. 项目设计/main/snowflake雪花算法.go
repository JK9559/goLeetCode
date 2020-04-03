package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	nodeBits  uint8 = 10
	stepBits  uint8 = 12
	nodeMax   int64 = -1 ^ (-1 << nodeBits)
	stepMax   int64 = -1 ^ (-1 << stepBits)
	timeShift       = nodeBits + stepBits
	nodeShift       = stepBits
)

var Epoch int64 = 1288834974657

type Node struct {
	mu        sync.Mutex
	timestamp int64
	node      int64
	step      int64
}

func NewNode(node int64) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}
	return &Node{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}

func (node *Node) Generate() int64 {
	node.mu.Lock()
	defer node.mu.Unlock()

	now := time.Now().UnixNano()

	if node.timestamp == now {
		node.step++

		if node.step > stepMax {
			for now <= node.timestamp {
				now = time.Now().UnixNano()
			}
		}
	} else {
		node.step = 0
	}

	node.timestamp = now

	res := (now-Epoch)<<timeShift | node.node<<nodeShift | node.step
	return res
}

func main() {
	node, err := NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int64)
	count := 100

	for i := 0; i < count; i++ {
		go func() {
			id := node.Generate()
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		fmt.Println(id)
		_, ok := m[id]
		if ok {
			fmt.Println("ID is not unique!")
			return
		}
		m[id] = i
	}
	fmt.Println("success!")
}
