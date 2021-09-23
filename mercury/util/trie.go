package util

type Node struct {
	//rune表示一个utf8字符（中文字符）
	char   rune
	Data   interface{}
	parent *Node
	Depth  int
	//childs 当前节点的所有孩子节点
	childs map[rune]*Node
	term   bool
}

type Trie struct {
	root *Node
	size int
}

func NewNode() *Node {
	return &Node{
		childs: make(map[rune]*Node, 32),
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

//假如我要把敏感词：“卧槽”加入
//Add("卧槽", nil)
//Add("色情片", nil)
func (p *Trie) Add(key string, data interface{}) (err error) {
	node := p.root
	runes := []rune(key)
	for _, r := range runes {
		ret, ok := node.childs[r]
		if !ok {
			ret = NewNode()
			ret.Depth = node.Depth + 1
			ret.char = r
			node.childs[r] = ret
		}
		node = ret
	}

	node.term = true
	node.Data = data
	return
}

//findNode("卧槽")
func (p *Trie) findNode(key string) (result *Node) {
	node := p.root
	runes := []rune(key)
	for _, v := range runes {
		ret, ok := node.childs[v]
		if !ok {
			return
		}
		node = ret
	}

	result = node
	return
}

func (p *Trie) Check(text, replace string) (isHit bool, str string) {
	chars := []rune(text)
	if p.root == nil {
		return
	}

	var left []rune
	node := p.root
	start := 0
	for index, v := range chars {
		ret, ok := node.childs[v]
		if !ok {
			left = append(left, chars[start:index+1]...)
			start = index + 1
			node = p.root
			continue
		}

		node = ret
		if ret.term {
			isHit = true
			node = p.root
			left = append(left, []rune(replace)...)
			start = index + 1
			continue
		}
	}

	str = string(left)

	return
}
