package main

// R 含有 R 个字符的字母表
const R = 26

type RTrie struct {
	val  int
	next [R]*RTrie
}

func (t *RTrie) Get(key string) int {
	node := t
	for _, ch := range key {
		ch -= 'a'
		if node.next[ch] == nil {
			return 0
		}
		node = node.next[ch]
	}
	return node.val
}

func (t *RTrie) Put(key string, val int) {
	node := t
	for _, ch := range key {
		ch -= 'a'
		if node.next[ch] == nil {
			node.next[ch] = &RTrie{}
		}
		node = node.next[ch]
	}
	node.val = val
}

//func (root *RTrie) String() string {
//	node := root
//	for node != {
//
//	}
//}
