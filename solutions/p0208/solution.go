package solution

// According to notes the solution has some simplifications like using only english alphabet and do not support unicode

// NOTES:
// You may assume that all inputs are consist of lowercase letters a-z.
// All inputs are guaranteed to be non-empty strings.

// Trie ...
type Trie struct {
	root *trieNode
}

type trieNode struct {
	childrens [26]*trieNode
	isWordEnd bool
}

// Constructor ...
func Constructor() Trie {
	return Trie{root: &trieNode{}}
}

// Insert ...
func (tr *Trie) Insert(word string) {
	wl := len(word)
	curr := tr.root
	for i := 0; i < wl; i++ {
		index := word[i] - 'a'
		if curr.childrens[index] == nil {
			curr.childrens[index] = &trieNode{}
		}
		curr = curr.childrens[index]

	}
	curr.isWordEnd = true
}

// Search ...
func (tr *Trie) Search(word string) bool {
	wl := len(word)
	curr := tr.root
	for i := 0; i < wl; i++ {
		index := word[i] - 'a'
		if curr.childrens[index] == nil {
			return false
		}
		curr = curr.childrens[index]
	}
	if curr.isWordEnd {
		return true
	}
	return false
}

// StartsWith ...
func (tr *Trie) StartsWith(prefix string) bool {
	pl := len(prefix)
	curr := tr.root
	for i := 0; i < pl; i++ {
		c := prefix[i] - 'a'
		if curr.childrens[c] == nil {
			return false
		}
		curr = curr.childrens[c]
	}
	return true
}
