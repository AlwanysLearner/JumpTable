package Jumptable

import (
	"fmt"
	"testing"
)

func TestJumpTable(t *testing.T) {
	for i := 0; i < 10; i++ {
		InsertNode(fmt.Sprint(i), i)
		Print()
	}
	fmt.Println(DeleteNode(fmt.Sprint(7)))
	Print()
	UpdateNode("1", 5)
	Print()
	res := SearchNode(3, 6)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%+v", res[i])
	}
}
