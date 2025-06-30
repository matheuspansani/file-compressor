package huffman

type Node struct {
	freq int
	char byte
	left *Node
	right *Node
}