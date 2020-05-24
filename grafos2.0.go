package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

func main() {
	var eleccion elec
	var matrix [][]int
	var unionA [][2]int
	var matrizInci [][]int
	var xy [][2]float64
	var array [2]int
	var nodos int
	var nombreNodos []string

	graphType := eleccion.elecGrafo() //almacena la cantidad de nodos que tendra el grafo
	fmt.Print("\n¿Cuantos nodos desea introducir?: ")
	fmt.Scan(&nodos)
	//grafo := graph.New(nodos)

	//declaramos como objeto y construimos la matriz
	consMatriz := matrixConst{nodos, matrix, nombreNodos}
	matrix = consMatriz.matrizAd()

	//Pedimos que nombre los nodos y cordenadas
	nombreNodos = consMatriz.nombreNd()
	xy = eleccion.XandY(xy, nodos, nombreNodos)

	if graphType == 1 {

		//comenzamos a unir los grafos
		g := desarrollo{nodos, matrix, unionA, array, nombreNodos, xy, graphType}
		unionA = g.grafoN()
		matrizInci = consMatriz.matrizIn(matrizInci, unionA)

		//Mostramos matriz de adyacencia
		fmt.Println("----------------------------------")
		fmt.Println("- Quiera ver la matriz Adyacente?: ")
		graphType2 := eleccion.elecgraphType()

		if strings.ToLower(graphType2) == "si" {
			//se muestra la matriz
			g.mostrarAdyacente()

			fmt.Println("--------------------------------------------")
			fmt.Println("- Ahora ¿Quiera ver la matriz de Incidencia?: ")
			graphType2 = eleccion.elecgraphType()

			if strings.ToLower(graphType2) == "si" {

				g.mostrarIncidencia(matrizInci, unionA)
				fmt.Println("Muchas gracias por usar el programa")

			} else {

				fmt.Println("Ok, Muchas gracias por usar el programa")

			}

		} else {

			fmt.Println("----------------------------------")
			fmt.Println("- Ha eledigo -NO-")
			fmt.Println("- Quiera ver la matriz Adyacente?: ")
			graphType2 = eleccion.elecgraphType()

			if strings.ToLower(graphType2) == "si" {

				g.mostrarIncidencia(matrizInci, unionA)
				fmt.Println("Muchas gracias por usar el programa")

			} else {
				fmt.Println("Ok, Muchas gracias por usar el programa")
			}

		}

	} else if graphType == 2 {

		//comenzamos a unir los grafos DIRIGIDOS
		g := desarrollo{nodos, matrix, unionA, array, nombreNodos, xy, graphType}
		unionA = g.grafoD()
		matrizInci = consMatriz.matrizIn(matrizInci, unionA)

		//Mostramos matriz de adyacencia
		fmt.Println("----------------------------------")
		fmt.Println("- Quiera ver la matriz Adyacente?: ")
		graphType2 := eleccion.elecgraphType()

		if strings.ToLower(graphType2) == "si" {
			//se muestra la matriz
			g.mostrarAdyacente()

			fmt.Println("--------------------------------------------")
			fmt.Println("- Ahora ¿Quiera ver la matriz de Incidencia?: ")
			graphType2 = eleccion.elecgraphType()

			if strings.ToLower(graphType2) == "si" {

				g.mostrarIncidencia(matrizInci, unionA)
				fmt.Println("Muchas gracias por usar el programa")

			} else {

				fmt.Println("Ok, Muchas gracias por usar el programa")

			}

		} else {

			fmt.Println("----------------------------------")
			fmt.Println("- Ha eledigo -NO-")
			fmt.Println("- Quiera ver la matriz Adyacente?: ")
			graphType2 = eleccion.elecgraphType()

			if strings.ToLower(graphType2) == "si" {

				g.mostrarIncidencia(matrizInci, unionA)
				fmt.Println("Muchas gracias por usar el programa")

			} else {
				fmt.Println("Ok, Muchas gracias por usar el programa")
			}

		}
	}
}

//------|Comienza la estructura de las graphTypes por consola|-------

type elec struct { //variables universales para la structura que actua como clase
	eleccion  int
	graphType string
}

