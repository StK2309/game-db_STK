package db

import (
	"game-db/game"
	"game-db/player"
)

// PlayerExists prüft, ob ein Spieler mit dem gegebenen Namen existiert.
func (db *GameDb) PlayerExists(name string) bool {
	return db.GetPlayer(name) != nil
}

// GetFavoriteGames liefert alle Spiele, die ein Spieler mindestens so lange
// gespielt hat, wie es für Empfehlungen erforderlich ist.
func (db *GameDb) GetFavoriteGames(name string) []*game.Game {
	return db.GetPlayedGames(name, db.MinHoursForRecommendation)
}

// GetPlayer sucht einen Spieler in der Datenbank anhand eines Namens.
// Liefert einen Zeiger auf den Spieler zurück, wenn er gefunden wird,
// oder nil, wenn er nicht gefunden wird.
func (db *GameDb) GetPlayer(name string) *player.Player {
	for _, p := range db.Players {
		if p.Name == name {
			return p
		}
	}
	return nil
}

// GetPlayersByGame sucht Spieler in der Datenbank, die ein bestimmtes Spiel gespielt haben.
// Erwartet den Titel des Spiels und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayersByGame(title string, min_played int) []*player.Player {
	players := []*player.Player{}
	for _, p := range db.Players {
		if p.HasPlayedMore(&game.Game{Title: title}, min_played) {
			players = append(players, p)
		}
	}
	return players
}
