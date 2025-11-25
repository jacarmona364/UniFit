package rutina

//Grupos musculares de la web para ejercitar
type GrupoMuscular struct {
	
	Nombre     string     //nombre del grupo muscular de la web 
	Ejercicios []Ejercicio // lista de ejercicos que compone un grupo muscular
	
	LimiteSobrecarga int // límite que no puede ser sobrepasado pues se produciría la sobrecarga
	CargaAcumulada int // valor que almacena como evoluciona la carga del grupo
	EstaSobrecargado bool // Estado que indica true si el musculo esta sobrecargado o false por defecto
	
}


