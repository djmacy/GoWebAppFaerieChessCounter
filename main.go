package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type PieceWorth int

const (
	PAWN PieceWorth = iota + 1
	PEASANT
	SOLDIER
	ROOK
	KNIGHT
	BISHOP
	CATAPULT
	CHAMBERLAIN
	COURTESAN
	HERALD
	INQUISITOR
	LANCER
	PONTIFF
	THIEF
	TOWER
	QUEEN
	JESTER
	KING
	REGENT
)

var pieceValues = map[PieceWorth]int{
	PAWN:        1,
	PEASANT:     2,
	SOLDIER:     3,
	ROOK:        9,
	KNIGHT:      4,
	BISHOP:      6,
	CATAPULT:    3,
	CHAMBERLAIN: 6,
	COURTESAN:   6,
	HERALD:      6,
	INQUISITOR:  8,
	LANCER:      5,
	PONTIFF:     8,
	THIEF:       5,
	TOWER:       10,
	QUEEN:       12,
	JESTER:      12,
	KING:        0,
	REGENT:      15,
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

var homeTmpl, resultTmpl *template.Template

func init() {
	homeTmpl = template.Must(template.ParseFiles("templates/index.html"))
	resultTmpl = template.Must(template.ParseFiles("templates/result.html"))
}

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

	pawn := parseInt(r.FormValue("pawn"))
	peasant := parseInt(r.FormValue("peasant"))
	soldier := parseInt(r.FormValue("soldier"))
	rook := parseInt(r.FormValue("rook"))
	knight := parseInt(r.FormValue("knight"))
	bishop := parseInt(r.FormValue("bishop"))
	catapult := parseInt(r.FormValue("catapult"))
	chamberlain := parseInt(r.FormValue("chamberlain"))
	courtesan := parseInt(r.FormValue("courtesan"))
	herald := parseInt(r.FormValue("herald"))
	inquisitor := parseInt(r.FormValue("inquisitor"))
	lancer := parseInt(r.FormValue("lancer"))
	pontiff := parseInt(r.FormValue("pontiff"))
	thief := parseInt(r.FormValue("thief"))
	tower := parseInt(r.FormValue("tower"))
	queen := r.FormValue("queen")
	king := r.FormValue("king")
	difficulty := r.FormValue("difficulty")

	var kingOrRegentValue int
	if king == "King" {
		kingOrRegentValue = 0
	} else {
		kingOrRegentValue = 15
	}
	var queenOrJesterValue int
	if queen == "Queen" {
		queenOrJesterValue = 12
	} else {
		queenOrJesterValue = 12
	}

	totalPoints :=
		kingOrRegentValue + queenOrJesterValue +
			getPieceValue(pawn, PAWN) +
			getPieceValue(peasant, PEASANT) +
			getPieceValue(soldier, SOLDIER) +
			getPieceValue(rook, ROOK) +
			getPieceValue(knight, KNIGHT) +
			getPieceValue(bishop, BISHOP) +
			getPieceValue(catapult, CATAPULT) +
			getPieceValue(chamberlain, CHAMBERLAIN) +
			getPieceValue(courtesan, COURTESAN) +
			getPieceValue(herald, HERALD) +
			getPieceValue(inquisitor, INQUISITOR) +
			getPieceValue(lancer, LANCER) +
			getPieceValue(pontiff, PONTIFF) +
			getPieceValue(thief, THIEF) +
			getPieceValue(tower, TOWER)

	remainingPoints := difficultyPoints.getDifficultyPoints(difficulty) - totalPoints

	data := struct {
		TotalPoints     int
		RemainingPoints int
	}{
		TotalPoints:     totalPoints,
		RemainingPoints: remainingPoints,
	}

	err = resultTmpl.Execute(w, data)
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

func getPieceValue(count int, pieceType PieceWorth) int {
	return count * pieceValues[pieceType]
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
