package models

// Grupos musculares que encontramos en la web
type GrupoMuscular struct {
	Nombre     string      // nombre del grupo muscular por el que se identifica en la web
	Ejercicios []Ejercicio // lista de ejercicios que contiene ese grupo
}
