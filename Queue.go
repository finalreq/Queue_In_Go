// +build

package Queue

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Queue struct {
	Head *Node
}

func (q *Queue) Enqueue(value int) {
	var n *Node = new(Node)
	n.Data = value

	if q.Head == nil {
		n.Next = nil
		q.Head = n
		return
	}

	n.Next = q.Head
	q.Head = n
	return
}

func (q *Queue) Dequeue() *Node {
	current := q.Head
	previous := q.Head

	for current.Next != nil {
		previous = current
		current = current.Next
	}
	previous.Next = nil
	return current
}

func (q *Queue) GetSize() int {
	n := q.Head
	i := 1

	if n == nil { return 0 }

	for n.Next != nil {
		i++
		n = n.Next
	}
	return i
}

func (q *Queue) IsEmpty() bool {
	if q.GetSize() == 0 {
		return true
	}
	return false
}


func (q *Queue) Print() {
	n := q.Head

	for n.Next != nil {
		fmt.Printf("%d -> ", n.Data)
		n = n.Next
	}
	fmt.Printf("%d\n", n.Data)
}
