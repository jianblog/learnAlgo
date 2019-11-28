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


if __name__ == "__main__":
    link =  basicOperation()
    reverseLink(link)
    print(link)

