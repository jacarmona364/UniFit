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
	nombre      string // nombre del ejercicio
	nivel  Dificultad  // dificultad que tiene 3 tipos
}