func (r elec) elecGrafo() int { //la R es un llamado para ejercutar la structura
	fmt.Println("\\Indique el tipo de grafo con el que se trabajara.")
	fmt.Print("(1 para Directed, 2 para Undirected \n")
	fmt.Print("Seleccione: ")
	fmt.Scan(&r.eleccion)

	for {
		if r.eleccion < 1 || r.eleccion > 2 {
			fmt.Println("\nIntrodujo un valor incorrecto. los valores son: 1 y 2")
			fmt.Print("1 - Grafo \n2 - Digrafo \n")
			fmt.Print("Seleccione: ")
			fmt.Scan(&r.eleccion)
		} else {
			break
		}
	}
	return r.eleccion
}

func (r elec) XandY(xy [][2]float64, nodos int, nombreNodos []string) [][2]float64 { //la R es un llamado para ejercutar la structura
	var x float64
	var y float64
	var array [2]float64
	var cont int

	fmt.Println("\n- Elija las cordenadas que tendran los nodos entre X = 1000/Y = 1000")
	fmt.Println("> Indique las coordenara X del nodo ", nombreNodos[cont], ": ")
	fmt.Scan(&x)
	for { //verificamos X
		if x < 1 || x > 999 {
			fmt.Println("-----------------------------------------------------------")
			fmt.Println("Valor o valores incorrentos. RECUERDA ---> X = 1000/Y = 1000")
			fmt.Println("> Indique las coordenara X del nodo ", nombreNodos[cont], ": ")
			fmt.Scan(&x)
			fmt.Println("-----------------------------------------------------------")
		} else {
			break
		}
	}

	fmt.Println("> Indique las coordenara Y del nodo ", nombreNodos[cont], ": ")
	fmt.Scan(&y)
	for { //verificamos Y
		if y < 1 || y > 999 {
			fmt.Println("-----------------------------------------------------------")
			fmt.Println("Valor o valores incorrentos. RECUERDA ---> X = 1000/Y = 1000")
			fmt.Println("> Indique las coordenara Y del nodo ", nombreNodos[cont], ": ")
			fmt.Scan(&y)
			fmt.Println("-----------------------------------------------------------")
		} else {
			break
		}
	}

	array[0] = x
	array[1] = y
	xy = append(xy, array)
	cont++

	for cont < nodos {

		fmt.Println("\n- Elija las cordenadas que tendran los nodos entre X = 1000/Y = 1000")
		fmt.Println("> Indique las coordenara X del nodo ", nombreNodos[cont], ": ")
		fmt.Scan(&x)
		for {
			if x < 1 || x > 999 {
				fmt.Println("-----------------------------------------------------------")
				fmt.Println("Valor o valores incorrentos. RECUERDA ---> X = 1000/Y = 1000")
				fmt.Println("> Indique las coordenara X del nodo ", nombreNodos[cont], ": ")
				fmt.Scan(&x)
				fmt.Println("-----------------------------------------------------------")
			} else {
				break
			}
		}

		fmt.Println("> Indique las coordenara Y del nodo ", nombreNodos[cont], ": ")
		fmt.Scan(&y)
		for {
			if y < 1 || y > 999 {
				fmt.Println("-----------------------------------------------------------")
				fmt.Println("Valor o valores incorrentos. RECUERDA ---> X = 1000/Y = 1000")
				fmt.Println("> Indique las coordenara Y del nodo ", nombreNodos[cont], ": ")
				fmt.Scan(&y)
				fmt.Println("-----------------------------------------------------------")
			} else {
				break
			}
		}
		array[0] = x
		array[1] = y
		xy = append(xy, array)
		cont++
	}
	fmt.Println(xy)
	return xy
}

func (r elec) elecgraphType() string {

	fmt.Println("----------------------------------")
	fmt.Print("- Seleccione (si/no): ")
	fmt.Scan(&r.graphType)

	return r.graphType
}

//------|Estructura de Construcion de matrices|-------

type matrixConst struct {
	nodos       int
	matrix      [][]int
	nombreNodos []string
}

func (m matrixConst) matrizAd() [][]int {
	m.matrix = make([][]int, m.nodos)
	for i := 0; i < m.nodos; i++ {
		m.matrix[i] = make([]int, m.nodos)
	}
	return m.matrix
}

func (m matrixConst) nombreNd() []string {
	fmt.Println("\n- Usted ha elegido ", m.nodos, " nodos.")
	fmt.Println("- Indique como quiere identificar cada nodo.")
	m.nombreNodos = make([]string, m.nodos)
	for i := 0; i < m.nodos; i++ {
		fmt.Print("- Identifique el nodo ", i+1, ": ")
		fmt.Scan(&m.nombreNodos[i])
	}
	return m.nombreNodos
}

