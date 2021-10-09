package graphServices

import (
	"fmt"
	"gostu/app/services/linkedListServices"
)

type Graph struct {
	vertex int
	abj []*linkedListServices.List
}

func (g *Graph) Graph(v int) {
	g.vertex = v
	g.abj = make([]*linkedListServices.List, v)
	for i := 0; i < v; i++ {
		g.abj[i] = &linkedListServices.List{}
	}
}

func (g *Graph) AddEdge(s, t int)  {
	g.abj[s].Append(t)
	g.abj[t].Append(s)
}

func (g *Graph) BreadthFirstSearch(s, t int) {
	if s == t {return}
	visited := make([]bool, g.vertex)
	visited[s] = true
	var queue []int
	queue = append(queue, s)
	prev := make([]int, g.vertex)
	for i := 0; i < g.vertex; i++ {
		prev[i] = -1
	}

	for len(queue) != 0 {
		w := queue[0]
		queue = queue[1:]
		for i := 0; i < g.abj[w].Length(); i++ {
			q, _ := g.abj[w].Get(i).(int)
			if !visited[q] {
				prev[q] = w
				if q == t {
					handlerFunc(prev, s, t)
					return
				}
				visited[q] = true
				queue = append(queue, q)
			}
		}
	}
}

var found bool = false

func (g *Graph) DepthFirstSearch(s, t int) {
	found = false
	visited := make([]bool, g.vertex)
	prev := make([]int, g.vertex)
	for i := 0; i < g.vertex; i++ {
		prev[i] = -1
	}
	g.recurDepthFirstSearch(s, t, visited, prev)
	handlerFunc(prev, s, t)
}

func (g *Graph) recurDepthFirstSearch(w, t int, visited []bool, prev []int)  {
	if found == true {return}
	visited[w] = true
	if w == t {
		found = true
		return
	}
	for i := 0; i < g.abj[w].Length(); i++ {
		q, _ := g.abj[w].Get(i).(int)
		if !visited[q] {
			prev[q] = w
			g.recurDepthFirstSearch(q, t, visited, prev)
		}
	}
}

func handlerFunc(prev []int, s, t int)  {
	if prev[t] != -1 && t != s {
		handlerFunc(prev, s, prev[t])
	}
	fmt.Println(t, " ")
}

func (g *Graph) ShowGraph()  {
	fmt.Println("图：邻接表")
	fmt.Printf("顶点个数：%v\n", g.vertex)
	fmt.Println("边：")
	for i, val := range g.abj {
		fmt.Printf("第%d顶点\n", i)
		val.ShowList()
		fmt.Println()
	}
}


