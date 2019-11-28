# coding = utf-8

import node

def newLink(s):
    thisLink = node.LinkNode()
    for i in str(s):
        thisLink.addNode(node.Node(i))
    return thisLink

def basicOperation():
    thisLink = newLink("world")
    # 追加节点
    # thisLink.addNode(node.Node("A"))
    # print(thisLink)

    # 头部插入节点
    # thisLink.insertNode("w", node.Node("1"))
    # print(thisLink)

    # 任意位插入
    # thisLink.insertNode("r", node.Node("2"))
    # print(thisLink)

    # 删除节点
    # thisLink.deleteNode("w")
    # print(thisLink)
    return thisLink


def reverseLink(linkNode):
    # 单链表反转
    # 使用三个指针 prev, cur, after

    prev = linkNode.head
    # 首先排除空链或单节点链
    if prev == None or prev.next == None:
        return linkNode
    else:
        cur = prev.next
        while cur:
            after = cur.next
            cur.next = prev
            prev = cur
            cur = after
        linkNode.head.next = None
        linkNode.head = prev

def createLoop():
    # 模拟创建一个包含回路的链表
    link = node.LinkNode()
    for i in "thisajoke":
        link.addNode(node.Node(i))

    pt = link.head
    pta = None
    while pt.value != "a":
        pt = pt.next
    pta = pt

    pt = link.head
    while pt.next:
        pt = pt.next
    pt.next = pta

    return link


def checkLoop(linkNode):
    # 检测链表是否存在环
    a = linkNode.head
    b = linkNode.head
    if a == None:
        return
    while a.next and a.next.next:
        a = a.next.next
        b = b.next
        if a == b:
            print("found loop")
            return
    return 

def mergeLink(la, lb):
    # 合并链表

    pta = la.head
    ptb = lb.head

    if not pta:
        return lb
    if not ptb:
        return la

    pNode = None
    p_tail  = None
    # 定位起始
    if pta.value < ptb.value:
        pNode = la
        p_tail = pta
        pta = pta.next
    else:
        pNode = lb
        p_tail = ptb
        ptb = ptb.next

    while True:
        if not pta:
            p_tail.next = ptb
            break
        if not ptb:
            p_tail.next = pta
            break

        if pta.value < ptb.value:
            p_tail.next = pta
            p_tail = pta
            pta = pta.next
        else:
            p_tail.next = ptb
            p_tail = ptb
            ptb = ptb.next
    return pNode


if __name__ == "__main__":
    link =  basicOperation()
    #reverseLink(link)
    #print(link)


    # 回环
    #loopLink = createLoop()
    #checkLoop(loopLink)


    # 链表合并
    la = newLink("ehr")
    lb = newLink("acnq")

    c = mergeLink(la, lb)
    print(c)
