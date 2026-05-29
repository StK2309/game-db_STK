package db

import "game-db/game"

// GetPopularGamesByGenre liefert alle Spiele eines Genres, die von mindestens
// der erforderlichen Anzahl an Spielern lange genug gespielt wurden.
func (db *GameDb) GetPopularGamesByGenre(genre string) []*game.Game {
	popularGames := []*game.Game{}
	for _, g := range db.GetGamesByGenre(genre) {
		if len(db.GetPlayersByGame(g.Title, db.MinHoursForRecommendation)) >= db.MinPlayersForRecommendation {
			popularGames = append(popularGames, g)
		}
	}
	return popularGames
}

// GetPlayedGames sucht alle Spiele, die ein gegebener Spieler gespielt hat.
// Erwartet dabei den Namen des Spielers und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayedGames(name string, min_played int) []*game.Game {
	games := []*game.Game{}
	if player := db.GetPlayer(name); player != nil {
		for _, g := range db.Games {
			if player.HasPlayedMore(g, min_played) {
				games = append(games, g)
			}
		}
	}
	return games
}

// HasPlayerPlayedGame prüft, ob ein Spieler ein bestimmtes Spiel mindestens so
// lange gespielt hat, wie es für Empfehlungen erforderlich ist.
func (db *GameDb) HasPlayerPlayedGame(playerName, gameTitle string) bool {
	player := db.GetPlayer(playerName)
	if player == nil {
		return false
	}
	for _, g := range db.Games {
		if g.Title == gameTitle {
			return player.HasPlayedMore(g, db.MinHoursForRecommendation)
		}
	}
	return false
}

// GetGamesByGenre sucht Spiele in der Datenbank anhand ihres Genres.
func (db *GameDb) GetGamesByGenre(genre string) []*game.Game {
	games := []*game.Game{}
	for _, g := range db.Games {
		if g.Genre == genre {
			games = append(games, g)
		}
	}
	return games
}
