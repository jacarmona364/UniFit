package planificacion

import "fmt"

//Errores para validar la rutina
var (
	ErrNombreVacio        = errors.New("el nombre de la rutina no puede estar vacío")
	ErrIntensidadInvalida = errors.New("la intensidad de la rutina debe estar entre 1 y 5") //por ejemplo
	ErrSinEjercicios      = errors.New("la rutina debe tener al menos un ejercicio")
)

// --- Estructura de Entrenamiento de un día ---

type EntrenamientoDiario struct {
	Nombre        string
	Intensidad    int // 0 (Descanso) a 5 (Máximo)
	Ejercicios    []Ejercicio
	GrupoMuscular GrupoMuscular
}

// --- Constructores de EntrenamientoDiario ---

//Un día de descanso no tendrá intensidad, ni ejercicios ni grupo muscular
func NewEntrenamientoDescanso() EntrenamientoDiario {
	return EntrenamientoDiario{
		Nombre:        "Descanso",
		Intensidad:    0,
		Ejercicios:    nil, // nil es un slice vacío
		GrupoMuscular: GrupoMuscular{Tipo: Ninguno},
	}
}

// un día de futbol siempre se entrenará las piernas, sin ejercicios y con la misma intensidad
func NewEntrenamientoFutbol() EntrenamientoDiario {
	return EntrenamientoDiario{
		Nombre:        "Fútbol",
		Intensidad:    4,
		Ejercicios:    nil, // El fútbol no tiene "ejercicios" de gimnasio
		GrupoMuscular: GrupoMuscular{Tipo: TrenInferior},
	}
}

//podemos crear una rutina de ejercicios 
func NewEntrenamientoRutina(nombre string, intensidad int, grupo GrupoMuscular, ejercicios []Ejercicio) EntrenamientoDiario {
	return EntrenamientoDiario{
		Nombre:        nombre,
		Intensidad:    intensidad,
		Ejercicios:    ejercicios,
		GrupoMuscular: grupo,
	}
}
