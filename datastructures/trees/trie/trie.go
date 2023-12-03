package trie

import (
	"fmt"

	"github.com/mhrdini/godsa/datastructures/trees"
)

const trie = "Trie"

type Trie struct {
	root *TrieNode
	size int
}

type TrieNode struct {
	parent     string
	parentNode *TrieNode
	children   map[string]*TrieNode
}

func New() *Trie {
	return &Trie{
		root: nil,
		size: 0,
	}
}

func (t *Trie) Name() string {
	return trie
}

func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) Empty() bool {
	return t.root == nil
}

func (t *Trie) Values() []string {
	return []string{}
}

func (t *Trie) String() string {
	return fmt.Sprint(t.root)
}

func (t *Trie) Reset() {
	t = &Trie{
		root: nil,
		size: 0,
	}
}

func (n *TrieNode) Value() (value string, ok bool) {
	if n != nil {
		value = n.parent
		ok = true
	}
	return
}

func (n *TrieNode) Children() []trees.INode[string] {
	children := make([]trees.INode[string], len(n.children))
	for _, v := range n.children {
		children = append(children, v)
	}
	return children
}

func (n *TrieNode) IsNil() bool {
	return n == nil
}

func (t *Trie) Root() trees.INode[string] {
	return t.root
}

func (t *Trie) Insert(v string) {
	var sentinel string
	currentNode := t.root
	for _, char := range v {
		stringChar := string(char)
		// fmt.Printf("Inserting %v\n", stringChar)
		if t.root == nil {
			// fmt.Printf("> Root is empty\n")
			t.root = &TrieNode{
				parent:     sentinel,
				parentNode: nil,
				children:   make(map[string]*TrieNode),
			}
			newNode := &TrieNode{
				parent:     stringChar,
				parentNode: t.root,
				children:   make(map[string]*TrieNode),
			}
			t.root.children[stringChar] = newNode
			// fmt.Printf("> New root: %v\n", t.root)
			currentNode = newNode
			t.size++
		} else if childNode, ok := currentNode.children[stringChar]; ok {
			// fmt.Printf("> Found %v at current node %v\n", stringChar, currentNode)
			currentNode = childNode
		} else {
			newNode := &TrieNode{
				parent:     stringChar,
				parentNode: currentNode,
				children:   make(map[string]*TrieNode),
			}
			currentNode.children[stringChar] = newNode
			// fmt.Printf("> Created new key for %v in current node %v\n", stringChar, currentNode)
			currentNode = newNode
			t.size++
		}
	}
	// fmt.Println("")
	currentNode.children[sentinel] = nil
}

func (t *Trie) Search(v string) *TrieNode {
	currentNode := t.root
	for _, char := range v {
		stringChar := string(char)
		if t.root == nil {
			return nil
		} else if childNode, ok := currentNode.children[stringChar]; ok {
			currentNode = childNode
		} else {
			return nil
		}
	}
	return currentNode
}

func (t *Trie) Remove(v string) {

}

func (t *Trie) CollectAllWords() []string {
	return t.root.CollectAllWords("")
}

func (n *TrieNode) CollectAllWords(word string) (words []string) {
	var sentinel string
	for key, childNode := range n.children {
		if key == sentinel {
			words = append(words, word)
		} else {
			words = append(words, childNode.CollectAllWords(word+key)...)
		}
	}
	return
}

// Returns an array of all possible string endings to the input prefix
func (t *Trie) Autocomplete(prefix string) (suffixes []string) {
	currentNode := t.Search(prefix)
	if currentNode == nil {
		return
	}
	suffixes = currentNode.CollectAllWords("")
	return
}

// Returns a list of words that have the largest prefix as the input word
func (t *Trie) Autocorrect(word string) (suggestions []string) {
	prefix := word
	suffixes := t.Autocomplete(word)
	for len(suffixes) == 0 {
		// fmt.Printf("word: %v suffixes: %v\n", word, suffixes)
		prefix = prefix[:len(word)-1]
		suffixes = t.Autocomplete(word)
	}
	for _, suffix := range suffixes {
		suggestions = append(suggestions, prefix+suffix)
	}
	return
}

func newDemoTrie() *Trie {
	t := New()
	t.Insert("ace")
	t.Insert("act")
	t.Insert("bad")
	t.Insert("bake")
	t.Insert("bat")
	t.Insert("batter")
	t.Insert("cab")
	t.Insert("cat")
	t.Insert("catnap")
	t.Insert("catnip")
	return t
}

func DemoCollectAllWords() {
	t := newDemoTrie()
	words := t.CollectAllWords()
	fmt.Println(words)
}

func DemoAutocomplete() {
	t := newDemoTrie()
	words := t.Autocomplete("ba")
	fmt.Println(words)
}

func DemoAutocorrect() {
	t := newDemoTrie()
	words := t.Autocorrect("catnar")
	fmt.Println(words)
}
