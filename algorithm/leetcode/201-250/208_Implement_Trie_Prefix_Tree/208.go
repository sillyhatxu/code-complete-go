package _08_Implement_Trie_Prefix_Tree

type Trie struct {
	children [26]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	current := this
	for i := range word {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &Trie{}
		}
		current = current.children[index]
	}
	current.isWord = true
}

func (this *Trie) Search(word string) bool {
	current := this
	for i := range word {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for i := range prefix {
		index := prefix[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
