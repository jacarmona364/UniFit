package models

// Definimos un tipo personalizado para los días (nuestro Enum)
type DiaSemana int

// Definimos las constantes usando iota para autoincrementar (Lunes=0, Martes=1, etc.)
const (
	Lunes DiaSemana = iota
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
	Domingo
)
