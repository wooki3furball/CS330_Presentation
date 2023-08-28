// Author: Michael Bopp
// Filename: btree.go
// Description: Go program demoing a binary tree, go routines, & channels.
// Date Created: 8/21/23
// Last Edited: 8/28/23

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}
	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else {
		root.Right = insert(root.Right, value)
	}
	return root
}

func inorderTraversal(root *Node, ch chan int) {
	if root != nil {
		inorderTraversal(root.Left, ch)
		ch <- root.Value
		inorderTraversal(root.Right, ch)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func runBinaryTree(values []int, treeName string) {
	startTime := time.Now()
	var root *Node
	for _, value := range values { // Traversing range of container
		root = insert(root, value)
	}
	insertTime := time.Since(startTime)

	ch := make(chan int, len(values))

	go func() {
		inorderTraversal(root, ch)
		close(ch)
	}()

	fmt.Printf("\n%s tree (sorted):\n", treeName)
	traversalStartTime := time.Now()
	for value := range ch { // Traversing range of channel type
		fmt.Println(value)
	}
	traversalTime := time.Since(traversalStartTime)

	fmt.Printf("%s tree - Insertion time: %v, Traversal time: %v\n", treeName, insertTime, traversalTime)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter the number of values to insert")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// Generate Fibonacci sequence and place it randomly in an array
	rand.Seed(time.Now().UnixNano())
	fibArray := make([]int, n)
	for i := 0; i < n; i++ {
		fibArray[i] = fibonacci(i)
	}
	rand.Shuffle(len(fibArray), func(i, j int) { fibArray[i], fibArray[j] = fibArray[j], fibArray[i] })

	// Generate a random array of integers
	randArray := make([]int, n)
	for i := 0; i < n; i++ {
		randArray[i] = rand.Intn(1000)
	}

	fmt.Println("Random Fibonacci sequence:")
	fmt.Println(fibArray)
	runBinaryTree(fibArray, "Fibonacci")

	fmt.Println("\nRandom integer sequence:")
	fmt.Println(randArray)
	runBinaryTree(randArray, "Random Integer")
}
