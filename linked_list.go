package linked_list

type User struct {
	Id int
	Name string
	Age int
}

type Node struct {
	Data User
	Prev *Node
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *List) AddFirst(node *Node) {
	if l.Head == nil {
		l.Tail = node
	} else {
		l.Head.Prev = node
	}
	node.Next = l.Head
	l.Head = node
	l.Size++
}

