package gee

type node struct {
	pattern string

	part string

	children []*node

	isWild bool
}

func (*node) insert() {

}

func (*node) search() *node {
	return nil
}
