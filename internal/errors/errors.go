package errors

import "errors"

var (
// Errores Generales
	ErrLecturaHTML = errors.New("error al leer el flujo de datos HTML")

	// Errores Index (Rutinas)
	// <div class="mainpage-category-list">
	ErrNoListaCategorias = errors.New("no se encontró el contenedor principal '.mainpage-category-list'")
	// <div class="category-name">
	ErrNoClaseNombreCategoria = errors.New("no se encontró la clase '.category-name' para el nombre del grupo")

	// Errores Detalle (Ejercicios)
	// <div class="view-exercise-term-list ...">
	ErrNoContenedorEjercicios = errors.New("no se encontró el contenedor de lista '.view-exercise-term-list'")
	// <div class="node-title">
	ErrNoClaseTituloNodo = errors.New("no se encontró la clase '.node-title' para el título del ejercicio")
	// <label>Exp. Level</label>
	ErrNoEtiquetaNivel = errors.New("no se encontró la etiqueta '<label>Exp. Level</label>'")
)
