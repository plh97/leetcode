# algorithm of graoph
# BSF
import heapq
import math
graph = {
    "A": {"B": 5, "C": 1},
    "B": {"A": 5, "C": 2, "D": 1},
    "C": {"A": 1, "B": 2, "D": 4, "E": 8},
    "D": {"B": 1, "C": 4, "E": 3, "F": 6},
    "E": {"C": 8, "D": 3},
    "F": {"D": 8},
}


def init_distance(graph, s):
    distance = {s: 0}
    for vertex in graph:
        if vertex != s:
            distance[vertex] = math.inf
    return distance


def dijkstra(graph, s):
    pqueue = []
    heapq.heappush(pqueue, {0, s})
    seen = set()
    parent = {s: None}
    distance = init_distance(graph, s)

    while (len(pqueue) > 0):
        pair = heapq.heappop(pqueue)
        dist = pair[0]
        vertex = pair[1]
        seen.add(vertex)

        nodes = graph[vertex].keys()
        for w in nodes:
            if w not in seen:
                if dist + graph[vertex][w] < distance[w]:
                    heapq.heappush(pqueue, (dist+graph[vertex][w], w))
                    # pqueue.append(w)
                    parent[w] = vertex
                    distance[w] = dist + graph[vertex][w]
                    # seen.add(w)
        print(vertex)
    return parent, distance


parent, distance = dijkstra(graph, "A")  # A->B->C->D->E->F

print(parent)
print(distance)
