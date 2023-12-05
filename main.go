package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var pieceValues2 = map[string]int{
	"PAWN":        1,
	"PEASANT":     2,
	"SOLDIER":     3,
	"ROOK":        9,
	"KNIGHT":      4,
	"BISHOP":      6,
	"CATAPULT":    3,
	"CHAMBERLAIN": 6,
	"COURTESAN":   6,
	"HERALD":      6,
	"INQUISITOR":  8,
	"LANCER":      5,
	"PONTIFF":     8,
	"THIEF":       5,
	"TOWER":       10,
	"QUEEN":       12,
	"JESTER":      12,
	"KING":        0,
	"REGENT":      15,
}

type DifficultyPoints struct {
	Beginner     int
	Intermediate int
	Advanced     int
}

var difficultyPoints = DifficultyPoints{
	Beginner:     65,
	Intermediate: 70,
	Advanced:     75,
}

// Used for interacting with the html files so we can dynamically adjust the html
var homeTmpl, resultTmpl *template.Template

// runs before the main function and ensures the templates are loaded
func init() {
	homeTmpl = template.Must(template.ParseFiles("templates/index.html"))
	resultTmpl = template.Must(template.ParseFiles("templates/result.html"))
}

// handles the requests to the home page. From my understanding its a controller of sorts
func homeHandler(w http.ResponseWriter, _ *http.Request) {
	err := homeTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get the number of pieces selected from the html
	pawnCount := parseInt(r.FormValue("pawn"))
	peasantCount := parseInt(r.FormValue("peasant"))
	soldierCount := parseInt(r.FormValue("soldier"))
	rookCount := parseInt(r.FormValue("rook"))
	knightCount := parseInt(r.FormValue("knight"))
	bishopCount := parseInt(r.FormValue("bishop"))
	catapultCount := parseInt(r.FormValue("catapult"))
	chamberlainCount := parseInt(r.FormValue("chamberlain"))
	courtesanCount := parseInt(r.FormValue("courtesan"))
	heraldCount := parseInt(r.FormValue("herald"))
	inquisitorCount := parseInt(r.FormValue("inquisitor"))
	lancerCount := parseInt(r.FormValue("lancer"))
	pontiffCount := parseInt(r.FormValue("pontiff"))
	thiefCount := parseInt(r.FormValue("thief"))
	towerCount := parseInt(r.FormValue("tower"))
	//these bottom two will return the actual piece name
	queen := r.FormValue("queen")
	king := r.FormValue("king")
	difficulty := r.FormValue("difficulty")

	var kingOrRegentValue int
	//if its not a king it will be regent
	if king == "King" {
		kingOrRegentValue = getPieceValue(1, "KING")
	} else {
		kingOrRegentValue = getPieceValue(1, "REGENT")
	}
	var queenOrJesterValue int
	//if its not a queen it will be a jester
	if queen == "Queen" {
		queenOrJesterValue = getPieceValue(1, "QUEEN")
	} else {
		queenOrJesterValue = getPieceValue(1, "JESTER")
	}

	totalPoints :=
		kingOrRegentValue + queenOrJesterValue +
			getPieceValue(pawnCount, "PAWN") +
			getPieceValue(peasantCount, "PEASANT") +
			getPieceValue(soldierCount, "SOLDIER") +
			getPieceValue(rookCount, "ROOK") +
			getPieceValue(knightCount, "KNIGHT") +
			getPieceValue(bishopCount, "BISHOP") +
			getPieceValue(catapultCount, "CATAPULT") +
			getPieceValue(chamberlainCount, "CHAMBERLAIN") +
			getPieceValue(courtesanCount, "COURTESAN") +
			getPieceValue(heraldCount, "HERALD") +
			getPieceValue(inquisitorCount, "INQUISITOR") +
			getPieceValue(lancerCount, "LANCER") +
			getPieceValue(pontiffCount, "PONTIFF") +
			getPieceValue(thiefCount, "THIEF") +
			getPieceValue(towerCount, "TOWER")

	remainingPoints := difficultyPoints.getDifficultyPoints(difficulty) - totalPoints

	type PointsCalculuted struct {
		TotalPoints     int
		RemainingPoints int
	}

	pointsCalculated := PointsCalculuted{
		TotalPoints:     totalPoints,
		RemainingPoints: remainingPoints,
	}

	err = resultTmpl.Execute(w, pointsCalculated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}

func getPieceValue(count int, pieceName string) int {
	return count * pieceValues2[pieceName]
}

func (d DifficultyPoints) getDifficultyPoints(difficulty string) int {
	switch difficulty {
	case "Beginner":
		return d.Beginner
	case "Intermediate":
		return d.Intermediate
	case "Advanced":
		return d.Advanced
	default:
		return 0
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/result", resultHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
