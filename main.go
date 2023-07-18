package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type ChessFormData struct {
	pawn        int
	peasant     int
	soldier     int
	rook        int
	knight      int
	bishop      int
	catapult    int
	chamberlain int
	courtesan   int
	herald      int
	inquisitor  int
	lancer      int
	pontiff     int
	thief       int
	tower       int
	queen       int
	king        int
	jester      int
	regent      int
	difficulty  string
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	formData := ChessFormData{
		pawn:        parseInt(r.FormValue("pawn")),
		peasant:     parseInt(r.FormValue("peasant")),
		soldier:     parseInt(r.FormValue("soldier")),
		rook:        parseInt(r.FormValue("rook")),
		knight:      parseInt(r.FormValue("knight")),
		bishop:      parseInt(r.FormValue("bishop")),
		catapult:    parseInt(r.FormValue("catapult")),
		chamberlain: parseInt(r.FormValue("chamberlain")),
		courtesan:   parseInt(r.FormValue("courtesan")),
		herald:      parseInt(r.FormValue("herald")),
		inquisitor:  parseInt(r.FormValue("inquisitor")),
		lancer:      parseInt(r.FormValue("lancer")),
		pontiff:     parseInt(r.FormValue("pontiff")),
		thief:       parseInt(r.FormValue("thief")),
		tower:       parseInt(r.FormValue("tower")),
		queen:       parseInt(r.FormValue("queen")),
		king:        parseInt(r.FormValue("king")),
		jester:      parseInt(r.FormValue("jester")),
		regent:      parseInt(r.FormValue("regent")),
		difficulty:  r.FormValue("difficulty"),
	}

	totalPoints := formData.pawn*1 +
		formData.peasant*2 +
		formData.soldier*3 +
		formData.rook*9 +
		formData.knight*4 +
		formData.bishop*6 +
		formData.catapult*3 +
		formData.chamberlain*6 +
		formData.courtesan*6 +
		formData.herald*6 +
		formData.inquisitor*8 +
		formData.lancer*5 +
		formData.pontiff*8 +
		formData.thief*5 +
		formData.tower*10 +
		formData.queen*12 +
		formData.king*0 +
		formData.jester*12 +
		formData.regent*15

	difficulties := map[string]int{
		"Beginner":     65,
		"Intermediate": 70,
		"Advanced":     75,
	}

	remainingPoints := difficulties[formData.difficulty] - totalPoints

	data := map[string]interface{}{
		"TotalPoints":     totalPoints,
		"RemainingPoints": remainingPoints,
	}

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		return
	}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/calculate", calculateHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
