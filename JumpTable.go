package Jumptable

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type JumpNode struct {
	key      string
	value    int
	NextNode []*JumpNode
}

var (
	cntHigh = 1
	root    = &JumpNode{value: math.MinInt, NextNode: make([]*JumpNode, 1)}
	first   = root.NextNode[0]
	tail    = root.NextNode[0]
	memHash = make(map[string]int)
)

func SearchNode(leftVal, rightVal int) []*JumpNode {
	result := searchNode(leftVal, rightVal, 0)
	res := make([]*JumpNode, 0)
	start := result.NextNode[0]
	for start != nil && start.value <= rightVal {
		res = append(res, start)
		start = start.NextNode[0]
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

// 1.无需进行单点查询的实现，只需实现范围查询，单点查询通过范围查询模拟。2.查询函数需要传入一个lowindex，用来表示你希望查询到第几层，方便后续insert
func searchNode(leftVal, rightVal, lowIndex int) *JumpNode {
	if first == nil {
		return root
	}
	if rightVal < first.value {
		return root
	}
	if leftVal > tail.value {
		return tail
	}
	leftNode, index := root, cntHigh-1
	for index >= lowIndex {
		if leftNode.NextNode[index] != nil && leftNode.NextNode[index].value < leftVal {
			leftNode = leftNode.NextNode[index]
		} else {
			index--
		}
	}
	return leftNode
}

// 考虑到并发还需要再改
func InsertNode(key string, val int) string {
	if _, ok := memHash[key]; ok {
		return "error,key is exist"
	}
	node := &JumpNode{key: key, value: val, NextNode: make([]*JumpNode, 1)}
	result := searchNode(val, val, 0)
	//fmt.Printf("%+v", result)
	if first == nil {
		first = node
		tail = node
	}
	if result == root {
		first = node
	}
	if result == tail {
		tail = node
	}
	result.NextNode[0], node.NextNode[0] = node, result.NextNode[0]
	//fmt.Println(1)
	rand.Seed(time.Now().UnixNano())
	cnt := 1
	for rand.Float64() < 0.5 {
		cnt++
		node.NextNode = append(node.NextNode, nil)
		if cnt > cntHigh {
			root.NextNode = append(root.NextNode, node)
		} else {
			r := searchNode(val, val, cnt-1)
			r.NextNode[cnt-1], node.NextNode[cnt-1] = node, r.NextNode[cnt-1]
		}
	}
	memHash[key] = val
	cntHigh = max(cnt, cntHigh)
	return "ok"
}

func DeleteNode(key string) string {
	if _, ok := memHash[key]; !ok {
		return "error,key is not exist"
	}
	result := searchNode(memHash[key], memHash[key], 0)
	for result.NextNode[0].key != key {
		result = result.NextNode[0]
	}
	p := result.NextNode[0]
	result.NextNode[0] = p.NextNode[0]
	for i := 1; i < len(p.NextNode); i++ {
		r := searchNode(memHash[key], memHash[key], i)
		for r.NextNode[i].key != key {
			r = r.NextNode[i]
		}
		r.NextNode[i] = p.NextNode[i]
	}
	delete(memHash, key)
	return "ok"
}

func UpdateNode(key string, val int) string {
	if _, ok := memHash[key]; !ok {
		return "error,key is not exist"
	}
	DeleteNode(key)
	InsertNode(key, val)
	return "ok"
}

func Print() {
	p := root
	for p.NextNode[0] != nil {
		for i := 0; i < len(p.NextNode[0].NextNode); i++ {
			if p.NextNode[0].NextNode[i] != nil {
				fmt.Printf("key:%v,val:%v,", p.NextNode[0].NextNode[i].key, p.NextNode[0].NextNode[i].value)
			}
		}
		if p == nil {
			break
		}
		p = p.NextNode[0]
		fmt.Println()
	}
	fmt.Println()
}
