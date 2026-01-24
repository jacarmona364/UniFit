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

func cargarHTML(t *testing.T, filename string) io.Reader {
	t.Helper()
	
	path := filepath.Join(testDataDir, filename)
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("SETUP ERROR: No se pudo cargar el fixture '%s': %v", filename, err)
	}
	
	t.Cleanup(func() { file.Close() })
	
	return file
}

func cargarHTMLString(t *testing.T, filename string) string {
	t.Helper()
	path := filepath.Join(testDataDir, filename)
	bytes, err := os.ReadFile(path) 
	if err != nil {
		t.Fatalf("SETUP ERROR: %v", err)
	}
	return string(bytes)
}

// -------------------------------------------------------------------------
// TESTS HAPPY PATH
// -------------------------------------------------------------------------

func TestExtraerRutinas(t *testing.T) {
	reader := cargarHTML(t, fileIndexRutina)

	rutinasMap, err := ExtraerRutinas(reader)

	require.NoError(t, err)
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
// TESTS SAD PATH 
// -------------------------------------------------------------------------

func TestExtraerRutinas_Errores(t *testing.T) {
	tests := []struct {
		name          string
		html          string
		expectedError error
	}{
		{
			name:          "Sad_Path_Sin_Lista_Categorias",
			html:          `<html><body><div class="otra-cosa">Nada</div></body></html>`,
			expectedError: internalErrors.ErrNoListaCategorias,
		},
		{
			name:          "Sad_Path_Sin_Clase_Nombre",
			html:          `<div class="mainpage-category-list"><div class="roto">...</div></div>`,
			expectedError: internalErrors.ErrNoClaseNombreCategoria,
		},
	}

for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ExtraerRutinas(strings.NewReader(tt.html))
			
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
			name:          "Sad_Path_Sin_Contenedor",
			html:          `<html>Nada</html>`,
			expectedError: internalErrors.ErrNoContenedorEjercicios,
		},
		{
			name:          "Sad_Path_Sin_Titulo_Nodo",
			html:          `<div class="view-exercise-term-list"><div class="cell">Sin Titulo</div></div>`,
			expectedError: internalErrors.ErrNoClaseTituloNodo,
		},
		{
			name:          "Sad_Path_Sin_Etiqueta_Nivel",
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
