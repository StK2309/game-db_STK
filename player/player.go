package player

import "game-db/game"

// Repräsentiert einen Spieler in einer Spiele-Datenbank.
type Player struct {
	Name        string
	playedGames map[string]int
}

// New erstellt einen neuen Spieler mit dem gegebenen Namen und Geburtsjahr.
func New(name string) *Player {
	return &Player{Name: name}
}

// PlayGame fügt ein Spiel zu den gespielten Spielen des Spielers hinzu und erhöht die Anzahl der gespielten Stunden.
func (p *Player) PlayGame(g *game.Game, hours int) {
	if p.playedGames == nil {
		p.playedGames = make(map[string]int)
	}
	p.playedGames[g.Title] += hours
}

// HasPlayed prüft, ob der Spieler ein bestimmtes Spiel gespielt hat.
func (p *Player) HasPlayed(g *game.Game) bool {
	if g == nil {
		return false
	}
	_, exists := p.playedGames[g.Title]
	return exists
}

// HasPlayedMore prüft, ob der Spieler ein bestimmtes Spiel mindestens `hours` Stunden gespielt hat.
func (p *Player) HasPlayedMore(g *game.Game, hours int) bool {
	if p.playedGames[g.Title] >= hours {
		return true
	}
	return false
}

// PlayedGames liefert eine Liste mit den Spielen, die der
// Spieler mehr als `hours` Stunden gespielt hat.
func (p *Player) PlayedGames(hours int) []string {
	games := []string{}
	for title, playedHours := range p.playedGames {
		if playedHours > hours {
			games = append(games, title)
		}
	}
	return games
}
