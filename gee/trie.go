package gee

import "strings"

type node struct {
	pattern string

	part string

	children []*node

	isWild bool
}

func (n *node) getFirstMatchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

func (n *node) getAllMatchChildren(part string) []*node {
	children := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}

	return children
}

//parts[height] 是要插入的下一个
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]

	child := n.getFirstMatchChild(part)
	if child == nil {
		child = &node{
			pattern:  "",
			part:     part,
			children: make([]*node, 0),
			isWild:   part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	//匹配到最后一个 part 或者匹配到第一个 * 就结束匹配
	if height == len(parts) || strings.HasPrefix(n.part, "*") {
		if n.part == "" {
			return nil
		}

		return n
	}

	part := parts[height]

	children := n.getAllMatchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)

		if result != nil {
			return result
		}
	}

	return nil
}
