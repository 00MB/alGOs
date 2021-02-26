package main

import (
	"fmt"
)


type Node struct {
	value	int
	children []*Node
	root bool
}

func minimax(tree Node, depth int, maximizingPlayer bool) (value int){
	if depth == 0 || tree.root {
		value := tree.value
		return value
	}
	if maximizingPlayer {
		value := -1000000
		for _, child := range tree.children {
			temp := minimax(*child, depth - 1, false)
			if temp > value {
				value = temp 
			}
		} 
		return value
	} else {
		value := 1000000
		for _, child := range tree.children {
			temp := minimax(*child, depth -1, true)
			if temp < value {
				value = temp
			}
		} 
		return value
	}
}

func main() {
	treelow11 := Node{
		value: 2,
		children: []*Node{},
		root: true,
	}
	treehigh11 := Node{
		value: 5,
		children: []*Node{},
		root: true,
	}
	treelow12 := Node{
		value: 6,
		children: []*Node{},
		root: true,
	}
	treehigh12 := Node{
		value: 8,
		children: []*Node{},
		root: true,
	}
	treelow1 := Node{
		value: 3,
		children: []*Node{&treelow11, &treehigh11},
		root: false,
	}
	treehigh1 := Node{
		value: 7,
		children: []*Node{&treelow12, &treehigh12},
		root: false,
	}
	tree := Node{
		value: 5,
		children: []*Node{&treelow1, &treehigh1},
		root: false,
	}
	value := minimax(tree, 2, true)
	fmt.Println("value returned:",value)
}