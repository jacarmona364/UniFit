package rutina

//Rutina única con grupos musculares y ejercicios
type Rutina struct {

	identificador string // identificador unico, distingue las entidades
	ejerciciosPorGrupo map[string]map[string]Ejercicio // estructura que organiza ejercicios en grupos musculares
	
}


