package graph

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostu/app/services/graphServices"
	"gostu/pkg/response"
)

func Graph(ctx *gin.Context) {
	graph := &graphServices.Graph{}
	graph.Graph(8)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 7)
	graph.ShowGraph()
	graph.BreadthFirstSearch(0, 6)
	fmt.Println()
	graph.DepthFirstSearch(0, 6)

	response.SuccessResponse(ctx, "succ", graph)
}
