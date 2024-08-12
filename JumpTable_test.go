package Jumptable

import (
	"testing"
)

func TestJumpTable(t *testing.T) {
	jt := NewJumpTable()
	jt.InsertNode("a", 1)
	jt.InsertNode("b", 2)
	jt.InsertNode("c", 3)
	jt.Print()
	jt.UpdateNode("b", 5)
	jt.Print()
	jt.DeleteNode("a")
	jt.Print()
}
