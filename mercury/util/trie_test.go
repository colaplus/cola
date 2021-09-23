package util

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Add("黄色的", nil)
	trie.Add("黄色", nil)
	trie.Add("绿色", nil)
	trie.Add("蓝色", nil)

	result, str := trie.Check("我们这里有个黄色的灯泡，他存在了很久。", "***")

	fmt.Printf("result:%#v, str:%#v", result, str)
}

func TestTrie11(t *testing.T) {
	trie := NewTrie()
	trie.Add("黄色", nil)
	trie.Add("绿色", nil)
	trie.Add("蓝色", nil)

	result, str := trie.Check("狗东西。", "***")

	fmt.Printf("result:%#v, str:%#v", result, str)
}
