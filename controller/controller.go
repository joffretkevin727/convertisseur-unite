package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"proj1/structure"
	"strconv"
	"text/template"
)

// Liste globale de toutes les unités supportées
var allUnits = []string{"mm", "cm", "m", "km", "mi", "nm", "g", "kg", "lb", "ml", "l", "°C", "°F"}

func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	// Analyse et rend le template spécifié
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Affiche la page initiale avec toutes les unités
	data := structure.PageData{
		Units: allUnits,
	}
	RenderTemplate(w, "home.html", data)
}

func Convertir(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	valStr := r.FormValue("valeur")
	source := r.FormValue("source")
	cible := r.FormValue("cible")
	valeur, _ := strconv.ParseFloat(valStr, 64)

	// Définition des groupes pour la validation
	unitesDistance := map[string]bool{"mm": true, "cm": true, "m": true, "km": true, "mi": true, "nm": true}
	unitesMasse := map[string]bool{"g": true, "kg": true, "lb": true}
	unitesVolume := map[string]bool{"ml": true, "l": true}

	// Vérification de compatibilité
	compatible := (unitesDistance[source] && unitesDistance[cible]) ||
		(unitesMasse[source] && unitesMasse[cible]) ||
		(unitesVolume[source] && unitesVolume[cible]) ||
		(source == "°C" || source == "°F") && (cible == "°C" || cible == "°F")

	data := structure.PageData{Units: allUnits}

	if !compatible {
		// Message d'erreur si les catégories sont différentes
		data.Resultat = "Erreur : Impossible de convertir des " + source + " en " + cible
	} else {
		// Logique de calcul existante
		coeffs := map[string]float64{
			"mm": 0.001, "cm": 0.01, "m": 1, "km": 1000, "mi": 1609.34, "nm": 1.852,
			"g": 1, "kg": 1000, "lb": 453.59,
			"ml": 0.001, "l": 1,
		}

		var resultat float64
		if source == "°C" && cible == "°F" {
			resultat = (valeur * 9 / 5) + 32
		} else if source == "°F" && cible == "°C" {
			resultat = (valeur - 32) * 5 / 9
		} else {
			resultat = (valeur * coeffs[source]) / coeffs[cible]
		}
		data.Resultat = fmt.Sprintf("%.2f %s = %.8f %s", valeur, source, resultat, cible)
	}

	RenderTemplate(w, "home.html", data)
}
