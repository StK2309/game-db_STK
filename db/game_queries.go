package db

import "game-db/game"

// GetPopularGamesByGenre liefert alle Spiele eines Genres, die von mindestens
// der erforderlichen Anzahl an Spielern lange genug gespielt wurden.
func (db *GameDb) GetPopularGamesByGenre(genre string) []*game.Game {
	popularGames := []*game.Game{}
	for _, g := range db.GetGamesByGenre(genre) {
		if db.qualifiedByGame[g.Title] >= db.MinPlayersForRecommendation {
			popularGames = append(popularGames, g)
		}
	}
	return popularGames
}

// GetPlayedGames sucht alle Spiele, die ein gegebener Spieler gespielt hat.
// Erwartet dabei den Namen des Spielers und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayedGames(name string, min_played int) []*game.Game {
	games := []*game.Game{}
	player := db.GetPlayer(name)
	if player == nil {
		return games
	}

	// Build a title -> *Game index (one pass over db.Games)
	gameIndex := make(map[string]*game.Game, len(db.Games))
	for _, g := range db.Games {
		gameIndex[g.Title] = g
	}

	// Iterate only the player's played titles (usually much smaller)
	playedTitles := player.PlayedGames(min_played)
	for _, title := range playedTitles {
		if g, ok := gameIndex[title]; ok {
			games = append(games, g)
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

// GetGamesByGenreIndexed zeigt, wie man für viele Genre-Abfragen einen Index
// aufbaut: map[genre] -> []*game.Game. Bei wiederholten Aufrufen amortisiert
// sich der einmalige Aufbau.
func (db *GameDb) GetGamesByGenreIndexed(genre string) []*game.Game {
	genreIndex := make(map[string][]*game.Game)
	for _, g := range db.Games {
		genreIndex[g.Genre] = append(genreIndex[g.Genre], g)
	}
	return genreIndex[genre]
}
