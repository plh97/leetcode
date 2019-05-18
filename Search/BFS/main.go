package BFS

// algorithm of graoph
// BSF

var graph = map[string][]string{
	"A": []string{"B", "C"},
	"B": []string{"A", "C", "D"},
	"C": []string{"A", "B", "D", "E"},
	"D": []string{"B", "C", "E", "F"},
	"E": []string{"C", "D"},
	"F": []string{"D"},
}

func BSF(graph map[string][]string, s string) []string {
	res := []string{}
	queue := []string{}
	queue = append([]string{s}, queue...) // 队列压入第一个
	seen := make(map[string]string)       // 可去重的集合映射关系
	seen[s] = s
	for len(queue) > 0 {
		vertex := queue[len(queue)-1] // 队列推出最后一个, 为什么在可选择下一个点的队列中推出最后一个呢, 这就是广度优先了
		queue = queue[:len(queue)-1]
		nodes := graph[vertex]
		for i := range nodes {
			node := nodes[i]
			_, ok := seen[node]
			if !ok {
				queue = append([]string{node}, queue...) // 插入第一个
				seen[node] = node
			}
		}
		res = append(res, vertex)
	}
	return res
}
