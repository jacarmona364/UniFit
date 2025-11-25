package rutina 

//Añado un enum para los distintos dias de la semana de la planificacion
type DiaSemana string

const (
	Lunes     DiaSemana = "LUNES"
	Martes    DiaSemana = "MARTES"
	Miercoles DiaSemana = "MIERCOLES"
	Jueves    DiaSemana = "JUEVES"
	Viernes   DiaSemana = "VIERNES"
	Sabado    DiaSemana = "SABADO"
	Domingo   DiaSemana = "DOMINGO"
)

//Estructura de la entidad planificadora
type Planificacion struct {
	Propietario       string // Persona a la que le pertenece la planificacion 
	RutinaSemanal map[DiaSemana][]GrupoMuscular // map que asocia a cada día grupos musculares
}




