package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Value int       `json:"value"`
	Left  *TreeNode `json:"left,omitempty"`
	Right *TreeNode `json:"right,omitempty"`
}

const (
	INVALID_RESPONSE = "The binary tree is not a valid BST"
	VALID_RESPONSE   = "The binary tree is a valid BST"
)

func main() {

	vf := flag.String("validate", "", "Validate a JSON string representing a binary tree")
	gf := flag.String("generate", "", "Generate a JSON string representing the binary tree with a given height")

	flag.Parse()

	if *vf != "" {
		treeJson := *vf // pointer to JSON string

		// Convert the JSON string to a binary tree
		tree := jsonToBinarySearchTree(treeJson)

		// Print the binary tree in ASCII format
		printTree(&tree, "", false)

		// Check if the binary tree is a valid BST
		if !isValidBinarySearchTree(&tree) {
			log.Fatal(INVALID_RESPONSE)
		}

		// Log the result
		log.Print(VALID_RESPONSE)

		os.Exit(0)
	}

	if *gf != "" {
		// Generate a binary tree of a given height and return its JSON representation
		height, err := strconv.Atoi(*gf)
		if err != nil {
			log.Fatalf("Failed to parse height: %s", err)
		}
		log.Print(generateBSTAsJSON(height))
		os.Exit(0)
	}
}

// isValidBinarySearchTreeUtil checks if the binary tree is a valid BST
func isValidBinarySearchTreeUtil(node *TreeNode, min, max int) bool {

	// If the node is nil, it is a valid BST
	if node == nil {
		return true
	}

	// If the node value is not within the min and max, it is not a valid BST
	if node.Value <= min || node.Value >= max {
		return false
	}

	// Check the left and right nodes recursively
	return isValidBinarySearchTreeUtil(node.Left, min, node.Value) && isValidBinarySearchTreeUtil(node.Right, node.Value, max)
}

// isValidBinarySearchTree checks if the binary tree is a valid BST
// It is separate from the util function to allow first call without min and max
func isValidBinarySearchTree(root *TreeNode) bool {
	return isValidBinarySearchTreeUtil(root, math.MinInt64, math.MaxInt64)
}

// generateBinarySearchTreeUtil creates a binary search tree of a given height
func generateBinarySearchTreeUtil(currentHeight, targetHeight, startValue int) *TreeNode {
	if currentHeight > targetHeight {
		return nil
	}

	valueStep := int(math.Pow(2, float64(targetHeight-currentHeight)))

	node := &TreeNode{
		Value: startValue,
		Left:  generateBinarySearchTreeUtil(currentHeight+1, targetHeight, startValue-valueStep),
		Right: generateBinarySearchTreeUtil(currentHeight+1, targetHeight, startValue+valueStep),
	}

	return node
}

// generateBSTAsJSON generates a binary search tree of a given height and returns its JSON representation.
// It is separate from the util function to allow first call without min and max
func generateBSTAsJSON(height int) string {
	root := generateBinarySearchTreeUtil(1, height, int(math.Pow(2, float64(height-1)))) // Starting value can be adjusted

	b, err := json.Marshal(root)
	if err != nil {
		log.Fatalf("Failed to generate JSON: %s", err)
	}
	return string(b)
}

// jsonToBinarySearchTree converts a JSON string to a binary tree
func jsonToBinarySearchTree(j string) TreeNode {
	var root TreeNode
	err := json.Unmarshal([]byte(j), &root)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}
	return root
}

// printTree prints the binary tree in ASCII format using / and \
func printTree(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	var linePrefix, valuePrefix string
	if node != nil {
		linePrefix = prefix + "    "
		valuePrefix = prefix
		if isLeft {
			valuePrefix += "└── "
		} else {
			valuePrefix += "┌── "
		}
		printTree(node.Right, linePrefix, false)
		fmt.Println(valuePrefix + strconv.Itoa(node.Value))
		printTree(node.Left, linePrefix, true)
	}
}
