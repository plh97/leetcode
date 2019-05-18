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

// 广度优先
function DSF(graph, s) {
  stack = []
  stack.push(s)
  seen = new Set()
  seen.add(s)
  while (stack.length > 0) {
    vertex = stack.pop()
    nodes = graph[vertex]
    for (let w of nodes) {
      if (!seen.has(w)) {
        stack.push(w)
        seen.add(w)
      }
    }
    console.log(vertex)
  }
}

// A->C->E->D->F->B
DSF(graph, "A")