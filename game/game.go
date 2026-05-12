package game

// Repräsentiert ein Spiel in einer Spiele-Datenbank.
type Game struct {
	Title string // Der Titel des Spiels
	Genre string // Das Genre des Spiels
}

// New erstellt ein neues Spiel mit dem gegebenen Titel.
func New(title string, genre string) *Game {
	return &Game{
		Title: title,
		Genre: genre,
	}
}

// HasGenre prüft, ob das Spiel ein bestimmtes Genre hat.
func (g *Game) HasGenre(genre string) bool {
	return g.Genre == genre
}
