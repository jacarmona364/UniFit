package errors

import "errors"

var (
	
	// Errores Generales
	ErrReadHTML = errors.New("error al leer el flujo de datos HTML")

	
	// Errores de la Página Principal
	// Basado en: <div class="mainpage-category-list">
	ErrNoMainPageList = errors.New("no se encontró el contenedor principal '.mainpage-category-list'")
	
	// Basado en: <div class="category-name">
	ErrNoCategoryNameClass = errors.New("no se encontró la clase '.category-name' para el nombre del grupo")

	
	// Errores de la Página de Detalle
	// Basado en: <div class="view-exercise-term-list ...">
	ErrNoExerciseListContainer = errors.New("no se encontró el contenedor de lista '.view-exercise-term-list'")

	// Basado en: <div class="node-title">
	ErrNoNodeTitleClass = errors.New("no se encontró la clase '.node-title' para el título del ejercicio")

	// Basado en: <label>Exp. Level</label>
	ErrNoExpLevelLabel = errors.New("no se encontró la etiqueta '<label>Exp. Level</label>'")
	
)
