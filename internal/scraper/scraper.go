package scraper

import (
	"io"
	"regexp"
	"strings"

	"UniFit/internal/errors"
	"UniFit/internal/models"
)

var reTags = regexp.MustCompile(`<[^>]+>`)

func ExtraerRutinas(htmlContent string) (map[models.GrupoMuscular]string, error) {
	
	bloqueTrabajo, err := obtenerBloquePrincipal(htmlContent)
	if err != nil {
		return nil, err
	}

	return procesarListaRutinas(bloqueTrabajo)
}

func obtenerBloquePrincipal(html string) (string, error) {
	const marcador = `class="mainpage-category-list"`
	idx := strings.Index(html, marcador)
	if idx == -1 {
		return "", errors.ErrNoListaCategorias
	}
	return html[idx:], nil
}

func procesarListaRutinas(bloque string) (map[models.GrupoMuscular]string, error) {
	rutinasMap := make(map[models.GrupoMuscular]string)
	celdas := strings.Split(bloque, `class="cell"`)
	encontrado := false

	for i := 1; i < len(celdas); i++ {
		celda := celdas[i]
		if !strings.Contains(celda, `class="category-name"`) {
			continue
		}

		nombre, url, err := extraerDatosRutina(celda)
		if err == nil && nombre != "" && url != "" {
			grupo := models.GrupoMuscular(nombre)
			rutinasMap[grupo] = url
			encontrado = true
		}
	}

	if !encontrado {
		return nil, errors.ErrNoClaseNombreCategoria
	}
	return rutinasMap, nil
}

func parsearDificultad(nivelHTML string) models.Dificultad {

	limpio := strings.ToUpper(strings.TrimSpace(nivelHTML))

	switch limpio {
	case "BEGINNER":
		return models.Beginner
	case "INTERMEDIATE":
		return models.Intermediate
	case "ADVANCED":
		return models.Advanced
	default:
		return models.Beginner 
	}
}

func ExtraerEjercicios(htmlContent string) ([]models.Ejercicio, error) {
	
	bloqueLista, err := obtenerBloqueEjercicios(htmlContent)
	if err != nil {
		return nil, err
	}

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

	for i := 1; i < len(celdas); i++ {
		if ej, err := construirEjercicio(celdas[i]); err == nil {
			ejercicios = append(ejercicios, ej)
		}
	}

	if len(ejercicios) == 0 && len(celdas) > 1 {
		return nil, diagnosticarErrorEstructura(celdas[1])
	}
	if len(ejercicios) == 0 {
		return nil, nil
	}

	return ejercicios, nil
}

func construirEjercicio(celda string) (models.Ejercicio, error) {
	if !strings.Contains(celda, "node-title") || !strings.Contains(celda, "Exp. Level") {
		return models.Ejercicio{}, errors.ErrLecturaHTML
	}

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


func leerHTML(r io.Reader) (string, error) {
	bodyBytes, err := io.ReadAll(r)
	if err != nil {
		return "", errors.ErrLecturaHTML
	}
	return string(bodyBytes), nil
}

func extraerDatosRutina(bloque string) (nombre, url string, err error) {
	startHref := strings.Index(bloque, `href="`)
	if startHref != -1 {
		startHref += 6 // Longitud de href="
		endHref := strings.Index(bloque[startHref:], `"`)
		if endHref != -1 {
			url = bloque[startHref : startHref+endHref]
		}
	}

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
	idxTitle := strings.Index(bloque, `class="node-title"`)
	if idxTitle == -1 {
		return "", errors.ErrNoClaseTituloNodo
	}
	
	fragmento := bloque[idxTitle:]
	idxFin := strings.Index(fragmento, "</div>")
	
	if idxFin != -1 {
		fragmento = fragmento[:idxFin]
	}
	
	return LimpiarTexto(match[1]), nil
}

func extraerNivelExperiencia(bloque string) (string, error) {
	marcador := "Exp. Level</label>"
	idx := strings.Index(bloque, marcador)
	if idx == -1 {
		return "", errors.ErrNoEtiquetaNivel
	}

	resto := bloque[idx+len(marcador):]
	idxFin := strings.Index(resto, "</div>")
	if idxFin == -1 {
		return "", nil
	}

	return LimpiarTexto(resto[:idxFin]), nil
}

func LimpiarTexto(textoSucio string) string {
	s := reTags.ReplaceAllString(textoSucio, "")
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.TrimSpace(s)
	return s
}
