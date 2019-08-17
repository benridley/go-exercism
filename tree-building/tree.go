package tree

import (
	"errors"
	"fmt"
	"sort"
)

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
	Children []*Node
}

// Build creates a tree structure from a slice of records using their ID and their parent IDs.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	// Sort records based on ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]*Node, len(records))

	// Create the root node, a special case with no parents.
	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("No valid root node. Check root node exists and has no parent nodes")
	}
	nodes[0] = &Node{
		ID:       0,
		Children: nil,
	}

	// Step through the slice, creating nodes as required. Because the slice is sorted, each record ID should be its index in the slice.
	for i := 1; i < len(records); i++ {
		if records[i].ID != i {
			return nil, fmt.Errorf("Missing record at ID: %d, records are non continuous", records[i].ID)
		}
		if records[i].Parent >= i && i != 0 {
			return nil, fmt.Errorf("Record has parent ID of itself or greater at ID: %d", records[i].ID)
		}

		newNode := Node{
			ID:       i,
			Children: nil,
		}
		nodes[i] = &newNode

		// Lookup this nodes parent and add this node as one of its children.
		nodes[records[i].Parent].Children = append(nodes[records[i].Parent].Children, &newNode)
	}

	return nodes[0], nil
}
