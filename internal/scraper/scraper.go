package scraper

import (
	"io"
	"regexp"
	"strings"

	"UniFit/internal/errors"
	"UniFit/internal/models"
)

// Limpieza

// reLink captura el contenido dentro de un <a>: <a ...>CONTENIDO</a>
var reLink = regexp.MustCompile(`<a[^>]*>(.*?)</a>`)

// reTags elimina cualquier etiqueta HTML <...>
var reTags = regexp.MustCompile(`<[^>]*>`)


// 1. EXTRAER RUTINAS (Index)

func ExtraerRutinas(r io.Reader) (map[models.GrupoMuscular]string, error){
	// 1. Cargamos el HTML en memoria
	htmlContent, err := leerHTML(r)
	if err != nil {
		return nil, err
	}

	// 2. Acotamos la zona de trabajo (fail fast si no está el contenedor)
	bloqueTrabajo, err := obtenerBloquePrincipal(htmlContent)
	if err != nil {
		return nil, err
	}

	// 3. Procesamos las celdas encontradas dentro del bloque
	return procesarListaRutinas(bloqueTrabajo)
}

func obtenerBloquePrincipal(html string) (string, error) {
	// Buscamos el div principal. Si no está, la estructura web ha cambiado.
	const marcador = `class="mainpage-category-list"`
	idx := strings.Index(html, marcador)
	if idx == -1 {
		return "", errors.ErrNoListaCategorias
	}

	return html[idx:], nil
}

func procesarListaRutinas(bloque string) (map[models.GrupoMuscular]string, error) {
	rutinasMap := make(map[models.GrupoMuscular]string)
	
	// Separamos por celda individual
	celdas := strings.Split(bloque, `class="cell"`)
	encontrado := false

	// Empezamos en 1 porque el split 0 es basura anterior a la primera coincidencia
	for i := 1; i < len(celdas); i++ {
		celda := celdas[i]

		// Filtramos celdas que no sean de rutinas (publicidad, layout, etc)
		if !strings.Contains(celda, `class="category-name"`) {
			continue
		}

		// Delegamos la extracción fina al helper que ya tenemos
		nombre, url, err := extraerDatosRutina(celda)
		
		if err == nil && nombre != "" && url != "" {
			grupo := models.GrupoMuscular(nombre)
			rutinasMap[grupo] = url
			encontrado = true
		}
	}

	// Si iteramos todo y no sacamos nada limpio, algo falla en los selectores internos (por ejemplo: cambió la clase category-name)
	if !encontrado {
		return nil, errors.ErrNoClaseNombreCategoria
	}

	return rutinasMap, nil
}


// 2. EXTRAER EJERCICIOS (Detalle)

func ExtraerEjercicios(r io.Reader) ([]models.Ejercicio, error) {
	// 1. Carga en memoria
	htmlContent, err := leerHTML(r)
	if err != nil {
		return nil, err
	}

	// 2. Validación y recorte del bloque principal
	bloqueLista, err := obtenerBloqueEjercicios(htmlContent)
	if err != nil {
		return nil, err
	}

	// 3. Iteración y conversión de datos
	return procesarListaEjercicios(bloqueLista)
}

func obtenerBloqueEjercicios(html string) (string, error) {
	const marcador = "view-exercise-term-list"
	idx := strings.Index(html, marcador)
	if idx == -1 {
		return "", errors.ErrNoContenedorEjercicios
	}
	return html[idx:], nil
}

func procesarListaEjercicios(bloque string) ([]models.Ejercicio, error) {
	celdas := strings.Split(bloque, `class="cell`)
	var ejercicios []models.Ejercicio

	// Empezamos en 1 porque el split 0 es basura anterior a la primera celda
	for i := 1; i < len(celdas); i++ {
		// Delegamos la validación y construcción a una función pura.
		// Si falla, simplemente ignoramos este item y seguimos.
		if ej, err := construirEjercicio(celdas[i]); err == nil {
			ejercicios = append(ejercicios, ej)
		}
	}

	// Si no hemos sacado ningún ejercicio, entonces diagnosticamos POR QUÉ.
	// Esto se hace UNA sola vez, fuera del bucle.
	if len(ejercicios) == 0 && len(celdas) > 1 {
		// Usamos la primera celda real como muestra para ver qué estructura falla
		return nil, diagnosticarErrorEstructura(celdas[1])
	}

	return ejercicios, nil
}

