package filter

func Replace(text string, replace string) (result string, isHit bool) {
	isHit, result = trie.Check(text, replace)
	return
}
