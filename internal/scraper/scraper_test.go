package scraper

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	internalErrors "UniFit/internal/errors"
	"UniFit/internal/models"
)

// Ruta relativa a la carpeta donde están los HTML
const testDataDir = "testdata"

// TESTS PARA EXTRAER RUTINAS (HAPPY PATH)

func TestExtraerRutinas(t *testing.T) {
	fileName := "index_rutina.html"
	filePath := filepath.Join(testDataDir, fileName)

	t.Run("Happy_Path_File_"+fileName, func(t *testing.T) {
		// 1. Abrir fichero
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatalf("No se pudo leer el archivo %s: %v", fileName, err)
		}
		defer file.Close()

		// 2. Ejecutar
		rutinasMap, err := ExtraerRutinas(file)

		// 3. Validaciones
		if err != nil {
			t.Errorf("No se esperaba error y se obtuvo: %v", err)
			return
		}

		if len(rutinasMap) == 0 {
			t.Fatal("El mapa de rutinas está vacío")
		}

		// 4. Integridad de datos (Los grupos existen en el HTML)
		gruposEsperados := []models.GrupoMuscular{models.Chest, models.Abs}
		for _, grupo := range gruposEsperados {
			if _, exists := rutinasMap[grupo]; !exists {
				t.Errorf("Falta el grupo muscular '%s' en el mapa extraído", grupo)
			}
		}
	})
}


// TESTS PARA EXTRAER EJERCICIOS (HAPPY PATH)

func TestExtraerEjercicios(t *testing.T) {
	fileName := "detalle_abs.html"
	filePath := filepath.Join(testDataDir, fileName)

	t.Run("Happy_Path_File_"+fileName, func(t *testing.T) {
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatalf("No se pudo leer el archivo %s: %v", fileName, err)
		}
		defer file.Close()

		ejercicios, err := ExtraerEjercicios(file)

		if err != nil {
			t.Errorf("No se esperaba error y se obtuvo: %v", err)
			return
		}

		if len(ejercicios) == 0 {
			t.Fatal("La lista de ejercicios está vacía")
		}

		// 4. Validación específica del primer ejercicio del HTML
		ej := ejercicios[0]
		nombreEsperado := "Lying Floor Leg Raise"
		
		if ej.Nombre != nombreEsperado {
			t.Errorf("Nombre incorrecto. Esperado: '%s', Obtenido: '%s'", nombreEsperado, ej.Nombre)
		}
		if ej.Dificultad != models.Beginner {
			t.Errorf("Dificultad incorrecta. Esperado: '%s', Obtenido: '%s'", models.Beginner, ej.Dificultad)
		}
	})
}


// TESTS ESTRUCTURALES RUTINAS (SAD PATH)

func TestExtraerRutinas_Estructura(t *testing.T) {
	
	t.Run("Sad_Path_Sin_Lista_Categorias", func(t *testing.T) {
		// HTML sin la clase principal
		html := `<html><body><div class="otra-cosa">Nada</div></body></html>`
		
		_, err := ExtraerRutinas(strings.NewReader(html))

		if !errors.Is(err, internalErrors.ErrNoListaCategorias) {
			t.Errorf("Se esperaba ErrNoListaCategorias, se obtuvo: %v", err)
		}
	})

	t.Run("Sad_Path_Sin_Clase_Nombre", func(t *testing.T) {
		// Existe el contenedor padre, pero cambia el nombre de la clase hija
		html := `
		<div class="mainpage-category-list">
			<a href="/link">
				<div class="nombre-roto">Pecho</div> 
			</a>
		</div>`
		
		_, err := ExtraerRutinas(strings.NewReader(html))

		if !errors.Is(err, internalErrors.ErrNoClaseNombreCategoria) {
			t.Errorf("Se esperaba ErrNoClaseNombreCategoria, se obtuvo: %v", err)
		}
	})
}


// TESTS ESTRUCTURALES EJERCICIOS (SAD PATH)

func TestExtraerEjercicios_Estructura(t *testing.T) {

	t.Run("Sad_Path_Sin_Contenedor", func(t *testing.T) {
		// Falta el contenedor "view-exercise-term-list"
		html := `<html><body><h1>Hola</h1></body></html>`
		
		_, err := ExtraerEjercicios(strings.NewReader(html))

		if !errors.Is(err, internalErrors.ErrNoContenedorEjercicios) {
			t.Errorf("Se esperaba ErrNoContenedorEjercicios, se obtuvo: %v", err)
		}
	})

	t.Run("Sad_Path_Sin_Titulo_Nodo", func(t *testing.T) {
		// Falta la clase "node-title"
		html := `
		<div class="view-exercise-term-list">
			<div class="cell">
				<div class="titulo-roto"><a href="#">Ejercicio</a></div>
				<label>Exp. Level</label> Beginner
			</div>
		</div>`
		
		_, err := ExtraerEjercicios(strings.NewReader(html))

		if !errors.Is(err, internalErrors.ErrNoClaseTituloNodo) {
			t.Errorf("Se esperaba ErrNoClaseTituloNodo, se obtuvo: %v", err)
		}
	})

	t.Run("Sad_Path_Sin_Etiqueta_Nivel", func(t *testing.T) {
		// Falta la etiqueta "Exp. Level"
		html := `
		<div class="view-exercise-term-list">
			<div class="node-title"><a href="#">Ejercicio</a></div>
			<div class="meta-box"><label>Level</label> Beginner</div>
		</div>`
		
		_, err := ExtraerEjercicios(strings.NewReader(html))

		if !errors.Is(err, internalErrors.ErrNoEtiquetaNivel) {
			t.Errorf("Se esperaba ErrNoEtiquetaNivel, se obtuvo: %v", err)
		}
	})
}
