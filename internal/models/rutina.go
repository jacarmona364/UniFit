package models

import (
	"time"
)

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

// Rutina asociada a una fecha
type Rutina struct {

	fecha time.Time // Identificador unico de la semana 
	rutinaSemanal map[DiaSemana]GrupoMuscular //cada dia de la semana tiene un grupo muscular
}
