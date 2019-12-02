package _06_linklist

/**
  算法练习： 单链表反转 go实现
  Author: Liujianbo
  Date: 2019-12-01
*/

import "fmt"

/**
* 链表节点, 单链表定义
 */

type ListNode struct {
	val  int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

// 方法定义
func (this *ListNode) getValue() int {
	return this.val
}

// 打印链表
func (this *LinkedList) Print() {
	cur := this.head
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.getValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

func NewListNode(v int) *ListNode {
	return &ListNode{v, nil}
}

func NewLinkedList(v int) *LinkedList {
	return &LinkedList{NewListNode(v)}
}

// 链表末尾增加节点
func (this *LinkedList) AddNode(v int) *LinkedList {
	newNode := NewListNode(v)
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	cur.next = newNode
	return this
}

// 链表翻转
func (this *LinkedList) ReverseLink() *LinkedList {
	var prev *ListNode = nil
	cur := this.head
	for cur.next != nil {
		last := cur.next
		cur.next = prev
		prev = cur
		cur = last
	}
	cur.next = prev
	this.head = cur
	return this
}
