package main

import "fmt"

var productosVendidos []string
var subtotales []float64

func RegistroVenta(nombre string, precio float64, cantidad int) {
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

		} else if opcion == 2 {
			MostrarEstadisticas()
		} else if opcion == 3 {
			break
		} else {
			fmt.Println("Opción inválida.")
		}
	}
}
