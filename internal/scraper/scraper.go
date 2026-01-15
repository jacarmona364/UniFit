package scraper

import (
	"io"
	"regexp"
	"strings"

	"UniFit/internal/errors"
	"UniFit/internal/models"
)

func ExtraerRutinas(htmlContent string) (map[models.GrupoMuscular]string, error) {
	
	bloqueTrabajo, err := obtenerBloquePrincipal(htmlContent)
	if err != nil {
		return nil, err
	}

	return procesarListaRutinas(bloqueTrabajo)
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
