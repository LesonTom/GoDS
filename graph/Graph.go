package graph

import "fmt"

type AdjMatrix struct {
	vertices  int
	adjMatrix [][]int
}

type AdjTable struct {
	vertices int
	adjList  map[int][]int
}

func NewGraphT(vertices int) *AdjTable {
	return &AdjTable{vertices, make(map[int][]int)}
}

func (g *AdjTable) AddEdge(from, to int) {
	g.adjList[from] = append(g.adjList[from], to)
	g.adjList[to] = append(g.adjList[to], from) // 无向图需要这一行
}

func (g *AdjTable) Print() {
	for vertex := 0; vertex < g.vertices; vertex++ {
		fmt.Printf("顶点 %d 的邻居：", vertex)
		for _, neighbor := range g.adjList[vertex] {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

func NewGraphM(vertices int) *AdjMatrix {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}
	return &AdjMatrix{vertices, matrix}
}
func (g *AdjMatrix) AddEdge(from, to int) {
	g.adjMatrix[from][to] = 1
	g.adjMatrix[to][from] = 1 // 无向图需要这一行
}

func (g *AdjMatrix) Print() {
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			fmt.Printf("%d ", g.adjMatrix[i][j])
		}
		fmt.Println()
	}
}
