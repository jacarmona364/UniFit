package scraper

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalErrors "UniFit/internal/errors"
	"UniFit/internal/models"
)

// Constantes de prueba (Fuera de los tests para fácil configuración)
const (
	testDataDir     = "testdata"
	fileIndexRutina = "index_rutina.html"
	fileDetalleAbs  = "detalle_abs.html"
)

// -------------------------------------------------------------------------
// HELPER (Infraestructura del test)
// -------------------------------------------------------------------------

// cargarHTML se encarga de abrir el archivo.
// Si falla, el test muere aquí y no ensucia la lógica del test principal.
func cargarHTML(t *testing.T, filename string) io.Reader {
	t.Helper() // Marca esta función como helper en los logs de error
	
	path := filepath.Join(testDataDir, filename)
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("SETUP ERROR: No se pudo cargar el fixture '%s': %v", filename, err)
	}
	
	// Registramos el cierre automático al terminar el test
	t.Cleanup(func() { file.Close() })
	
	return file
}

// -------------------------------------------------------------------------
// TESTS HAPPY PATH (Usando Testify)
// -------------------------------------------------------------------------

func TestExtraerRutinas(t *testing.T) {
	reader := cargarHTML(t, fileIndexRutina)
	rutinasMap, err := ExtraerRutinas(reader)

	require.NoError(t, err, "La extracción de rutinas no debería fallar")
	require.NotEmpty(t, rutinasMap, "El mapa de rutinas no debería estar vacío")

	assert.Contains(t, rutinasMap, models.Chest, "Debería contener Chest")
	assert.Contains(t, rutinasMap, models.Abs, "Debería contener Abs")
}

func TestExtraerEjercicios(t *testing.T) {
	reader := cargarHTML(t, fileDetalleAbs)

	ejercicios, err := ExtraerEjercicios(reader)

	require.NoError(t, err)
	require.NotEmpty(t, ejercicios)

	assert.Equal(t, "Lying Floor Leg Raise", ejercicios[0].Nombre)
	assert.Equal(t, models.Beginner, ejercicios[0].Dificultad)
}

// -------------------------------------------------------------------------
// TESTS SAD PATH (Validación de Errores con Testify)
// -------------------------------------------------------------------------

func TestExtraerRutinas_Errores(t *testing.T) {
	tests := []struct {
		name          string
		html          string
		expectedError error
	}{
		{
			name:          "Sin Lista Categorias",
			html:          `<html><body><div class="otra-cosa">Nada</div></body></html>`,
			expectedError: internalErrors.ErrNoListaCategorias,
		},
		{
			name:          "Sin Clase Nombre",
			html:          `<div class="mainpage-category-list"><div class="roto">...</div></div>`,
			expectedError: internalErrors.ErrNoClaseNombreCategoria,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ExtraerRutinas(strings.NewReader(tt.html))
			
			// Testify comprueba que el error sea EXACTAMENTE el esperado
			assert.ErrorIs(t, err, tt.expectedError)
		})
	}
}

func TestExtraerEjercicios_Errores(t *testing.T) {
	tests := []struct {
		name          string
		html          string
		expectedError error
	}{
		{
			name:          "Sin Contenedor",
			html:          `<html>Nada</html>`,
			expectedError: internalErrors.ErrNoContenedorEjercicios,
		},
		{
			name:          "Sin Titulo Nodo",
			html:          `<div class="view-exercise-term-list"><div class="cell">Sin Titulo</div></div>`,
			expectedError: internalErrors.ErrNoClaseTituloNodo,
		},
		{
			name:          "Sin Etiqueta Nivel",
			html:          `<div class="view-exercise-term-list"><div class="cell"><div class="node-title"></div>Sin Level</div></div>`,
			expectedError: internalErrors.ErrNoEtiquetaNivel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ExtraerEjercicios(strings.NewReader(tt.html))
			assert.ErrorIs(t, err, tt.expectedError)
		})
	}
}
