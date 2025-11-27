package models
// Entidad 
//Única por usuario
type Rutina struct {
	catalogoEjercicios map[GrupoMuscular][]Ejercicio //conjunto de datos que tenemos de la web, ejercicios asociados a un grupo muscular
}
