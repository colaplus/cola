package filter

import (
	"bufio"
	"io"
	"mercury/util"
	"os"
	"strings"
)

var (
	trie *util.Trie
)

func Init(filename string) (err error) {
	trie = util.NewTrie()
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		word, errRet := reader.ReadString('\n')
		word = strings.TrimSpace(word)
		if errRet == io.EOF {
			return
		}
		if errRet != nil {
			err = errRet
			return
		}

		err = trie.Add(word, nil)
		if err != nil {
			return
		}
	}

	return
}