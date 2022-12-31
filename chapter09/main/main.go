package main

import (
	"fmt"
	"math/rand"
	"time"

	bst "chapter09/binarysearchtree"
)

// Satisfies OrderedStringer because of ~int
// Also satisfies OrderedStringer because of String() method below
type Number int

func (num Number) String() string {
	return fmt.Sprintf("%d", num)
}

type Float float64

func (num Float) String() string {
	return fmt.Sprintf("%0.1f", num)
}

func inorderOperator(val Float) {
	fmt.Println(val.String())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Generate a random search tree
	randomSearchTree := bst.BinarySearchTree[Float]{nil, 0}
	for i := 0; i < 30; i++ {
		rn := 1.0 + 99.0*rand.Float64()
		randomSearchTree.Insert(Float(rn))
	}
	time.Sleep(3 * time.Second)
	bst.ShowTreeGraph(randomSearchTree)
	randomSearchTree.InOrderTraverse(inorderOperator)
	min := randomSearchTree.Min()
	max, _ := randomSearchTree.Max()
	fmt.Printf("\nMinimum value in random search tree is %0.1f  \nMaximum value in random search tree is %0.1f", *min, *max)

	start := time.Now()
	tree := bst.BinarySearchTree[Number]{nil, 0}
	for val := 0; val < 100_000; val++ {
		tree.Insert(Number(val))
	}
	elapsed := time.Since(start)
	_, ht := tree.Max()
	fmt.Printf("\nTime to build BST tree with 100,000 nodes in sequential order: %s. Height of tree: %d", elapsed, ht)
}

/*
3.5
3.6
8.6
9.0
13.2
25.9
27.3
31.4
41.7
46.4
47.6
49.4
57.0
58.4
60.0
60.4
63.3
64.3
66.0
67.8
69.9
74.2
77.5
78.1
89.3
93.0
93.4
95.1
96.9
99.4

Minimum value in random search tree is 3.5
Maximum value in random search tree is 99.4
Time to build BST tree with 100,000 nodes in sequential order: 21.330900852s. Height of tree: 100000
*/
