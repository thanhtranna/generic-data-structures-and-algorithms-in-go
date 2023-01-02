package main

import (
	"fmt"
	"math/rand"
	"time"

	avl "chapter10/genericavl/avl"
)

func inorderOperator(val Float) {
	fmt.Println(val.String())
}

// Float Satisfies OrderedStringer because of ~float64
// Also satisfies OrderedStringer because of String() method below
type Float float64

func (num Float) String() string {
	return fmt.Sprintf("%0.1f", num)
}

type Integer int

func (num Integer) String() string {
	return fmt.Sprintf("%d", num)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var data [100_000]int
	for i := 0; i < 100_000; i++ {
		data[i] = rand.Intn(1_000_000)
	}

	// Generate a random search tree
	randomSearchTree := avl.AVLTree[Float]{nil, 0}
	for i := 0; i < 30; i++ {
		rn := 1.0 + 99.0*rand.Float64()
		randomSearchTree.Insert(Float(rn))
	}
	time.Sleep(3 * time.Second)
	avl.ShowTreeGraph(randomSearchTree)

	randomSearchTree.InOrderTraverse(inorderOperator)
	min := randomSearchTree.Min()
	max := randomSearchTree.Max()
	fmt.Printf("\nMinimum value in tree is %0.1f  Maximum value in tree is %0.1f", *min, *max)

	start := time.Now()
	tree := avl.AVLTree[Integer]{nil, 0}
	for i := 0; i < 100_000; i++ {
		tree.Insert(Integer(data[i]))
	}
	elapsed := time.Since(start)
	fmt.Printf("\nInsertion time for AVL tree: %s.  Height of tree: %d", elapsed, tree.Height())

	start = time.Now()
	for i := 0; i < 100_000; i++ {
		_ = tree.Search(Integer(i))
	}
	elapsed = time.Since(start)
	fmt.Println("\nSearch time for AVL tree: ", elapsed)
}

/* Output
1.9
7.7
9.3
15.0
20.1
26.1
26.3
30.4
32.5
33.1
36.4
37.6
37.7
38.6
48.1
48.5
49.0
50.6
51.4
52.7
66.3
72.7
73.8
75.0
76.0
87.1
92.9
95.3
98.8
99.0

Minimum value in tree is 1.9  Maximum value in tree is 99.0
Insertion time for AVL tree: 48.596183ms.  Height of tree: 20
Search time for AVL tree:  4.113031ms
*/
