package rutina

// TODO: Definir el límite real. Pongo 100 temporalmente para que compile pero debe ser un valor del cliente
const LimiteSobrecarga = 100

//Grupos musculares de la web para ejercitar
type GrupoMuscular struct {
	
	Nombre     string     //nombre del grupo muscular de la web 
	Ejercicios []Ejercicio // lista de ejericicos que lo compone
	SobrecargaAcumulada int // valor que almacena como evoluciona la sobrecarga del grupo
	Sobrecarga bool // Estado que indica true si el musculo esta sobrecargado o false por defecto
	
}


