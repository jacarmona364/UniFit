ackage models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//GOOD PATH (El camino feliz) 

func TestNuevoEjercicio_CreacionCorrecta(t *testing.T) {
	ej, err := NuevoEjercicio("Press Banca", Beginner)

	assert.NoError(t, err, "No debería dar error con datos válidos")
	assert.NotNil(t, ej, "El objeto ejercicio no debería ser nulo")
	
}

//SAD PATH (Gestión de errores)

func TestNuevoEjercicio_ErrorSiNombreVacio(t *testing.T) {
	_, err := NuevoEjercicio("", Advanced)

	assert.Error(t, err, "Debería lanzar error")
	assert.Is(t, err, ErrNombreVacio, "Debe devolver ErrNombreVacio")
}

func TestNuevoEjercicio_ErrorSiDificultadInvalida(t *testing.T) {
	_, err := NuevoEjercicio("Sentadilla", Dificultad("ADVANCED"))

	assert.Error(t, err)
	assert.Is(t, err, ErrDificultadInvalida, "Debe validar que la dificultad exista")
}
