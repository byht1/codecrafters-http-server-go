package controllers

import (
	"strings"
)

type Node struct {
	children  map[string]*Node
	endpoints map[string]*Endpoint
	param     string
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	trie := Trie{root: &Node{
		children:  make(map[string]*Node),
		endpoints: make(map[string]*Endpoint),
	}}

	trie.AddRoute(root)
	trie.AddRoute(echo)

	return &trie
}

func (t *Trie) AddRoute(endpoint *Endpoint) {
	method := endpoint.Method
	path := endpoint.Path
	current := t.root
	segments := strings.Split(path, "/")

	for _, segment := range segments {
		if segment == "" {
			continue
		}

		var next *Node
		if segment[0] == ':' {
			param := segment[1:]
			if _, exists := current.children[param]; !exists {
				current.children[param] = &Node{
					children:  make(map[string]*Node),
					endpoints: make(map[string]*Endpoint),
					param:     param,
				}
			}
			next = current.children[param]
		} else {
			if _, exists := current.children[segment]; !exists {
				current.children[segment] = &Node{
					children:  make(map[string]*Node),
					endpoints: make(map[string]*Endpoint),
				}
			}
			next = current.children[segment]
		}
		current = next
	}

	current.endpoints[method] = endpoint
}

func (t *Trie) FindRoute(method, path string) (*Endpoint, map[string]string) {
	current := t.root
	segments := strings.Split(path, "/")
	params := make(map[string]string)

	for _, segment := range segments {
		if segment == "" {
			continue
		}

		if next, exists := current.children[segment]; exists {
			current = next
		} else {
			found := false
			for _, node := range current.children {
				if node.param != "" {
					params[node.param] = segment
					current = node
					found = true
					break
				}
			}
			if !found {
				return nil, nil
			}
		}
	}

	if endpoint, exists := current.endpoints[method]; exists {
		return endpoint, params
	}

	return nil, nil
}
