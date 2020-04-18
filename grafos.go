package main

import (
	"fmt"
	//importando paquete de grafos de github
	"github.com/yourbasic/graph"
)

var vertexCount int     //almacena la cantidad de nodos que tendra el grafo
var graphType int       //si el grafo sera undirected o directed
var addEdgeCheck string //bandera para añadir aristas
var edgeCost string     //para definir si la arista tendra peso o no
var vertex [2]int       //de donde a donde va la arista sin peso
var cost int64          //costo de la arista

func main() {

	//tamaño del grafo
	fmt.Print("Introduzca la cantidad de 'Nodos' o ' Vertices' que tendra el Grafo: ")
	fmt.Scanf("%d", &vertexCount)
	//construyo el grafo con el tamaño solicitado antes
	g := graph.New(vertexCount)
	//matriz de adyacencia
	matrix := make([][]int, vertexCount)
	for i := 0; i < vertexCount; i++ {
		matrix[i] = make([]int, vertexCount)
	}
	//coleccion para matriz de incidencia
	var coleccion [][2]int
	var element [2]int
	//pregunto que tipo de grafo sera
	fmt.Println("El grafo será undirected o directed? (0 para undirected, 1 para directed)")
	fmt.Scanf("%d", &graphType)

	//grafo no dirigido
	if graphType == 0 {
		fmt.Println("Desea añadir aristas al nodo? (y/n)")
		fmt.Scanf("%s", &addEdgeCheck)
		//si o no añadir aristas al nodo
		if addEdgeCheck == "y" {
			//emulando while con for infinito que sera detenido con un BREAK statement
			for {
				fmt.Print("La arista tendra peso? (y/n): ")
				fmt.Scanf("%s", &edgeCost)

				if edgeCost == "n" {
					//origen y fin de arista
					fmt.Println("Indique de que vertice a que vertice se creara la arista, no importa el orden")
					//vertice 1
					fmt.Print("Vertice 1: ")
					fmt.Scanf("%d", &vertex[0])
					//vertice 2
					fmt.Print("Vertice 2: ")
					fmt.Scanf("%d", &vertex[1])
					//creando arista
					g.AddBoth(vertex[0], vertex[1])
					//fill matriz de adyacencia
					matrix[vertex[0]][vertex[1]] = 1
					//fill matriz de incidencia
					element[0] = vertex[0]
					element[1] = vertex[1]
					coleccion = append(coleccion, element)
				} else {
					//origen y fin de arista
					fmt.Println("Indique de que vertice a que vertice se creara la arista, no importa el orden")
					//vertice 1
					fmt.Print("Vertice 1: ")
					fmt.Scanf("%d", &vertex[0])
					//vertice 2
					fmt.Print("Vertice 2: ")
					fmt.Scanf("%d", &vertex[1])
					//peso de arista
					fmt.Print("Indique el peso de esta arista: ")
					fmt.Scanf("%d", &cost)
					//creando arista
					g.AddBothCost(vertex[0], vertex[1], cost)
					matrix[vertex[0]][vertex[1]] = 1
				}
				fmt.Print("Desea añadir otra arista? (y/n): ")
				fmt.Scanf("%s", &addEdgeCheck)
				if addEdgeCheck == "n" {
					break
				}
			}
		} else {
			fmt.Println("grafo nulo\n")
		}
	} else { //grafo dirigido
		fmt.Println("Desea añadir aristas al nodo? (y/n)")
		fmt.Scanf("%s", &addEdgeCheck)

		if addEdgeCheck == "y" {
			for {
				fmt.Print("La arista tendra peso? (y/n): ")
				fmt.Scanf("%s", &edgeCost)
				if edgeCost == "n" {
					//origen y fin de arista
					fmt.Println("Indique de que vertice a que vertice se creara la arista.")
					fmt.Println("AVISO! El orden en que introduzca los vertices influira en la direccion de la arista (1 => 2)")
					//vertice 1
					fmt.Print("Vertice 1: ")
					fmt.Scanf("%d", &vertex[0])
					//vertice 2
					fmt.Print("Vertice 2: ")
					fmt.Scanf("%d", &vertex[1])
					//vertex[2] = nil
					//creando arista
					g.Add(vertex[0], vertex[1])
					matrix[vertex[0]][vertex[1]] = 1
				} else {
					//origen y fin de arista
					fmt.Println("Indique de que vertice a que vertice se creara la arista.")
					fmt.Println("AVISO! El orden en que introduzca los vertices influira en la direccion de la arista (1 => 2)")
					//vertice 1
					fmt.Print("Vertice 1: ")
					fmt.Scanf("%d", &vertex[0])
					//vertice 2
					fmt.Print("Vertice 2: ")
					fmt.Scanf("%d", &vertex[1])
					//peso de arista
					fmt.Print("Indique el peso de esta arista: ")
					fmt.Scanf("%d", &cost)
					//creando arista
					g.AddCost(vertex[0], vertex[1], cost)
					matrix[vertex[0]][vertex[1]] = 1
				}
				fmt.Print("Desea añadir otra arista? (y/n): ")
				fmt.Scanf("%s", &addEdgeCheck)
				if addEdgeCheck == "n" {
					break
				}
			}
		} else {
			fmt.Println("grafo nulo\n")
		}
	}

	//CODIGO GUIA NO BORRAR
	/*g.AddBoth(0, 1)// 0 -- 1
	g.AddBoth(0, 2) //  |    |
	g.AddBoth(2, 3) //  2 -- 3
	g.AddBoth(1, 3)*/

	// creando matriz de incidecia
	matrix1 := make([][]int, vertexCount)
	for i := 0; i < vertexCount; i++ {
		matrix1[i] = make([]int, len(coleccion))
	}

	// rellenando matriz de incidencia
	//3 for para recorrer filas y columnas de la matriz y el array dentro de cada elemento de la matriz
	for i := 0; i < vertexCount; i++ {
		for j := 0; j < len(coleccion); j++ {
			for k := 0; k < 2; k++ {
				matrix1[j][coleccion[j][k]] = 1
			}
		}
	}

	fmt.Println("\n\n==================================================================\n")
	fmt.Println("Cantidad de nodos y Cantidad de aristas representadas en una lista")
	fmt.Println("Cada elemento de la lista posee dos enteros dentro de el.\nEstos indican de que vertice a que vertice van las aristas.")
	fmt.Println("\n====================================================================\n")
	fmt.Println(g)
	fmt.Println("\n====================================================================\n")
	fmt.Println("Matriz de Adyacencia")
	fmt.Println("\n====================================================================\n")

	//imprimiendo matriz de adyacencia
	for i := 0; i < vertexCount; i++ {
		for j := 0; j < vertexCount; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println(" ")
	}

	//imprimiendo matriz de incidencia
	fmt.Println("\n====================================================================\n")
	fmt.Println("Matriz de Incidencia")
	fmt.Println("\n====================================================================\n")
	for i := 0; i < vertexCount; i++ {
		for j := 0; j < len(coleccion); j++ {
			fmt.Printf("%d ", matrix1[i][j])
		}
		fmt.Println(" ")
	}
}
