package main

type TireSt struct {
	R    int   //基数
	Root *node //单词查找树
}

type node struct {
	val  any
	next []*node
}

func newNode(val any, R int) *node {
	return &node{
		val:  val,
		next: make([]*node, R),
	}
}

func NewTireSt(R int) *TireSt {
	return &TireSt{R: R}
}
func (t *TireSt) Get(key string) any {
	x := t.get(t.Root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}
func (t *TireSt) get(root *node, key string, d int) *node {
	if root == nil {
		return nil
	}
	if d == len(key) {
		return root
	}
	c := key[d]
	return t.get(root.next[c], key, d+1)
}

func (t *TireSt) Put(key string, val any) {
	t.Root = t.put(t.Root, key, val, 0)
}

func (t *TireSt) put(root *node, key string, val any, d int) *node {
	if root == nil {
		root = newNode(nil, t.R)
	}
	if d == len(key) {
		root.val = val
		return root
	}
	c := key[d]
	root.next[c] = t.put(root.next[c], key, val, d+1)
	return root
}
func (t *TireSt) Keys() []string {
	return t.KeyWithPrefix("")
}
func (t *TireSt) KeyWithPrefix(key string) []string {
	q := make([]string, 0)
	t.collect(t.Root, key, "", &q)
	return q
}

//	func (t *TireSt) collect(x *node, pre string, q *[]string) {
//		if x == nil {
//			return
//		}
//		if x.val != nil {
//			*q = append(*q, pre)
//		}
//		for i := 0; i < t.R; i++ {
//			t.collect(x.next[i], pre+string(rune(i)), q)
//		}
//	}
func (t *TireSt) KeysThatMatch(pat string) []string {
	q := make([]string, 0)
	t.collect(t.Root, "", pat, &q)
	return q
}
func (t *TireSt) collect(x *node, pre, pat string, q *[]string) {
	d := len(pre)
	if x == nil {
		return
	}
	if d == len(pat) && x.val != nil {
		*q = append(*q, pre)
	}
	if d == len(pat) {
		return
	}
	c := pat[d]
	for i := 0; i < t.R; i++ {
		if c == '.' || c == byte(i) {
			t.collect(x.next[i], pre+string(rune(i)), pat, q)
		}
	}

}

func (t *TireSt) LongestPrefixOf(prefix string) string {
	length := t.search(t.Root, prefix, 0, 0)
	return prefix[0:length]
}

func (t *TireSt) search(x *node, s string, d, length int) int {
	if x == nil {
		return length
	}
	if x.val != nil {
		length = d
	}
	if d == len(s) {
		return length
	}
	c := s[d]
	return t.search(x.next[c], s, d+1, length)
}

func (t *TireSt) Delete(key string) {
	t.Root = t.delete(t.Root, key, 0)
}
func (t *TireSt) delete(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		x.val = nil
	} else {
		c := key[d]
		x.next[c] = t.delete(x.next[c], key, d+1)
	}
	if x.val != nil {
		return x
	}
	for i := 0; i < t.R; i++ {
		if x.next[i] != nil {
			return x
		}
	}
	return nil
}
