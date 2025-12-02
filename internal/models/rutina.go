package models
type GrupoMuscular int
const (
	// iota empieza en 0.
	Abductors GrupoMuscular = iota
	Abs
	Biceps
	//... se deben añadir los grupos a trabajar
)

//Entidad única que gestiona los datos obtenidos
type Rutina struct {
	catalogoEjercicios map[GrupoMuscular][]Ejercicio //conjunto de datos que tenemos de la web, ejercicios asociados a un grupo muscular
}
