package rbtree

const (
	// RED color
	RED = true
	// BLACK color with initial value since a nil node is black
	BLACK = false
)

type (
	// KeyType is the type with CompareTo behavior.
	KeyType interface {
		CompareTo(c interface{}) int
	}

	// KeyTypeInt is the int that implements the KeyType interface.
	KeyTypeInt int

	// Node is the red-black tree node
	Node struct {
		Key    KeyType
		Val    interface{}
		Color  bool
		Left   *Node
		Right  *Node
		Parent *Node
	}

	// RBT is the red-black tree
	RBT struct {
		root *Node
		size uint64
	}
)

// CompareTo implementation for type KeyTypeInt
func (k KeyTypeInt) CompareTo(c interface{}) int {
	return int(k) - int(c.(KeyTypeInt))
}

// isRed returns whether a node is in red color, nil is black
func isRed(n *Node) bool {
	if n == nil {
		return BLACK
	}
	return n.Color
}

// NewRBT returns a red-black tree
func NewRBT() *RBT {
	return &RBT{}
}

// Size returns the number of nodes that the red-black tree hold
func (t *RBT) Size() uint64 {
	return t.size
}

// min find the minimum node in n's sub-tree
func min(n *Node) *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

// max find the maximum node in n's sub-tree
func max(n *Node) *Node {
	for n.Right != nil {
		n = n.Right
	}
	return n
}

// predecessor returns n's predecessor node
func (t *RBT) predecessor(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.Left != nil {
		return max(n.Left)
	}

	p := n.Parent
	lc := n
	for p != nil && lc == p.Left {
		lc = p
		p = p.Parent
	}
	return p

}

// unused in this code
// successor returns n's successor node
func (t *RBT) successor(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.Right != nil {
		return min(n.Right)
	}

	p := n.Parent
	lc := n
	for p != nil && lc == p.Right {
		lc = p
		p = p.Parent
	}
	return p
}

// LeftRotate left rotate the node n, acting on a red link
//
//    p               p
//    |               |
//    n               m
//   / \             / \
//  nl  m    ==>    n  mr
//     / \         / \
//    ml mr       nl ml
//
func (t *RBT) leftRotate(n *Node) {
	m := n.Right
	n.Right = m.Left
	if m.Left != nil {
		m.Left.Parent = n
	}
	m.Parent = n.Parent
	if n.Parent == nil {
		t.root = m
	} else if n == n.Parent.Left {
		n.Parent.Left = m
	} else {
		n.Parent.Right = m
	}
	m.Left = n
	n.Parent = m
	// repaint color
	m.Color = n.Color
	n.Color = RED
}

// RightRotate right rotate the node n, acting on a red link
//
//     p                p
//     |                |
//     n                m
//    / \              / \
//   m  nr    ==>     ml  n
//  / \                  / \
// ml mr                mr nr
//
func (t *RBT) rightRotate(n *Node) {
	m := n.Left
	n.Left = m.Right
	if m.Right != nil {
		m.Right.Parent = n
	}
	m.Parent = n.Parent
	if n.Parent == nil {
		t.root = m
	} else if n == n.Parent.Left {
		n.Parent.Left = m
	} else {
		n.Parent.Right = m
	}

	m.Right = n
	n.Parent = m

	m.Color = n.Color
	n.Color = RED
}

func (t *RBT) flipColors(n *Node) {
	n.Color = RED
	n.Left.Color = BLACK
	n.Right.Color = BLACK
}

// Search returns the node by the given key if it exists
func (t *RBT) Search(key KeyType) interface{} {
	if n := t.search(t.root, key); n != nil {
		return n.Val
	}
	return nil
}

func (t *RBT) search(r *Node, key KeyType) *Node {
	for r != nil {
		cmp := key.CompareTo(r.Key)
		if cmp < 0 {
			r = r.Left
		} else if cmp > 0 {
			r = r.Right
		} else {
			break
		}
	}
	return r
}

// Insert inserts a key with associated data(val) to a place in the red-black
// tree by compare the keys, and if the key exists, update it with the new val.
func (t *RBT) Insert(key KeyType, val interface{}) {
	if t.root == nil {
		t.root = &Node{Key: key, Val: val}
		t.size = 1
		return
	}

	// find parent node to attach
	root := t.root
	var p *Node
	for root != nil {
		p = root
		cmp := key.CompareTo(root.Key)
		if cmp < 0 {
			root = root.Left
		} else if cmp > 0 {
			root = root.Right
		} else {
			root.Val = val
			return
		}
	}

	newnode := &Node{Key: key,
		Val:    val,
		Color:  RED,
		Parent: p,
	}

	cmp := key.CompareTo(p.Key)
	if cmp < 0 {
		p.Left = newnode
	} else {
		p.Right = newnode
	}

	t.insertFix(newnode)
	t.size++
}

