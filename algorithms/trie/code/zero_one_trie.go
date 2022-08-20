package main

// HIGH_BIT 数字的最大值，例如要表示的数字范围为 0 <= num <= (2^31 - 1)
const HIGH_BIT = 30

type Trie struct {
	// left 表示 0
	// right 表示 1
	left, right *Trie
}

func (t *Trie) Put(num int) {
	cur := t

	for i := HIGH_BIT; i >= 0; i-- {
		bit := (num >> i) & 1
		if bit == 0 {
			// 走 left
			if cur.left == nil {
				// 左子树不存在，新建
				cur.left = &Trie{}
			}
			cur = cur.left
		} else {
			// bit == 1, 走 right
			if cur.right == nil {
				// 右子树不存在，新建
				cur.right = &Trie{}
			}
			cur = cur.right
		}
	}
}

// Check 计算
func (t *Trie) Check(num int) (x int) {
	cur := t
	for i := HIGH_BIT; i >= 0; i-- {
		bit := (num >> i) & 1
		if bit == 0 {
			// a_i 的第 k 个二进制位为 0，应当往表示 1 的子节点 right 走
			if cur.right != nil {
				cur = cur.right
				x = x*2 + 1
			} else {
				cur = cur.left
				x = x * 2
			}
		} else {
			// bit == 1, a_i 的第 k 个二进制位为 1，应当往表示 0 的子节点 left 走
			if cur.left != nil {
				cur = cur.left
				x = x*2 + 1
			} else {
				cur = cur.right
				x = x * 2
			}
		}
	}
	return
}
