
package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//GOOD PATH

func TestRutina_GenerarSemana_AsignacionCorrecta(t *testing.T) {
  
	rutina := NuevaRutina()
	
	ejPecho, _ := NuevoEjercicio("Press Banca", Intermediate)
	ejPierna, _ := NuevoEjercicio("Sentadilla", Advanced)


	semana, err := rutina.GenerarSemana()

	assert.NoError(t, err)
	assert.NotNil(t, semana)

	assert.NotEmpty(t, semana["Lunes"], "El lunes debería tener ejercicios")
	assert.Equal(t, 1, len(semana["Lunes"]))
	
		assert.NotEmpty(t, semana["Miércoles"], "El miércoles debería tener ejercicios")
	
	assert.Empty(t, semana["Martes"], "El martes no añadimos nada, debería estar vacío")
}

//SAD PATH

func TestRutina_GenerarSemana_ErrorSiVacia(t *testing.T) {
	rutina := NuevaRutina()

	plan, err := rutina.GenerarSemana()

	assert.Error(t, err)
	assert.Nil(t, plan)
  
	assert.Is(t, err, ErrRutinaVacia, "Debe avisar si intentamos generar rutina sin ejercicios")
}
