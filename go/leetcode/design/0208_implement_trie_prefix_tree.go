package design

type TrieNode struct {
	isWord   bool
	next map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	next := make(map[rune]*TrieNode)
	return &TrieNode{next:next}
}

type Trie struct {
	root *TrieNode
	size int
}

func NewTrie() Trie {
	return Trie{NewTrieNode(), 0}
}

func (this *Trie) Insert(word string) {
	runes := []rune(word)
	cur := this.root
	for i := range runes {
		r := runes[i]
		if cur.next[r] == nil {
			cur.next[r] = NewTrieNode()
		}
		cur = cur.next[r]
	}

	if !cur.isWord {
		cur.isWord = true
		this.size++
	}
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	runes := []rune(word)
	for i := range runes {
		r := runes[i]
		if cur.next[r] == nil {
			return false
		}
		cur = cur.next[r]
	}
	return cur.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	runes := []rune(prefix)
	for i := range runes {
		r := runes[i]
		if cur.next[r] == nil {
			return false
		}
		cur = cur.next[r]
	}
	return true
}
