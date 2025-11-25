package rutina

//Grupos musculares de la web para ejercitar
type GrupoMuscular struct {
	
	Nombre     string     //nombre del grupo muscular de la web 
	Ejercicios []Ejercicio // lista de ejericicos que lo compone
	LimiteSobrecarga int // límite que no puede ser sobrepasado pues se produciría la sobrecarga
	SobrecargaAcumulada int // valor que almacena como evoluciona la sobrecarga del grupo
	Sobrecarga bool // Estado que indica true si el musculo esta sobrecargado o false por defecto
	
}


