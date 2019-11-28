# coding = utf-8

import node

def basicOperation():
    # 创建一个包含5节点的单向链表
    thisLink = node.LinkNode()
    for s in "world":
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

if __name__ == "__main__":
    basicOperation()
