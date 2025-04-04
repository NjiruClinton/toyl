package trie

// TrieNode represents each node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie is the main data structure
type Trie struct {
	root *TrieNode
}

// NewTrie returns a new instance of Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// Insert adds a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search returns true if the word is in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith checks if any word in the Trie starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return true
}

//func main() {
//	trie := NewTrie()
//	trie.Insert("hello")
//	trie.Insert("world")
//
//	fmt.Println(trie.Search("hello"))    // true
//	fmt.Println(trie.Search("hell"))     // false
//	fmt.Println(trie.StartsWith("hell")) // true
//	fmt.Println(trie.StartsWith("wo"))   // true
//	fmt.Println(trie.Search("hi"))       // false
//}
