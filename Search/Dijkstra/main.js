// # algorithm of graoph
// # BSF 

graph = {
  "A": ["B", "C"],
  "B": ["A", "C", "D"],
  "C": ["A", "B", "D", "E"],
  "D": ["B", "C", "E", "F"],
  "E": ["C", "D"],
  "F": ["D"],
}

function BSF(graph, s) {
  // queue 队列  [a, b, c, d]  移入 [new, a, b, c] , 移除 pop
  queue = []
  queue.unshift(s)
  seen = new Set()
  seen.add(s)       // 用于记录曾经访问过的节点, 不再访问
  while (queue.length > 0) {
    vertex = queue.pop()
    nodes = graph[vertex]
    for (let w of nodes) {
      if (!seen.has(w)) {
        queue.unshift(w)
        seen.add(w)
      }
    }
    console.log(vertex)
  }
}


// A->B->C->D->E->F
BSF(graph, "A")