func (t *RBT) insertFix(n *Node) {
	var u *Node // n's uncle node
	for n.Parent != nil && isRed(n.Parent) {
		if n.Parent == n.Parent.Parent.Left {
			u = n.Parent.Parent.Right
			if u != nil && isRed(u) {
				n = n.Parent.Parent
				t.flipColors(n)
			} else {
				if n == n.Parent.Right {
					n = n.Parent
					t.leftRotate(n)
				}
				t.rightRotate(n.Parent.Parent)
			}
		} else {
			u = n.Parent.Parent.Left
			if u != nil && isRed(u) {
				n = n.Parent.Parent
				t.flipColors(n)
			} else {
				if n == n.Parent.Left {
					n = n.Parent
					t.rightRotate(n)
				}
				t.leftRotate(n.Parent.Parent)
			}
		}
	}
	// root should be black node at the end
	t.root.Color = BLACK
}

// Remove delete a node by the given key if it exits
func (t *RBT) Remove(key KeyType) {
	d := t.search(t.root, key)

	if d == nil {
		return // not found
	}

	// find replacement node
	//   in our code, if finally the rpl hold a child node, the rpl must be a
	// predecessor or successor
	rpl := d
	if rpl.Left != nil && rpl.Right != nil {
		// rpl = rpl.successor()
		rpl = t.predecessor(rpl)
	} else {
		if rpl.Left != nil {
			rpl = rpl.Left
		} else if rpl.Right != nil {
			rpl = rpl.Right
		}
	}
	// delete replacement node
	if rpl != t.root {
		if d != rpl { // d is not a leaf node
			d.Key = rpl.Key
			d.Val = rpl.Val
		}

		if rpl.Left != nil { // rpl is a predecessor
			if !isRed(rpl) {
				rpl.Left.Color = BLACK
			}

			rpl.Left.Parent = rpl.Parent
			if rpl == rpl.Parent.Left {
				rpl.Parent.Left = rpl.Left
			} else {
				rpl.Parent.Right = rpl.Left
			}
			// unlink rpl
			rpl.Parent = nil
			rpl.Left = nil
		} else {
			// rpl is a leaf node
			// fix
			if !isRed(rpl) {
				t.removeFix(rpl)
			}
			// then delete
			if rpl == rpl.Parent.Left {
				rpl.Parent.Left = nil
			} else {
				rpl.Parent.Right = nil
			}
			// unlink rpl
			rpl.Parent = nil
		}
	} else { // single node tree
		t.root = nil
	}
	t.size--
}

// fixAfterRemove do fix if the deleted node is in black color
func (t *RBT) removeFix(n *Node) {
	for n != t.root && !isRed(n) {
		if n == n.Parent.Left {
			rBro := n.Parent.Right
			// find real brother node
			if isRed(rBro) {
				t.leftRotate(n.Parent)
				rBro = n.Parent.Right
			}
			if !isRed(rBro.Left) && !isRed(rBro.Right) { // can't borrow
				rBro.Color = RED
				n = n.Parent
			} else { // can borrow
				// 3-node
				if !isRed(rBro.Right) {
					t.rightRotate(rBro)
					rBro = n.Parent.Right
				}
				t.leftRotate(n.Parent)
				// since we pull down n's parent to replace n and
				// we borrow two n's bro's childern, we need repaint
				n.Parent.Color = BLACK
				rBro.Right.Color = BLACK
				n = t.root
			}
		} else {
			lBro := n.Parent.Left

			if isRed(lBro) {
				t.rightRotate(n.Parent)
				lBro = n.Parent.Left
			}
			if !isRed(lBro.Left) && !isRed(lBro.Right) {
				lBro.Color = RED
				n = n.Parent
			} else {
				if !isRed(lBro.Left) {
					t.leftRotate(lBro)
					lBro = n.Parent.Left
				}
				t.rightRotate(n.Parent)
				n.Parent.Color = BLACK
				lBro.Left.Color = BLACK
				n = t.root
			}
		}
	}
	n.Color = BLACK
}
