package graph

// 图
type Graph struct {
	nodes []*Graph
	value string
}

type Queue struct {
	head *node
	tail *node
}

type node struct {
	value string
	next  *node
}

func BFS(g Graph, name string) string {
	q := &Queue{nil, nil}

}
