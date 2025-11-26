package rutina

// Ejercicios asociados a un grupo muscular

// La dificultad puede ser de tres tipos
type Dificultad string
const (
	Beginner     Dificultad = "BEGINNER"
	Intermediate Dificultad = "INTERMEDIATE"
	Advanced     Dificultad = "ADVANCED"
)

// Estructura de ejercicio
type Ejercicio struct {
	Nombre      string // nombre del ejercicio
	Nivel  Dificultad  // dificultad que tiene 3 tipos
}


