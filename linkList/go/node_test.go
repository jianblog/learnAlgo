package _06_linklist

import (
	"fmt"
	"testing"
)

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
	fmt.Println("hello world")
	l.Print()
	l = l.ReverseLink()
	l.Print()
}

func TestIntersection(t *testing.T) {
	la := NewLinkedList(10)
	lb := NewLinkedList(1)
	pta, ptb := la.head, lb.head
	la.Print()
	lb.Print()
	for pta != ptb {
		if pta == nil {
			pta = lb.head
		} else {
			fmt.Println(pta.val)
			pta = pta.next
		}
		if ptb == nil {
			ptb = la.head
		} else {
			ptb = ptb.next
		}
	}
	fmt.Println(pta.val)
}
