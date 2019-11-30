#!/usr/bin/env python3
# coding =utf-8

###
#  算法学习 - 数组构成队列
#  coding by Liujianbo
#  Date: 2019-11-29
###

class ArrayQueue(object):
    def __init__(self, capacity: int):
        self.n = capacity
        self.head = 0
        self.tail = 0
        self.items = []

    def enqueue(self, item):
        # 入队
        if self.tail == self.n:
            # 队列全满 或 头部需要释放空间
            if self.head == 0:
                return False
            else:
                for i in range(self.head, self.tail):
                    self.items[i-self.head] = self.item[i]
                self.tail = self.tail - self.head
                self.head = 0

        self.items[tail] = item
        self.tail += 1
        return True

    def dequeue(self):
        # 出队
        if self.head == self.tail:
            return None
        ret = self.items[self.head]
        head += 1
        return ret
