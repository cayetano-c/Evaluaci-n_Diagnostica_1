package main

import "fmt"

var productosVendidos []string
var subtotales []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)
}

func MostrarEstadisticas() {
	if len(subtotales) == 0 {
		fmt.Println("No existen ventas registradas.")
		return
	}

	total := 0.0

	for _, s := range subtotales {
		total += s
	}

	fmt.Println("Total recaudado:", total)
}

func main() {
	for {
		fmt.Println("\n--- Menú ---")
		fmt.Println("1. Registrar venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")

		var opcion int
		fmt.Scanln(&opcion)

		if opcion == 1 {
			fmt.Println("Productos disponibles:")
			fmt.Println("1. Arroz ($1.25)")
			fmt.Println("2. Leche ($0.95)")
			fmt.Println("3. Pan ($0.50)")

			var numProducto, cantidad int

			fmt.Print("Número de producto: ")
			fmt.Scanln(&numProducto)

			fmt.Print("Cantidad: ")
			fmt.Scanln(&cantidad)

			if numProducto == 1 {
				RegistrarVenta("Arroz", 1.25, cantidad)
			} else if numProducto == 2 {
				RegistrarVenta("Leche", 0.95, cantidad)
			} else if numProducto == 3 {
				RegistrarVenta("Pan", 0.50, cantidad)
			} else {
				fmt.Println("Producto inválido.")
			}

		} else if opcion == 2 {
			MostrarEstadisticas()

		} else if opcion == 3 {
			fmt.Println("Programa finalizado.")
			break

		} else {
			fmt.Println("Opción inválida.")
		}
	}
}
