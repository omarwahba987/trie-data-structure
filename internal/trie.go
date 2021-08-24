package internal

import (
	"errors"
	"fmt"
)

type Node struct {
	Value    string
	Children map[string]Node
	Address  string
}

func (node Node) String() string {
	return fmt.Sprintf("\n node value : %v \n node children : %v  \n ",
		node.Value, node.Children)
}

type Trie struct {
	Root Node
}

func InitTrie() *Trie {
	return &Trie{
		Root: Node{
			Value:    "",
			Children: map[string]Node{},
		},
	}
}
func (trie *Trie) Insert(name, address string) error {
	currentNode := trie.Root
	for pos, i := range name {
		if _, ok := currentNode.Children[string(i)]; ok {
			currentNode = currentNode.Children[string(i)]
		} else {
			if pos == len(name)-1 {
				currentNode.Children[string(i)] = Node{
					Value:    name[0 : pos+1],
					Address:  address,
					Children: map[string]Node{},
				}
				return nil
			} else {
				currentNode.Children[string(i)] = Node{
					Value:    name[0 : pos+1],
					Children: map[string]Node{},
				}
				currentNode = currentNode.Children[string(i)]

			}
		}
	}
	return errors.New("person already exist")
}
func (trie *Trie) Search(name string) (string, error) {
	currentNode := trie.Root
	for pos, i := range name {
		if _, ok := currentNode.Children[string(i)]; ok {
			if pos == len(name)-1 {
				currentNode = currentNode.Children[string(i)]
				return currentNode.Address, nil
			} else {
				currentNode = currentNode.Children[string(i)]
			}

		} else {
			return "", errors.New("name doesn't exist")
		}

	}
	return "", errors.New("unexpected error happened")
}
