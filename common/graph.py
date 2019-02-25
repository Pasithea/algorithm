"""
    Graph bfs and dfs implementation.
"""

from collections import deque
from typing import List


class Graph:
    def __init__(self, vertex_num: int=10):
        self._adj = [[] for _ in range(vertex_num)]
        self._vertex_num = vertex_num

    def add_edge(self, s: int, t: int):
        self._adj[s].append(t)
        self._adj[t].append(s)

    def bfs(self, s: int, t: int) -> None:
        if s == t:
            return
        visited = [False for _ in range(self._vertex_num)]
        queue = deque()
        prev = [None for _ in range(self._vertex_num)]
        queue.append(s)
        while len(queue) != 0:
            w = queue.popleft()
            for i in range(len(self._adj[w])):
                v = self._adj[w][i]
                if not visited[v]:
                    prev[v] = w
                    if v == t:
                        self._print_way(prev, s, t)
                        return
                    visited[v] = True
                    queue.append(v)

    def dfs(self, s: int, t: int) -> None:
        found = False
        visited = [False for _ in range(self._vertex_num)]
        prev = [None for _ in range(self._vertex_num)]

        def _dfs(w: int):
            nonlocal found
            if found:
                return
            visited[w] = True
            if w == t:
                found = True
                return
            for i in range(len(self._adj[w])):
                v = self._adj[w][i]
                if not visited[v]:
                    prev[v] = w
                    _dfs(v)
        _dfs(s)
        self._print_way(prev, s, t)


    def _print_way(self, prev: List, s: int, t: int):
        if prev[t] is not None and t != s:
            self._print_way(prev, s, prev[t])
        print(str(t))


if __name__ == '__main__':
    g = Graph()
    g.add_edge(1, 2)
    g.add_edge(1, 4)
    g.add_edge(2, 3)
    g.add_edge(2, 5)
    g.add_edge(3, 6)
    g.add_edge(4, 5)
    g.add_edge(5, 7)
    g.add_edge(6, 8)
    print('bfs:')
    g.bfs(1, 8)
    print('dfs:')
    g.dfs(1, 8)
