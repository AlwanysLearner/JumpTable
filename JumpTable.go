package Jumptable

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type JumpNode struct {
	key      string
	value    int
	NextNode []*JumpNode
}

type JumpTable struct {
	root    *JumpNode
	first   *JumpNode
	tail    *JumpNode
	memHash map[string]int
	cntHigh int
	mutex   sync.RWMutex
}

func NewJumpTable() *JumpTable {
	rand.Seed(time.Now().UnixNano())
	root := &JumpNode{value: math.MinInt, NextNode: make([]*JumpNode, 1)}
	return &JumpTable{
		root:    root,
		first:   root.NextNode[0],
		tail:    root.NextNode[0],
		memHash: make(map[string]int),
		cntHigh: 1,
	}
}

func (jt *JumpTable) SearchNode(leftVal, rightVal int) []*JumpNode {
	jt.mutex.RLock()
	defer jt.mutex.RUnlock()

	result := jt.searchNode(leftVal, rightVal, 0)
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

func (jt *JumpTable) searchNode(leftVal, rightVal, lowIndex int) *JumpNode {
	if jt.first == nil {
		return jt.root
	}
	if rightVal < jt.first.value {
		return jt.root
	}
	if leftVal > jt.tail.value {
		return jt.tail
	}
	leftNode, index := jt.root, jt.cntHigh-1
	for index >= lowIndex {
		if leftNode.NextNode[index] != nil && leftNode.NextNode[index].value < leftVal {
			leftNode = leftNode.NextNode[index]
		} else {
			index--
		}
	}
	return leftNode
}

func (jt *JumpTable) InsertNode(key string, val int) string {
	jt.mutex.Lock()
	defer jt.mutex.Unlock()

	if _, ok := jt.memHash[key]; ok {
		return "error, key already exists"
	}
	node := &JumpNode{key: key, value: val, NextNode: make([]*JumpNode, 1)}
	result := jt.searchNode(val, val, 0)
	if jt.first == nil {
		jt.first = node
		jt.tail = node
	}
	if result == jt.root {
		jt.first = node
	}
	if result == jt.tail {
		jt.tail = node
	}
	result.NextNode[0], node.NextNode[0] = node, result.NextNode[0]

	cnt := 1
	for rand.Float64() < 0.5 {
		cnt++
		node.NextNode = append(node.NextNode, nil)
		if cnt > jt.cntHigh {
			jt.root.NextNode = append(jt.root.NextNode, node)
		} else {
			r := jt.searchNode(val, val, cnt-1)
			r.NextNode[cnt-1], node.NextNode[cnt-1] = node, r.NextNode[cnt-1]
		}
	}
	jt.memHash[key] = val
	jt.cntHigh = max(cnt, jt.cntHigh)
	return "ok"
}

func (jt *JumpTable) DeleteNode(key string) string {
	jt.mutex.Lock()
	defer jt.mutex.Unlock()

	if _, ok := jt.memHash[key]; !ok {
		return "error, key does not exist"
	}
	result := jt.searchNode(jt.memHash[key], jt.memHash[key], 0)
	for result.NextNode[0].key != key {
		result = result.NextNode[0]
	}
	p := result.NextNode[0]
	result.NextNode[0] = p.NextNode[0]
	for i := 1; i < len(p.NextNode); i++ {
		r := jt.searchNode(jt.memHash[key], jt.memHash[key], i)
		for r.NextNode[i].key != key {
			r = r.NextNode[i]
		}
		r.NextNode[i] = p.NextNode[i]
	}
	delete(jt.memHash, key)
	return "ok"
}

func (jt *JumpTable) UpdateNode(key string, val int) string {
	jt.mutex.RLock()
	if _, ok := jt.memHash[key]; !ok {
		jt.mutex.RUnlock()
		return "error, key does not exist"
	}
	jt.mutex.RUnlock()
	jt.DeleteNode(key)
	jt.InsertNode(key, val)
	return "ok"
}

func (jt *JumpTable) Print() {
	jt.mutex.RLock()
	defer jt.mutex.RUnlock()

	for level := jt.cntHigh - 1; level >= 0; level-- {
		fmt.Printf("Level %d: ", level)
		node := jt.root
		for node != nil {
			if level < len(node.NextNode) {
				fmt.Printf("(%s, %d) ", node.key, node.value)
				node = node.NextNode[level]
			} else {
				node = nil
			}
		}
		fmt.Println()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