func (m matrixConst) matrizIn(matrizInci [][]int, unionA [][2]int) [][]int {
	matrizInci = make([][]int, m.nodos)
	for i := 0; i < m.nodos; i++ {
		matrizInci[i] = make([]int, len(unionA))
	}

	return matrizInci
}

//------|Comienza estructura de grafos|-------

type desarrollo struct {
	nodos       int
	matrix      [][]int
	unionA      [][2]int
	array       [2]int
	nombreNodos []string
	xy          [][2]float64
	graphType   int
}

func (g desarrollo) grafoN() [][2]int {
	var n1 int //i
	var n2 int //j
	var letra1 string
	var letra2 string
	var muestra string
	var peso int64
	var x float64
	var y float64
	ultimaL := len(g.nombreNodos)

	// Creamos pantalla
	pantalla := gg.NewContext(1000, 1000)
	pantalla.SetRGB(255, 221, 80)

	for i := 0; i < g.nodos; i++ {
		for j := 0; j < 2; j++ {
			if j == 0 {
				x = g.xy[i][j]
			} else if j == 1 {
				y = g.xy[i][j]
			}
		}
		pantalla.DrawString(g.nombreNodos[i], x, y)
		pantalla.DrawCircle(x, y, 20)
	}

	fmt.Println("\n- Comienza la union de los Nodos con Aristas- ")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("- Recuerde que las posiciones son de [", g.nombreNodos[0], "] hasta [", g.nombreNodos[ultimaL-1], "]")
	fmt.Println("-----------------------------------------------------------")
	fmt.Print("- Nodo: ")
	fmt.Scan(&letra1)
	fmt.Print("- ¿Desde el nodo ", letra1, " hasta?: ")
	fmt.Scan(&letra2)

	for i := 0; i < g.nodos; i++ {
		if letra1 == g.nombreNodos[i] {
			n1 = i
		} else if letra2 == g.nombreNodos[i] {
			n2 = i
		}
	}
	pantalla.DrawLine(g.xy[n1][0], g.xy[n1][1], g.xy[n2][0], g.xy[n2][1])

	fmt.Print("- Diga el peso de la Arista que une el nodo[", letra1, "] Con el nodo[", letra2, "]: ")
	fmt.Scan(&peso)

	g.array[0] = n1
	g.array[1] = n2
	g.unionA = append(g.unionA, g.array)
	g.matrix[n1][n2] = 1

	//grafo.AddBothCost(n1, n2, peso)

	fmt.Println("\n----------------------------------------------------------------------------------")
	fmt.Println("- Grafo actual")
	fmt.Println(g.unionA)
	fmt.Println("- Nodos + Grafo actual + peso")
	//fmt.Println(grafo)
	fmt.Println("----------------------------------------------------------------------------------")

	fmt.Print("\n- ¿Desa agregar otra union de nodos (Aristas)? (si/no): ")
	fmt.Scan(&muestra)

	for strings.ToLower(muestra) == "si" {

		fmt.Println("- Recuerde que las posiciones son de [", g.nombreNodos[0], "] hasta [", g.nombreNodos[ultimaL-1], "]")
		fmt.Println("-----------------------------------------------------------")
		fmt.Print("- Nodo: ")
		fmt.Scan(&letra1)
		fmt.Print("- ¿Desde el nodo ", letra1, " hasta?: ")
		fmt.Scan(&letra2)

		for i := 0; i < g.nodos; i++ {
			if letra1 == g.nombreNodos[i] {
				n1 = i
			} else if letra2 == g.nombreNodos[i] {
				n2 = i
			}
		}
		pantalla.DrawLine(g.xy[n1][0], g.xy[n1][1], g.xy[n2][0], g.xy[n2][1])

		fmt.Print("- Diga el peso de la Arista que une el nodo[", letra1, "] Con el nodo[", letra2, "]: ")
		fmt.Scan(&peso)

		g.array[0] = n1
		g.array[1] = n2
		g.unionA = append(g.unionA, g.array)

		//grafo.AddBothCost(n1, n2, peso)

		fmt.Println("\n-----------------------------------------------------------")
		fmt.Println("- Grafo actual")
		fmt.Println(g.unionA)
		fmt.Println("- Nodos + Grafo actual + peso")
		//fmt.Println(grafo)
		fmt.Println("-------------------------------------------------------------")

		if g.matrix[n1][n2] >= 1 {
			g.matrix[n1][n2] = g.matrix[n1][n2] + 1
		} else {
			g.matrix[n1][n2] = 1
		}

		fmt.Print("\n- ¿Desa agregar otra union de nodos (Aristas)? (si/no): ")
		fmt.Scan(&muestra)
	}
	pantalla.SetLineWidth(1)
	pantalla.Stroke()
	pantalla.SavePNG("out.png")

	return g.unionA
}

