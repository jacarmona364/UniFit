package models

import "errors"

var (
	ErrNombreVacio        = errors.New("el nombre del ejercicio no puede estar vacío")
	ErrDificultadInvalida = errors.New("el nivel de dificultad proporcionado no es reconocido")
	ErrGrupoSinEjercicios = errors.New("no se encontraron ejercicios para el grupo muscular solicitado")
)
