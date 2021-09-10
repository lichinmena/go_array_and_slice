package main

import "fmt"

func main() {
	//Los arrays no se ven a menudo en programas go
	//Por que el tamaño de su matriz es parte de su tipo
	//Lo que limita su poder expresivo.
	//En go los array tienen un tamaño fijo y no puede alterarse

	var edades [3]uint8
	edades[0] = 1
	edades[1] = 2
	edades[2] = 3

	//Forma corta
	//edades := [3]uint8{8, 16, 32}

	fmt.Println(edades[0])
	fmt.Println(edades[1])
	fmt.Println(edades[2])

	for index, value := range edades {
		fmt.Printf("En el indice %d esta el valor %d\n", index, value)
	}

	//Longitud del array
	fmt.Printf("La longitud del array es : %d", len(edades))

	//Slices ===============================
	//Es una estructura ligera que encapsula y representa una porcion
	// de un array

	fmt.Println()
	fmt.Println()
	fmt.Println()

	//Primer modo de crear un slice
	frutas := []string{"Manzana", "peras"}
	fmt.Println(frutas)

	//Segunda forma
	verduras := make([]string, 2)
	verduras[0] = "Calabaza"
	verduras[1] = "Zanahoria"
	//Usamos make por que la declaracion de un slice
	//Implica mas que solo reservar un espacio en memoria
	//En concreto se tiene que separar memoria
	//Para un array subyacente, y ademas se tiene que inicializar el slice

	verduras2 := make([]string, 0, 10) //Indicar longitud y capacidad
	//append agrega elementos al slice sin importar su longitud actual
	//Append retorna un slice nuevo
	verduras2 = append(verduras2, "Brocolli")
	fmt.Println(verduras2)

	//redifinir slice
	verduras2 = verduras2[0:6]
	verduras2[5] = "Papa"
	fmt.Println(verduras2)

	edadesSlice := make([]uint8, 0, 10)
	edadesSlice = edadesSlice[0:6]
	edadesSlice[5] = 28
	fmt.Println(edadesSlice)

	//Append crea un slice

	edadesSlice2 := make([]uint8, 0, 5)
	//Para saber la capacidad actual de un slice
	fmt.Println(cap(edadesSlice2))

	//Para saber la longitud actual de un slice
	fmt.Println(len(edadesSlice2))

	testSlice := make([]uint8, 0, 5)
	capacidad := cap(testSlice)

	fmt.Printf("Capacidad Inicial : %d\n", capacidad)
	fmt.Println(testSlice)
	for i := uint8(0); i < 25; i++ {
		testSlice = append(testSlice, i)

		fmt.Println(testSlice)

		if cap(testSlice) != capacidad {
			capacidad = cap(testSlice)
			fmt.Printf("La capacidad es ahora: %d\n", capacidad)
		}
	}

	//Tercera forma de declarar un slice
	var sliceTercero []uint8
	fmt.Println(sliceTercero)

	//Cuarta forma de declarar un slice
	var sliceCuarto = []uint8{28, 32, 56, 45}
	fmt.Println(sliceCuarto)

	//Quinta forma
	sliceQuinto := []uint8{10, 20, 30, 40}
	fmt.Println(sliceQuinto)

}
