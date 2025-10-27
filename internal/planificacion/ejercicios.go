package planificacion

// --- Estructura del Ejercicio ---

type Ejercicio struct {
	NombreEjercicio string
	Descripcion     string
	Intensidad      int // intensidad del 1 al 10 por ejemplo de la dificultad
	MusculosObjetivo []Musculo // musculo o musculos que se ejercitarán
}
