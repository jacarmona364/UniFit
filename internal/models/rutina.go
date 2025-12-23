package models
type GrupoMuscular string
const (
	Abductors GrupoMuscular = "Abductors"
	Abs       GrupoMuscular = "Abs"
	Biceps    GrupoMuscular = "Biceps"
	Chest     GrupoMuscular = "Chest"
	//... se deben añadir los grupos a trabajar
)

//Entidad única que gestiona los datos obtenidos
type Rutina struct {
	catalogoEjercicios map[GrupoMuscular][]Ejercicio //conjunto de datos que tenemos de la web, ejercicios asociados a un grupo muscular
}