func construirEjercicio(celda string) (models.Ejercicio, error) {
	// 1. Validaciones mínimas requeridas para que sea un ejercicio válido
	if !strings.Contains(celda, "node-title") || !strings.Contains(celda, "Exp. Level") {
		return models.Ejercicio{}, errors.ErrLecturaHTML // Error genérico para indicar fallo
	}

	// 2. Extracción de datos
	nombre, _ := extraerNombreEjercicio(celda)
	nivelStr, _ := extraerNivelExperiencia(celda)

	if nombre == "" || nivelStr == "" {
		return models.Ejercicio{}, errors.ErrLecturaHTML
	}

	return models.Ejercicio{
		Nombre:     nombre,
		Dificultad: models.ParseDificultad(nivelStr),
	}, nil
}

func diagnosticarErrorEstructura(celdaMuestra string) error {
	if !strings.Contains(celdaMuestra, "node-title") {
		return errors.ErrNoClaseTituloNodo
	}
	if !strings.Contains(celdaMuestra, "Exp. Level") {
		return errors.ErrNoEtiquetaNivel
	}
	return nil
}

// FUNCIONES AUXILIARES (HELPERS)

func leerHTML(r io.Reader) (string, error) {
	bodyBytes, err := io.ReadAll(r)
	if err != nil {
		return "", errors.ErrLecturaHTML
	}
	return string(bodyBytes), nil
}

// extraerDatosRutina saca nombre y URL de un bloque de celda de rutina
func extraerDatosRutina(bloque string) (nombre, url string, err error) {
	// Buscamos el href manualmente con strings para la URL
	// Estructura: <a href="/exercises/..."
	startHref := strings.Index(bloque, `href="`)
	if startHref != -1 {
		startHref += 6 // Longitud de href="
		endHref := strings.Index(bloque[startHref:], `"`)
		if endHref != -1 {
			url = bloque[startHref : startHref+endHref]
		}
	}

	// Para el nombre buscamos dentro de la clase category-name
	idxCat := strings.Index(bloque, `class="category-name">`)
	if idxCat != -1 {
		idxCat += len(`class="category-name">`)
		endCat := strings.Index(bloque[idxCat:], `</div>`)
		if endCat != -1 {
			rawName := bloque[idxCat : idxCat+endCat]
			nombre = LimpiarTexto(rawName)
		}
	}
	return nombre, url, nil
}

func extraerNombreEjercicio(bloque string) (string, error) {
	// Buscamos el bloque del título
	idxTitle := strings.Index(bloque, `class="node-title"`)
	if idxTitle == -1 {
		return "", errors.ErrNoClaseTituloNodo
	}
	
	fragmento := bloque[idxTitle:]
	// Usamos Regex solo para sacar el texto limpio del <a> dentro del título
	match := reLink.FindStringSubmatch(fragmento)
	if len(match) < 2 {
		return "", nil // No se encontró enlace, algo raro
	}
	
	return LimpiarTexto(match[1]), nil
}

func extraerNivelExperiencia(bloque string) (string, error) {
	// Buscamos la etiqueta Exp. Level
	marcador := "Exp. Level</label>"
	idx := strings.Index(bloque, marcador)
	if idx == -1 {
		return "", errors.ErrNoEtiquetaNivel
	}

	// El texto está después del marcador y antes del cierre del div
	resto := bloque[idx+len(marcador):]
	idxFin := strings.Index(resto, "</div>")
	if idxFin == -1 {
		return "", nil
	}

	return LimpiarTexto(resto[:idxFin]), nil
}

// Elimina tags, &nbsp;, saltos de línea y espacios extra.
func LimpiarTexto(textoSucio string) string {
	s := reTags.ReplaceAllString(textoSucio, "")
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.TrimSpace(s)
	return s
}
