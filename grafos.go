package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fogleman/gg"

	//importando paquete de grafos de github
	"github.com/yourbasic/graph"
)

var vertexCount int     //almacena la cantidad de nodos que tendra el grafo
var graphType int       //si el grafo sera undirected o directed
var addEdgeCheck string //bandera para añadir aristas
var edgeCost string     //para definir si la arista tendra peso o no
var vertex [2]int       //de donde a donde va la arista sin peso
var cost int64          //costo de la arista
var tipoGrafo string
var cuentaAviso int = 0 //Variable que permite llevar el tiempo del aviso
var xy [2]float64       //coordenadas nodos para plot
var pares [][2]int      //posiciones de aristas plot
//coleccion para matriz de incidencia
var coleccion [][2]int //slice de vectores

func main() {

	//Función de bienvenida y algunas instrucciones
	welcomeAndInstrucciones()

	//tamaño del grafo
	fmt.Print("\nIntroduzca la cantidad de 'Nodos' o ' Vertices' que tendra el Grafo: ")
	fmt.Scanf("%d", &vertexCount)

	//solicito coordenadas de los nodos
	var nodeCoor [][2]float64 //almacena los xy de cada nodo
	for i := 0; i < vertexCount; i++ {
		fmt.Print("Ingrese la coordenada X del nodo ", i, ": ")
		fmt.Scanf("%g", &xy[0])
		fmt.Print("Ingrese la coordenada Y del nodo ", i, ": ")
		fmt.Scanf("%g", &xy[1])
		nodeCoor = append(nodeCoor, xy)
	}

	//construyo el grafo con el tamaño solicitado antes
	g := graph.New(vertexCount)
	//matriz de adyacencia
	matrix := make([][]int, vertexCount)
	for i := 0; i < vertexCount; i++ {
		matrix[i] = make([]int, vertexCount)
	}
	// //coleccion para matriz de incidencia
	// var coleccion [][2]int //slice de vectores
	//pregunto que tipo de grafo sera
	fmt.Print("¿El grafo será undirected o directed? \nIngrese 0 para undirected, 1 para directed: ")
	fmt.Scanf("%d", &graphType)

	//grafo no dirigido
	if graphType == 0 {
		tipoGrafo = "No Dirigido"
		fmt.Print("Desea añadir aristas al nodo? (y/n): ")
		fmt.Scanf("%s", &addEdgeCheck)
		//si o no añadir aristas al nodo
		if strings.ToLower(addEdgeCheck) == "y" {
			//emulando while con for infinito que sera detenido con un BREAK statement
			for {
				fmt.Print("La arista tendra peso? (y/n): ")
				fmt.Scanf("%s", &edgeCost)

				if strings.ToLower(edgeCost) == "n" {
					//origen y fin de arista
					//Revisando si ya se hizo el aviso de empezar por el nodo 0
					if cuentaAviso == 0 {
						recordatorioCero(vertexCount)
					}
					fmt.Println("\nIndique de que vertice a que vertice se creara la arista: ")
					//vertice 1
					fmt.Print("\nVertice 1: ")
					fmt.Scanf("%d", &vertex[0])
					//vertice 2
					fmt.Print("\nVertice 2: ")
					fmt.Scanf("%d", &vertex[1])
					//creando arista
					g.AddBoth(vertex[0], vertex[1])
					//fill matriz de adyacencia
					matrix[vertex[0]][vertex[1]] = 1
					//fill matriz de incidencia
					coleccion = append(coleccion, vertex)
					cuentaAviso = cuentaAviso + 1
				} else {
					//origen y fin de arista
					//Revisando si ya se hizo el aviso de empezar por el nodo 0
					if cuentaAviso == 0 {
						recordatorioCero(vertexCount)
					}
					fmt.Println("\nIndique de que vertice a que vertice se creara la arista:")
					//vertice 1
					fmt.Print("\nVertice 1: ")
					fmt.Scanf("%d", &vertex[0])
					//vertice 2
					fmt.Print("\nVertice 2: ")
					fmt.Scanf("%d", &vertex[1])
					//peso de arista
					fmt.Print("\nIndique el peso de esta arista: ")
					fmt.Scanf("%d", &cost)
					//creando arista
					g.AddBothCost(vertex[0], vertex[1], cost)
					matrix[vertex[0]][vertex[1]] = 1
					coleccion = append(coleccion, vertex)
					cuentaAviso = cuentaAviso + 1
				}
				fmt.Print("\nDesea añadir otra arista? (y/n): ")
				fmt.Scanf("%s", &addEdgeCheck)
				if addEdgeCheck == "n" {
					break
				}
			}
		} else {
			tipoGrafo = "Grafo Nulo"
		}
	} else { //grafo dirigido
		tipoGrafo = "Grafo Dirgido"
		fmt.Print("Desea añadir aristas al nodo? (y/n): ")
		fmt.Scanf("%s", &addEdgeCheck)

		if addEdgeCheck == "y" {
			for {
				fmt.Print("La arista tendra peso? (y/n): ")
				fmt.Scanf("%s", &edgeCost)
				if edgeCost == "n" {
					//origen y fin de arista
					fmt.Println("Indique de que vertice a que vertice se creara la arista.")
					//Revisando si ya se hizo el aviso de empezar por el nodo 0
					if cuentaAviso == 0 {
						recordatorioCero(vertexCount)
					}
					fmt.Println("\n¡AVISO! El orden en que introduzca los vertices influira en la direccion de la arista (1 => 2)")
					//vertice 1
					fmt.Print("\nVertice 1: ")
					fmt.Scanf("%d", &vertex[0])
					//vertice 2
					fmt.Print("\nVertice 2: ")
					fmt.Scanf("%d", &vertex[1])
					//vertex[2] = nil
					//creando arista
					g.Add(vertex[0], vertex[1])
					matrix[vertex[0]][vertex[1]] = 1
					coleccion = append(coleccion, vertex)
					cuentaAviso = cuentaAviso + 1
				} else {
					//origen y fin de arista
					fmt.Println("Indique de que vertice a que vertice se creara la arista.")
					//Revisando si ya se hizo el aviso de empezar por el nodo 0
					if cuentaAviso == 0 {
						recordatorioCero(vertexCount)
					}
					fmt.Println("\n¡AVISO! El orden en que introduzca los vertices influira en la direccion de la arista (1 => 2)")
					//vertice 1
					fmt.Print("\nVertice 1: ")
					fmt.Scanf("%d", &vertex[0])
					//vertice 2
					fmt.Print("\nVertice 2: ")
					fmt.Scanf("%d", &vertex[1])
					//peso de arista
					fmt.Print("\nIndique el peso de esta arista: ")
					fmt.Scanf("%d", &cost)
					//creando arista
					g.AddCost(vertex[0], vertex[1], cost)
					matrix[vertex[0]][vertex[1]] = 1
					coleccion = append(coleccion, vertex)
					cuentaAviso = cuentaAviso + 1
				}
				fmt.Print("\nDesea añadir otra arista? (y/n): ")
				fmt.Scanf("%s", &addEdgeCheck)
				if addEdgeCheck == "n" {
					break
				}
			}
		} else {
			tipoGrafo = "Grafo Nulo"
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
	fmt.Println("La cantidad de nodos y la cantidad de aristas están representadas en una lista.")
	fmt.Println("Cada elemento de la lista posee dos enteros dentro de él.\nEstos indican de cuál vértice a cuál vértice van las aristas.")
	fmt.Println("De tener ponderación, las aristas se mostrarán luego de cada una de ellas seguidas de ':'")
	fmt.Println("\n====================================================================\n")
	fmt.Println(g)
	fmt.Printf("\nTipo de Grafo: %s\n", tipoGrafo)
	fmt.Println("\n====================================================================\n")
	fmt.Println("                Matriz de Adyacencia\n")
	//imprimiendo matriz de adyacencia

	for i := 0; i < vertexCount; i++ {
		if i == 0 {
			fmt.Print("                     [", i, "]")
		} else {
			fmt.Print("[", i, "]")
		}
	}
	fmt.Println("")
	for i := 0; i < vertexCount; i++ {
		fmt.Print("		  [", i, "] ")
		for j := 0; j < vertexCount; j++ {
			fmt.Printf("%d  ", matrix[i][j])
		}
		fmt.Println(" ")
	}
	//imprimiendo matriz de incidencia
	fmt.Println("\n====================================================================\n")
	fmt.Println("                Matriz de Incidencia\n")

	for i := 0; i < vertexCount; i++ {
		if i == 0 {
			fmt.Print("                     [", i, "]")
		} else {
			fmt.Print("[", i, "]")
		}
	}
	fmt.Println("")
	for i := 0; i < vertexCount; i++ {
		fmt.Print("		  [", i, "] ")
		for j := 0; j < len(coleccion); j++ {
			fmt.Printf("%d  ", matrix1[j][i])
		}
		fmt.Println(" ")
	}
	fmt.Println("\n====================================================================\n")

	plot(nodeCoor, coleccion)
} //Fin funión main

//Creando función con bienvenida e instrucciones
func welcomeAndInstrucciones() {
	fmt.Println("\n¡Bienvenido a Grafos! Programa elaborado por:\nSebastián Avendaño C.I: 26.765.567.\nSebastián Álvarez  C.I: 26.900.740.")
	fmt.Println("Cátedra: Estructuras Discretas II. Profesor Ing. Jetro López. \nIngeniería en Computación - Universidad José Antonio Páez.")
	fmt.Println("\nA continuación, deberá ingresar el número de nodos que tendrá el grafo, así como el tipo del grafo.")
} //Final de función con bienvenida e instrucciones

//Función con instrucciones para la introducción de aristas
func recordatorioCero(nodos int) {
	fmt.Println("\n¡Recuerde que debe comenzar por el primer nodo (el nodo 0) y debe llegar hasta el último nodo, es decir, el nodo ", (nodos - 1), "!")
	fmt.Println("En caso contrario, podría caer en condición de error")
} //Fin función de instrucciones para la introduccion de aristas

//creando canvas
func plot(nodeCoor [][2]float64, coleccion [][2]int) {
	//seteando canvas
	canvas := gg.NewContext(1000, 1000)
	canvas.SetRGB(255, 221, 80)

	//pintar nodos
	var x, y float64
	for i := 0; i < len(nodeCoor); i++ {
		for j := 0; j < 2; j++ {
			if j == 0 {
				x = nodeCoor[i][j]
			} else if j == 1 {
				y = nodeCoor[i][j]
			}
		}
		canvas.DrawString(strconv.Itoa(i), x, y)
		canvas.DrawCircle(x, y, 20)
	}

	//pintar lineas
	var x1, y1, x2, y2 float64
	for i := 0; i < len(coleccion); i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				if j == 0 {
					x1 = nodeCoor[coleccion[i][j]][0]
					y1 = nodeCoor[coleccion[i][j]][1]
				} else {
					x2 = nodeCoor[coleccion[i][j]][0]
					y2 = nodeCoor[coleccion[i][j]][1]
				}
			}
		}
		canvas.DrawLine(x1, y1, x2, y2)
	}

	canvas.SetLineWidth(2)
	canvas.Stroke()
	canvas.SavePNG("out.png")

}