func (g desarrollo) grafoD() [][2]int {
	var n1 int //i
	var n2 int //j
	var letra1 string
	var letra2 string
	var muestra string
	var peso int64
	var x float64
	var y float64
	ultimaL := len(g.nombreNodos)

	// Creamos pantalla
	pantalla := gg.NewContext(1000, 1000)
	pantalla.SetRGB(255, 221, 80)

	for i := 0; i < g.nodos; i++ {
		for j := 0; j < 2; j++ {
			if j == 0 {
				x = g.xy[i][j]
			} else if j == 1 {
				y = g.xy[i][j]
			}
		}
		pantalla.DrawString(g.nombreNodos[i], x, y)
		pantalla.DrawCircle(x, y, 20)
	}

	fmt.Println("- Comienza la union de los Nodos con Aristas [Con entorno de DIRECCION]- ")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("- Recuerde que las posiciones son de [", g.nombreNodos[0], "] hasta [", g.nombreNodos[ultimaL-1], "]")
	fmt.Println("- Tambien que la direccion solo sera de un sentido [---->] [", g.nombreNodos[0], "] hasta [", g.nombreNodos[ultimaL-1], "]")
	fmt.Println("- Ejemplo: [A] ------> [B]")
	fmt.Println("-----------------------------------------------------------")
	fmt.Print("- Desde el Nodo: ")
	fmt.Scan(&letra1)
	fmt.Print("- ¿Desde el nodo ", letra1, " hasta? (Distino Flecha): ")
	fmt.Scan(&letra2)

	for i := 0; i < g.nodos; i++ {
		if letra1 == g.nombreNodos[i] {
			n1 = i
		} else if letra2 == g.nombreNodos[i] {
			n2 = i
		}
	}
	pantalla.DrawLine(g.xy[n1][0], g.xy[n1][1], g.xy[n2][0], g.xy[n2][1])

	fmt.Print("- Diga el peso de la Arista que une el nodo[", letra1, "] Con el nodo[", letra2, "]: ")
	fmt.Scan(&peso)

	g.array[0] = n1
	g.array[1] = n2
	g.unionA = append(g.unionA, g.array)
	g.matrix[n1][n2] = 1

	//grafo.AddBothCost(n1, n2, peso)

	fmt.Println("\n----------------------------------------------------------------------------------")
	fmt.Println("- Grafo actual")
	fmt.Println(g.unionA)
	fmt.Println("- Nodos + Grafo actual + peso")
	//fmt.Println(grafo)
	fmt.Println("----------------------------------------------------------------------------------")

	fmt.Print("\n- ¿Desa agregar otra union de nodos (Aristas)? (si/no): ")
	fmt.Scan(&muestra)

	for strings.ToLower(muestra) == "si" {

		fmt.Println("- Recuerde que las posiciones son de [", g.nombreNodos[0], "] hasta [", g.nombreNodos[ultimaL-1], "]")
		fmt.Println("-----------------------------------------------------------")
		fmt.Print("- Nodo: ")
		fmt.Scan(&letra1)
		fmt.Print("- ¿Desde el nodo ", letra1, " hasta?: ")
		fmt.Scan(&letra2)

		for i := 0; i < g.nodos; i++ {
			if letra1 == g.nombreNodos[i] {
				n1 = i
			} else if letra2 == g.nombreNodos[i] {
				n2 = i
			}
		}
		pantalla.DrawLine(g.xy[n1][0], g.xy[n1][1], g.xy[n2][0], g.xy[n2][1])

		fmt.Print("- Diga el peso de la Arista que une el nodo[", letra1, "] Con el nodo[", letra2, "]: ")
		fmt.Scan(&peso)

		g.array[0] = n1
		g.array[1] = n2
		g.unionA = append(g.unionA, g.array)

		//grafo.AddBothCost(n1, n2, peso)

		fmt.Println("\n-----------------------------------------------------------")
		fmt.Println("- Grafo actual")
		fmt.Println(g.unionA)
		fmt.Println("- Nodos + Grafo actual + peso")
		//fmt.Println(grafo)
		fmt.Println("-------------------------------------------------------------")

		if g.matrix[n1][n2] >= 1 {
			g.matrix[n1][n2] = g.matrix[n1][n2] + 1
		} else {
			g.matrix[n1][n2] = 1
		}

		fmt.Print("\n- ¿Desa agregar otra union de nodos (Aristas)? (si/no): ")
		fmt.Scan(&muestra)
	}
	pantalla.SetLineWidth(1)
	pantalla.Stroke()
	pantalla.SavePNG("out.png")

	return g.unionA
}

