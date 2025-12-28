package errors

type ErrorScraper string

func (e ErrorScraper) Error() string {
	return string(e)
}

const (
	// Errores Generales
	ErrLecturaHTML = ErrorScraper("error al leer el flujo de datos HTML")
	
	// Errores Index (Rutinas)
	// <div class="mainpage-category-list">
	ErrNoListaCategorias = ErrorScraper("no se encontró el contenedor principal '.mainpage-category-list'")	
	// <div class="category-name">
	ErrNoClaseNombreCategoria = ErrorScraper("no se encontró la clase '.category-name' para el nombre del grupo")
	
	// Errores Detalle (Ejercicios)
	// <div class="view-exercise-term-list ...">
	ErrNoContenedorEjercicios = ErrorScraper("no se encontró el contenedor de lista '.view-exercise-term-list'")
	// <div class="node-title">
	ErrNoClaseTituloNodo = ErrorScraper("no se encontró la clase '.node-title' para el título del ejercicio")
	// <label>Exp. Level</label>
	ErrNoEtiquetaNivel = ErrorScraper("no se encontró la etiqueta '<label>Exp. Level</label>'"))
