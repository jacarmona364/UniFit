package scraper

import (
	"io"
	"regexp"
	"strings"

	"UniFit/internal/errors"
	"UniFit/internal/models"
)

var reTags = regexp.MustCompile(`<[^>]+>`)

func ExtraerRutinas(r io.Reader) (map[models.GrupoMuscular]string, error) {
	htmlContent, err := leerHTML(r)
	if err != nil {
		return nil, err
	}

	bloqueTrabajo, err := obtenerBloquePrincipal(htmlContent)
	if err != nil {
		return nil, err
	}

	return procesarListaRutinas(bloqueTrabajo)
}

func ParseDificultad(nivelHTML string) models.Dificultad {
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

func ExtraerEjercicios(r io.Reader) ([]models.Ejercicio, error) {
	htmlContent, err := leerHTML(r)
	if err != nil {
		return nil, err
	}
	
	bloqueLista, err := obtenerBloqueEjercicios(htmlContent)
	if err != nil {
		return nil, err
	}

	return procesarListaEjercicios(bloqueLista)
}

func diagnosticarErrorEstructura(celdaMuestra string) error {
	if !strings.Contains(celdaMuestra, "node-title") {
		return errors.ErrNoClaseTituloNodo
	}

	nivel, err := extraerNivelExperiencia(celdaMuestra)
	if err != nil || nivel == "" {
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

func extraerNivelExperiencia(bloque string) (string, error) {
	marcador := "Exp. Level</label>"
	idx := strings.Index(bloque, marcador)
	if idx == -1 {
		return "", errors.ErrNoEtiquetaNivel
	}

	resto := bloque[idx+len(marcador):]
	idxFin := strings.Index(resto, "</div>")
	
	if idxFin == -1 {
		return "", errors.ErrLecturaHTML
	}

	return LimpiarTexto(resto[:idxFin]), nil
}


// -------------------------------------------------------------------------
// LÓGICA INTEGRADA DESDE PARSER.GO
// -------------------------------------------------------------------------

func obtenerBloquePrincipal(html string) (string, error) {
	const marcador = "mainpage-category-list"
	idx := strings.Index(html, marcador)
	if idx == -1 {
		return "", errors.ErrNoListaCategorias
	}
	return html[idx:], nil
}

func procesarListaRutinas(bloque string) (map[models.GrupoMuscular]string, error) {
	rutinasMap := make(map[models.GrupoMuscular]string)
	celdas := strings.Split(bloque, "cell") 
	encontrado := false

	for i := 1; i < len(celdas); i++ {
		celda := celdas[i]
		if !strings.Contains(celda, "category-name") {
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

func extraerDatosRutina(bloque string) (nombre, url string, err error) {
	const long = 6
	startHref := strings.Index(bloque, `href="`)
	if startHref != -1 {
		startHref += long
		endHref := strings.Index(bloque[startHref:], `"`)
		if endHref != -1 {
			url = bloque[startHref : startHref+endHref]
		}
	}

	idxCat := strings.Index(bloque, "category-name")
	if idxCat != -1 {
		sub := bloque[idxCat:]
		idxCierre := strings.Index(sub, ">")
		if idxCierre != -1 {
			startName := idxCat + idxCierre + 1
			endName := strings.Index(bloque[startName:], "</div>")
			if endName != -1 {
				nombre = LimpiarTexto(bloque[startName : startName+endName])
			}
		}
	}
	return nombre, url, nil
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
	if len(celdas) <= 1 {
		celdas = strings.Split(bloque, "cell")
	}

	var ejercicios []models.Ejercicio

	for i := 1; i < len(celdas); i++ {
		if !strings.Contains(celdas[i], "node-title") {
			continue
		}
		
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
		Dificultad: ParseDificultad(nivelStr),
	}, nil
}

func extraerNombreEjercicio(bloque string) (string, error) {
	idxTitle := strings.Index(bloque, "node-title")
	if idxTitle == -1 {
		return "", errors.ErrNoClaseTituloNodo
	}
	
	fragmento := bloque[idxTitle:]
	
	idxFinLink := strings.Index(fragmento, "</a>")
	if idxFinLink == -1 {
		 return "", errors.ErrLecturaHTML
	}
	
	subFragmento := fragmento[:idxFinLink]
	idxInicioTexto := strings.LastIndex(subFragmento, ">")
	
	if idxInicioTexto == -1 {
		return "", errors.ErrLecturaHTML
	}

	return LimpiarTexto(subFragmento[idxInicioTexto+1:]), nil
}

func LimpiarTexto(textoSucio string) string {
	s := reTags.ReplaceAllString(textoSucio, "")
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.TrimSpace(s)
	return s
}
