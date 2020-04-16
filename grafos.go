package main

import (
	"fmt"

	"github.com/yourbasic/graph"
)

func main() {

	var vertexCount int
	fmt.Print("Introduzca la cantidad de 'Nodos' o ' Vertices' que tendra el Grafo: ")
	fmt.Scanf("%d", &vertexCount)
	fmt.Println(vertexCount)
	g := graph.New(vertexCount)
	g.AddBoth(0, 1) //  0 -- 1
	g.AddBoth(0, 2) //  |    |
	g.AddBoth(2, 3) //  2 -- 3
	g.AddBoth(1, 3)
	fmt.Println(g.String())
}
