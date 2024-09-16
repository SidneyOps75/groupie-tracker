package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	database "groupie-trackers/data"
	"groupie-trackers/models"
)

func HandleArtistDetail(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorPage(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	var artists []models.Artist
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		log.Printf("Error fetching artists data: %v", err)
		ErrorPage(w, "Error fetching artists data", http.StatusInternalServerError)
		return
	}

	var artist models.Artist
	for _, a := range artists {
		if a.ID == id {
			artist = a
			break
		}
	}
	if artist.ID == 0 {
		ErrorPage(w, "Artist not found", http.StatusNotFound)
		return
	}

	var locationsResponse models.LocationsResponse
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/locations", &locationsResponse); err != nil {
		log.Printf("Error fetching locations data: %v", err)
		ErrorPage(w, "Error fetching locations data", http.StatusInternalServerError)
		return
	}

	var location models.Location
	for _, loc := range locationsResponse.Index {
		if loc.ID == id {
			location = loc
			break
		}
	}

	var datesResponse models.DatesResponse
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/dates", &datesResponse); err != nil {
		log.Printf("Error fetching dates data: %v", err)
		ErrorPage(w, "Error fetching dates data", http.StatusInternalServerError)
		return
	}

	var date models.Date
	for _, d := range datesResponse.Index {
		if d.ID == id {
			date = d
			break
		}
	}

	var relationsResponse models.RelationsResponse
	if err := database.FetchData("https://groupietrackers.herokuapp.com/api/relation", &relationsResponse); err != nil {
		log.Printf("Error fetching relations data: %v", err)
		ErrorPage(w, "Error fetching relations data", http.StatusInternalServerError)
		return
	}

	var relation models.Relation
	for _, rel := range relationsResponse.Index {
		if rel.ID == id {
			relation = rel
			break
		}
	}

	data := models.ArtistDetailData{
		Artist:   artist,
		Location: location,
		Date:     date,
		Relation: relation,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorPage(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorPage(w, "Error executing template", http.StatusInternalServerError)
	}
}
