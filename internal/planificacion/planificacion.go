package planificacion

// --- Estructura de la Agenda ---

// DiaSemana
type DiaSemana int

// dia de la semana asociado a un numero
const (
	Lunes DiaSemana = iota // 0
	Martes                 // 1
	Miercoles              // 2
	Jueves                 // 3
	Viernes                // 4
	Sabado                 // 5
	Domingo                // 6
)

// PlanSemanal asocia un EntrenamientoDiario a cada DiaSemana.
// Usar un 'map'
type PlanSemanal struct {
	Semana map[DiaSemana]EntrenamientoDiario
}