func (g desarrollo) mostrarAdyacente() {
	fmt.Println("\n------Matriz Adyacente------")

	//Mostramos [a][b][c]... en horizontal

	for i := 0; i < len(g.matrix); i++ {
		if i == 0 {
			fmt.Print("    [" + g.nombreNodos[i] + "]")
		} else if i > 0 {
			fmt.Print("[" + g.nombreNodos[i] + "]")
		}
	}

	//Mostramos matriz Adyacente
	sumeA := make([]int, g.nodos)
	fmt.Println()
	for i := 0; i < len(g.matrix); i++ {
		fmt.Print("[" + g.nombreNodos[i] + "]-")
		for j := 0; j < len(g.matrix); j++ {
			fmt.Print("[" + strconv.Itoa(g.matrix[i][j]) + "]")
			sumeA[i] = sumeA[i] + g.matrix[i][j]
		}
		fmt.Print("-[" + strconv.Itoa(sumeA[i]) + "]")
		fmt.Println()
	}
	fmt.Println("\n- La ultima Columna, Representa los -Grados- de las aristas.")
}

func (g desarrollo) mostrarIncidencia(matrizInci [][]int, unionA [][2]int) {

	//------
	aristas := make([]int, len(unionA))
	suma := 0
	fmt.Println("\n------Matriz de Incidencia------")
	for i := 0; i < len(unionA); i++ {
		if i == 0 {
			aristas[i] = 1
			suma = suma + aristas[i]
			fmt.Print("    [E" + strconv.Itoa(aristas[i]) + "]")
		} else if i > 0 {
			suma = suma + 1
			aristas[i] = suma
			fmt.Print("[E" + strconv.Itoa(aristas[i]) + "]")
		}
	}
	fmt.Println()
	for i := 0; i < len(aristas); i++ {
		if i == 0 {
			fmt.Print("     ")
		} else if i > 0 {
			fmt.Print(".  ")
		}
	}

	if g.graphType == 1 {
		valida := 0
		valida2 := 0
		for i := 0; i < g.nodos; i++ {
			for j := 0; j < len(unionA); j++ {

				valida = unionA[j][0]
				valida2 = unionA[j][1]
				if i == valida || i == valida2 {
					matrizInci[i][j] = 1
				}

			}

		}
	} else if g.graphType == 2 {

		valida := 0
		valida2 := 0

		for i := 0; i < g.nodos; i++ {
			for j := 0; j < len(unionA); j++ {

				valida = unionA[j][0]
				valida2 = unionA[j][1]

				if i == valida {
					matrizInci[i][j] = 1
				} else if i == valida2 {
					matrizInci[i][j] = -1
				}

			}

		}
	}

	//usamos la matriz

	fmt.Println()
	for i := 0; i < g.nodos; i++ {
		fmt.Print("[" + g.nombreNodos[i] + "]-")
		for j := 0; j < len(unionA); j++ {
			fmt.Print("[" + strconv.Itoa(matrizInci[i][j]) + "] ")
		}
		fmt.Println()
	}

	if g.graphType == 1 {
		fmt.Println("- Es un grado No dirigido de ", g.nodos, " nodos y ", len(unionA), " Aristas")
	} else if g.graphType == 2 {
		fmt.Println("- Es un Digrado de ", g.nodos, " nodos y ", len(unionA), " Aristas [DIRIGIDAS]")
	}

}
