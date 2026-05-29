package db

import (
	"game-db/game"
	"game-db/player"
)

// Repräsentiert eine Spiele-Datenbank, die Spieler und Spiele verwaltet.
type GameDb struct {
	// Players enthält alle Spieler in der Datenbank.
	Players []*player.Player

	// Games enthält alle Spiele in der Datenbank.
	Games []*game.Game

	// MinHoursForRecommendation definiert die Mindestanzahl gespielter Stunden,
	// um ein Spiel für Empfehlungen zu berücksichtigen.
	MinHoursForRecommendation int

	// MinPlayersForRecommendation definiert die Mindestanzahl von Spielern,
	// die ein Spiel gespielt haben müssen, damit es für Empfehlungen berücksichtigt wird.
	MinPlayersForRecommendation int

	// Neu: Indizes
	gamesByGenre    map[string][]*game.Game
	qualifiedByGame map[string]int // Titel → Anzahl qualifizierter Spieler
}

// New erstellt eine neue leere Datenbank.
func New(minHours, minPlayers int) *GameDb {
	return &GameDb{
		Players:                     []*player.Player{},
		Games:                       []*game.Game{},
		MinHoursForRecommendation:   minHours,
		MinPlayersForRecommendation: minPlayers,
	}
}

// AddPlayer fügt einen Spieler zur Datenbank hinzu.
func (db *GameDb) AddPlayer(p *player.Player) {
	db.Players = append(db.Players, p)
}

// AddGame fügt ein Spiel zur Datenbank hinzu.
func (db *GameDb) AddGame(g *game.Game) {
	db.Games = append(db.Games, g)
}
