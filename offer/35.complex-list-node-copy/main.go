package _35_complex_list_node_copy

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cacheMap map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if cache, ok := cacheMap[node]; ok {
		return cache
	}
	newNode := &Node{Val: node.Val}
	cacheMap[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	cacheMap = map[*Node]*Node{}
	return deepCopy(head)
}
