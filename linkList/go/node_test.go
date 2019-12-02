package _06_linklist

import "testing"

func TestPrintLink(t *testing.T) {
	l := NewLinkedList(1)
	for i := 2; i < 6; i++ {
		l.AddNode(i)
	}
	l.Print()
}

func TestReverseLink(t *testing.T) {
	l := NewLinkedList(1)
	for i := 2; i < 6; i++ {
		l.AddNode(i)
	}
	l.Print()
	l = l.ReverseLink()
	l.Print()
}
