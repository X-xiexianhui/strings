package main

// 三向单词查找树
type TST struct {
	Root *tstNode //根节点
}
type tstNode struct {
	c                byte     //字符
	left, mid, right *tstNode //左中右子三向单词查找树
	value            any      //值
}

func newTstNode(c byte) *tstNode {
	return &tstNode{c: c}
}

func (t *TST) Get(key string) any {
	x := t.get(t.Root, key, 0)
	if x == nil {
		return nil
	}
	return x.value

}
func (t *TST) get(x *tstNode, key string, d int) *tstNode {
	if x == nil {
		return nil
	}
	c := key[d]
	if c < x.c {
		return t.get(x.left, key, d)
	} else if c > x.c {
		return t.get(x.right, key, d)
	} else if d < len(key)-1 {
		return t.get(x.mid, key, d+1)
	} else {
		return x
	}
}

func (t *TST) Put(key string, val any) {
	t.Root = t.put(t.Root, key, val, 0)
}
func (t *TST) put(x *tstNode, key string, val any, d int) *tstNode {
	c := key[d]
	if x == nil {
		x = newTstNode(c)
	}
	if c < x.c {
		x.left = t.put(x.left, key, val, d)
	} else if c > x.c {
		x.right = t.put(x.right, key, val, d)
	} else if d < len(key)-1 {
		x.mid = t.put(x.mid, key, val, d+1)
	} else {
		x.value = val
	}
	return x

}
