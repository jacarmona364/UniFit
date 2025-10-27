package planificacion

// --- Estructura Genérica del Deporte ---

type ActividadDeportiva struct {
	Nombre string 
	GrupoPrincipal GrupoMuscular 
	Intensidad int 
}

// --- Constantes y Objeto Fijo de Fútbol ---

const NombreActividadFutbol = "Fútbol" 
const IntensidadFutbol = 4 

// Objeto Constante de Fútbol
var ActividadDeportivaFutbol = ActividadDeportiva{
	Nombre: NombreActividadFutbol,
	Intensidad: IntensidadFutbol,
	GrupoPrincipal: GrupoMuscular{
		Tipo: TrenInferior,
	},
}