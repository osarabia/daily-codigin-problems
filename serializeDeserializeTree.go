package main

import (
	"fmt"
	"strings"
)

/*
Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
*/

type Node struct{
	val string
	left *Node
	right *Node
}

func main() {
	leftleft := Node{"left.left", nil, nil}
	left := Node{"left", &leftleft, nil}
	right := Node{"right", nil, nil}
	root := Node{"root", &left, &right}
	//fmt.Printf("%v\n", root)

	btSerialized := make([]string, 0, 5)
	ss := serialize(root, btSerialized)
	strBTrepre := strings.Join(ss, ",")
	fmt.Println(strBTrepre)

	rootNode := deserialize(strBTrepre)
	fmt.Printf("%v\n", rootNode.left)
        fmt.Printf("%v\n", rootNode.left.left)
	fmt.Printf("%v\n", rootNode.right)
}

func serialize(root Node, output []string) []string {

	output = append(output, root.val)
        if root.left  != nil {
		output = serialize(*root.left, output)
	} else {
		output = append(output, "#")
	}

	if root.right != nil {
		output = serialize(*root.right, output)
	} else {
		output = append(output, "#")
	}

	return output
}

func helperDeserialize(data []string, k *int) *Node{
	val := data[*k]
	if val == "#" {
		n := Node{"#", nil, nil}
		return &n
	}

	//parent Node
	p := Node{val, nil, nil}
	*k += 1
	p.left = helperDeserialize(data, k)
	*k += 1
	p.right = helperDeserialize(data, k)

	return &p
}

func deserialize(input string) Node{
	var k int
	data := strings.Split(input, ",")
	//fmt.Printf("%v\n", data)

	val := data[0]
	if val  == "#" {
	    return Node{"#", nil, nil}
        }
	//root Node
	r := Node{val, nil, nil}
	k += 1
	r.left = helperDeserialize(data, &k)
	k += 1
	r.right = helperDeserialize(data, &k)

	return r
}
