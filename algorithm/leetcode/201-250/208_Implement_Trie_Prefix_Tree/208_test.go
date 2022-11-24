package _08_Implement_Trie_Prefix_Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Trie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	assert.EqualValues(t, true, trie.Search("apple"))
	assert.EqualValues(t, false, trie.Search("app"))
	assert.EqualValues(t, true, trie.StartsWith("app"))
	trie.Insert("app")
	assert.EqualValues(t, true, trie.Search("app"))
}

func Test_Trie1(t *testing.T) {
	trie := Constructor()
	trie.Insert("a")
	assert.EqualValues(t, true, trie.Search("a"))
	assert.EqualValues(t, true, trie.StartsWith("a"))
}
