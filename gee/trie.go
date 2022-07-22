package gee

type node struct {
	pattern string

	part string

	children []*node

	isWild bool
}

func (*node) insert(pattern string, parts []string, height int) {

}

func (*node) search(parts []string, height int) *node {
	return nil
}
