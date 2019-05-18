package DFS

// algorithm of graoph
// DFS -> deep first search

var graph = map[string][]string{
	"A": []string{"B", "C"},
	"B": []string{"A", "C", "D"},
	"C": []string{"A", "B", "D", "E"},
	"D": []string{"B", "C", "E", "F"},
	"E": []string{"C", "D"},
	"F": []string{"D"},
}

func DFS(graph map[string][]string, s string) []string {
	res := []string{}
	stack := []string{}
	stack = append([]string{s}, stack...) // 队列压入最后一个
	seen := make(map[string]string)       // 可去重的集合映射关系
	seen[s] = s
	for len(stack) > 0 {
		vertex := stack[0] // 队列推出最后一个
		stack = stack[1:]
		nodes := graph[vertex]
		for i := range nodes {
			node := nodes[i]
			_, ok := seen[node]
			if !ok {
				stack = append([]string{node}, stack...) // 队列压入最后一个
				seen[node] = node
			}
		}
		res = append(res, vertex)
	}
	return res
}
