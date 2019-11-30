#!/usr/bin/env python3
# coding = utf-8

###
#  算法学习 - 数组构成队列
#  coding by Liujianbo
#  Date: 2019-11-30
###
import sys
sys.path.append("..")
from linkList import node

class LinkQueue(object):
    def __init__(self):
        self.head = None
        self.tail = None

    def dequeue(self):
        if not self.head:
            return False
        ret = self.head
        self.head = self.head.next
        return ret


    def enqueue(self, value):
        if not self.head:
            self.head = node.Node(value)
            self.tail = self.head
            return
        self.tail.next = node.Node(value)
        self.tail = self.tail.next

    def __str__(self):
        pt = self.head
        s = ""
        while pt:
            s += pt.value + "-"
            pt = pt.next
        return s

