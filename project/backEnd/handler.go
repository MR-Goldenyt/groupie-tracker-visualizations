package backend

import (
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker-visualizations/project"
	"groupie-tracker-visualizations/project/backEnd/helpers"
)

// Handler routes HTTP requests to the correct templates or static files.
func Handler(w http.ResponseWriter, r *http.Request) {
	artists, _, err := helpers.LoadDataSet()
	if err != nil {
		helpers.ServeError(w, r, http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		helpers.ServeError(w, r, http.StatusMethodNotAllowed)
		return
	}

	switch r.URL.Path {
	case "/", "/home":
		helpers.RenderTemplate(w, r, "frontEnd/template/home.html", project.PageData{Artists: artists})

	// case "/artists":
	// 	helpers.RenderTemplate(w, r, "frontEnd/template/home.html", project.PageData{Artists: artists})

	case "/aboutus":
		helpers.ServeEmbeddedFile(w, r, "frontEnd/template/aboutus.html")

	case "/submit":
		handleSubmit(w, r, artists)

	default:
		helpers.ServeError(w, r, http.StatusNotFound)
	}
}

func AssetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/static/")
		data, err := project.StaticFS.ReadFile("frontEnd/static/" + path)
		if err != nil {
			helpers.ServeError(w, r, http.StatusNotFound)
			return
		}

		switch {
		case strings.HasSuffix(path, ".css"):
			w.Header().Set("Content-Type", "text/css")
		case strings.HasSuffix(path, ".svg"):
			w.Header().Set("Content-Type", "image/svg+xml")
		}

		w.Write(data)
	})
}

// handleSubmit processes the /submit route, validating the 'value' parameter and rendering the artist page.
func handleSubmit(w http.ResponseWriter, r *http.Request, artists []project.Artist) {
	if r.Method != http.MethodGet {
		helpers.ServeError(w, r, http.StatusMethodNotAllowed)
		return
	}

	valueStr := r.URL.Query().Get("value")
	idx, err := strconv.Atoi(valueStr)
	if err != nil {
		helpers.ServeError(w, r, http.StatusBadRequest)
		return
	} else if idx < 1 || idx > len(artists) {
		helpers.ServeError(w, r, http.StatusNotFound)
		return
	}

	// 🔄 Add this block to compute previous and next artist IDs with wraparound
	prevIdx := (idx-2+len(artists))%len(artists) + 1
	nextIdx := idx%len(artists) + 1

	// ✅ Include PrevID and NextID in the page data
	pageData := project.PageDataArtist{
		Artist:      artists[idx-1],
		LocDateList: helpers.BuildLocationDateList(idx - 1),
		PrevID:      prevIdx,
		NextID:      nextIdx,
	}

	helpers.RenderTemplate(w, r, "frontEnd/template/artist.html", pageData)
}
