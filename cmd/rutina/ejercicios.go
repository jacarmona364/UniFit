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
	
	nivel  Dificultad  // dificultad que puede ser de 3 tipos
	intensidad int // intensidad del ejercicio calculada de la dificultad
	
}


