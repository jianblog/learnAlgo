package tree

type BST struct {
	*BinaryTree

	//比较方法
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BST {
	return &BST{BinaryTree: NewBinaryTree(rootV), compareFunc: compareFunc}
}

func (this *BST) Find(v interface{}) *Node {
	p := this.root
	for nil != p {
		compareResult := this.compareFunc(v, p.data)
		if compareResult == 0 {
			return p
		} else if compareResult > 0 {
			p = p.right
		} else {
			p = p.left
		}
	}
	return nil
}

func (this *BST) Insert(v interface{}) bool {
	p := this.root
	for nil != p {
		compareResult := this.compareFunc(v, p.data)
		if compareResult == 0 {
			return false
		} else if compareResult > 0 {
			if nil == p.right {
				p.right = NewNode(v)
				break
			}
			p = p.right
	 	} else {
			if nil == p.left {
				p.left = NewNode(v)
				break
			}
			p = p.left
		}
	}
	return true
}

func (this *BST) Delete(v interface{}) bool {
	var pp *Node = nil
	p := this.root
	deleteLeft := false
	for nil != p {
		compareResult := this.compareFunc(v, p.data)
		if compareResult > 0 {
			pp = p
			p = p.right
			deleteLeft = false
		} else if compareResult < 0 {
			pp = p
			p = p.left
			deleteLeft = true
		} else {
			// 找到待删除点p时，同时也得到了其父节点pp
			break
		}
	}

	if nil == p {
		return false
	} else if nil == p.left && nil == p.right {
		if nil != pp {
			if deleteLeft {
				pp.left = nil
			} else {
				pp.right = nil
			}
		} else {
			this.root = nil
		}
	} else if nil != p.right {
		pq := p
		q := p.right
		fromRight := true
		for nil != q.left {    // 寻找待删除节点右侧分支中最小的值
			pq = q
			q = q.left
			fromRight = false
		}
		// 替换节点的父节点删除下级关系,才可以将它移走
		if fromRight {
			pq.right = nil
		} else {
			pq.left = nil
		}
		// 替换节点取代被替换节点的关系链
		q.left = p.left
		q.right = p.right
		if nil == pp {
			this.root = p
		} else {
			if deleteLeft {
				pq.left = q
		 
