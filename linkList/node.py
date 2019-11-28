#coding =utf-8

## 算法-单链表-利用数组实现LRU

class Node(object):
    def __init__(self, value):
        self.value = value
        self.next = None

    def setNext(self, pt):
        self.next = pt

    def __str__(self):
        return "node(%s)" % self.value

class LinkNode(object):
    def __init__(self):
        self.head = None

    def addNode(self, node):
        # 给链表添加追加
        if not self.head:
            self.head = node
            return
        pt = self.head
        while pt.next:
            pt = pt.next
        pt.next = node

    def insertNode(self, beforeValue, node):
        # 在指定节点查插入
        pt = self.head
        if pt.value == beforeValue:
            node.setNext(pt)
            self.head = node
            return
        while pt.next.value != beforeValue:
            pt = pt.next
        if pt.next:
            node.setNext(pt.next)
            pt.setNext(node)

    def deleteNode(self, value):
        # 删除指定值得节点
        pt = self.head
        if pt.value == value:
            self.head = pt.next
            return
        # 需要修改前一节点的next，所以遍历时的当前指向需在前一位，然后比较next指向的value
        while pt.next and pt.next.value != value:
            pt = pt.next
        if pt.next:
            pt.next = pt.next.next

    def __str__(self):
        if not self.head:
            return "None"
        else:
            pt = self.head
            ret = ""
            while pt:
                if pt.next:
                    ret += "%s -->" % pt.value
                else:
                    ret += "%s" % pt.value
                pt = pt.next
            return ret

        

    
