package containers

type Queue struct {
	MaxSize int
	count   int
	head    *Node
	tail    *Node
}

type Node struct {
	value string
	next  *Node
}

func (q *Queue) add(data string) (err error) {
}
