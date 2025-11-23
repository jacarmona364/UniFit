package rutina

//Grupos musculares de la web para ejercitar
type GrupoMuscular struct {
	Nombre     string     //nombre del grupo muscular de la web 
	Ejercicios []Ejercicio // lista de ejericicos que lo compone
	Sobrecarga bool // Estado que indica true si el musculo esta sobrecargado o false por defecto
}


