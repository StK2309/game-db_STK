package db

import (
	"game-db/game"
)

// RecommendGames erwartet einen Spielernamen und generiert Spiele-Empfehlungen.
//
// Die Funktion sucht nach Spielen, die der Spieler oft gespielt hat.
// Für diese Spiele werden Spiele mit gleichem Genre gesucht,
// die von anderen Spielern häufig gespielt wurden.
func (db *GameDb) RecommendGames(playerName string) []*game.Game {
	if !db.PlayerExists(playerName) {
		return nil
	}

	// Schritt 1: Finde die Spiele, die der Spieler oft gespielt hat.
	favoriteGames := db.GetFavoriteGames(playerName)

	// Schritt 2: Sammle Genres der Lieblingsspiele.
	genres := make(map[string]bool)
	for _, game := range favoriteGames {
		genres[game.Genre] = true
	}

	// Schritt 3: Suche nach Spielen in den gleichen Genres, die von anderen Spielern häufig gespielt wurden.
	recommendedGames := make(map[string]*game.Game)
	for genre := range genres {
		popularGames := db.GetPopularGamesByGenre(genre)
		for _, popularGame := range popularGames {
			if !db.HasPlayerPlayedGame(playerName, popularGame.Title) {
				recommendedGames[popularGame.Title] = popularGame
			}
		}
	}

	// Schritt 4: Konvertiere die Map in eine Liste und gib sie zurück.
	result := make([]*game.Game, 0, len(recommendedGames))
	for _, game := range recommendedGames {
		result = append(result, game)
	}

	return result
}
