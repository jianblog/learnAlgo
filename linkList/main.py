# coding = utf-8

import node

def basicOperation():
    # 创建一个包含5节点的单向链表
    thisLink = node.LinkNode()
    for s in "wor":
        thisLink.addNode(node.Node(s))
    print(thisLink)

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

if __name__ == "__main__":
    link =  basicOperation()
    #reverseLink(link)
    #print(link)


    # 回环
    loopLink = createLoop()
    checkLoop(loopLink)
