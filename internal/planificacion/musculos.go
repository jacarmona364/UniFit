package planificacion

// --- Estructuras Musculares ---

type Musculo struct {
	Nombre string
}

type TipoGrupoMuscular string

const (
	TrenInferior TipoGrupoMuscular = "Tren inferior"
	TrenSuperior TipoGrupoMuscular = "Tren superior"
	Core         TipoGrupoMuscular = "Core"
	Ninguno      TipoGrupoMuscular = ""
)

type GrupoMuscular struct {
	Tipo     TipoGrupoMuscular // 3 tipos
	Musculos []Musculo         // grupo muscular compuesto de varios musculos
}