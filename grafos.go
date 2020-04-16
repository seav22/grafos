package main

import (
	"fmt"
	//importando paquete de grafos de github
	"github.com/yourbasic/graph"
)

func main() {

	var vertexCount int //almacena la cantidad de nodos que tendra el grafo
	var graphType int
	var addEdgeCheck int //bandera para añadir aristas
	var edgeCost int     //para definir si la arista tendra peso o no
	var vertex [2]int    //de donde a donde va la arista sin peso
	var cost int64

	//tamaño del grafo
	fmt.Print("Introduzca la cantidad de 'Nodos' o ' Vertices' que tendra el Grafo: ")
	fmt.Scanf("%d", &vertexCount)
	//construyo el grafo con el tamaño solicitado antes
	g := graph.New(vertexCount)
	//pregunto que tipo de grafo sera
	fmt.Println("El grafo será undirected o directed? (0 para undirected, 1 para directed)")
	fmt.Scanf("%d", &graphType)

	//grafo no dirigido
	if graphType == 0 {
		fmt.Println("Desea añadir aristas al nodo? (0 para no 1 para si)")
		fmt.Scanf("%d", &addEdgeCheck)
		//si o no añadir aristas al nodo
		if addEdgeCheck == 1 {
			fmt.Print("La arista tendra peso? (0 para no, 1 para si): ")
			fmt.Scanf("%d", &edgeCost)
			if edgeCost == 0 {
				//origen y fin de arista
				fmt.Println("Indique de que vertice a que vertice se creara la arista, no importa el orden")
				//vertice 1
				fmt.Print("Vertice 1: ")
				fmt.Scanf("%d", vertex[0])
				//vertice 2
				fmt.Print("Vertice 2: ")
				fmt.Scanf("%d", vertex[1])
				//vertex[2] = nil
				//creando arista
				g.AddBoth(vertex[0], vertex[1])
			} else {
				//origen y fin de arista
				fmt.Println("Indique de que vertice a que vertice se creara la arista, no importa el orden")
				//vertice 1
				fmt.Print("Vertice 1: ")
				fmt.Scanf("%d", vertex[0])
				//vertice 2
				fmt.Print("Vertice 2: ")
				fmt.Scanf("%d", vertex[1])
				//peso de arista
				fmt.Print("Indique el peso de esta arista: ")
				fmt.Scanf("%d", &cost)
				//creando arista
				g.AddBothCost(vertex[0], vertex[1], cost)
			}
		} else {
			fmt.Print("grafo nulo")
		}
	} else { //grafo no dirigido
		fmt.Println("Desea añadir aristas al nodo? (0 para no 1 para si)")
		fmt.Scanf("%d", &addEdgeCheck)
		if addEdgeCheck == 1 {
			fmt.Print("La arista tendra peso? (0 para no, 1 para si): ")
			fmt.Scanf("%d", &edgeCost)
			if edgeCost == 0 {
				//origen y fin de arista
				fmt.Println("Indique de que vertice a que vertice se creara la arista.")
				fmt.Println("AVISO! El orden en que introduzca los vertices influira en la direccion de la arista (1 => 2)")
				//vertice 1
				fmt.Print("Vertice 1: ")
				fmt.Scanf("%d", vertex[0])
				//vertice 2
				fmt.Print("Vertice 2: ")
				fmt.Scanf("%d", vertex[1])
				//vertex[2] = nil
				//creando arista
				g.Add(vertex[0], vertex[1])
			} else {
				//origen y fin de arista
				fmt.Println("Indique de que vertice a que vertice se creara la arista.")
				fmt.Println("AVISO! El orden en que introduzca los vertices influira en la direccion de la arista (1 => 2)")
				//vertice 1
				fmt.Print("Vertice 1: ")
				fmt.Scanf("%d", vertex[0])
				//vertice 2
				fmt.Print("Vertice 2: ")
				fmt.Scanf("%d", vertex[1])
				//peso de arista
				fmt.Print("Indique el peso de esta arista: ")
				fmt.Scanf("%d", &cost)
				//creando arista
				g.AddCost(vertex[0], vertex[1], cost)
			}
		} else {
			fmt.Print("grafo nulo")
		}
	}
	/*g.AddBoth(0, 1) //  0 -- 1
	g.AddBoth(0, 2) //  |    |
	g.AddBoth(2, 3) //  2 -- 3
	g.AddBoth(1, 3)*/
	fmt.Println(g.String())
}
