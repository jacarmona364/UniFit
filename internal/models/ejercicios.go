package models

// La dificultad de un ejercicio puede ser de tres tipos 
type Dificultad string
const (
	Beginner     Dificultad = "BEGINNER"
	Intermediate Dificultad = "INTERMEDIATE"
	Advanced     Dificultad = "ADVANCED"
)

// Ejercicios que pertenecen a un grupo muscular de la web
type Ejercicio struct {
	Nombre      string  // nombre del ejercicio obtenido de la web
	Dificultad  Dificultad // un nivel de dificultad
}




