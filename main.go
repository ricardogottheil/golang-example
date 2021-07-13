package main

import (
	"fmt"

	cf "CursoProfesionalGo/CodigoFacilito"
)



func main() {
	curso := cf.Curso { Titulo: "Curso profesional de go" }

	fmt.Println(curso.GetTitulo())
}