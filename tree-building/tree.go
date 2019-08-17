package tree

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields ID and Parent
// and Node is a struct containing int field ID and []*Node field Children.

// Record contains the ID of a record and its parent id.
type Record struct {
	ID     int
	Parent int
}

// Node represents a forum post in a tree structure. The node has an ID and child nodes.
type Node struct {
	ID       int
	Children *Node
}

func Build(records []Record) (*Node, error) {

}
