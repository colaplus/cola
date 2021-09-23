package filter

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	err := Init("../data/filter.dat.txt")
	if err != nil {
		t.Errorf("load filter data failed, err:%v", err)
		return
	}
	trie.Add("黄色", nil)

	data := `
		asdfa阿萨德反服阿萨德发送到阿萨德fuck发送到乱伦阿斯蒂芬裸体，喜欢小黄片，黄色的
`
	result, isReplace := Replace(data, "***")
	fmt.Printf("isReplace:%#v, str:%v\n", isReplace, result)
}