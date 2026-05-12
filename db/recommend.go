package db

import (
	"game-db/game"
	"game-db/player"
)

// RecommendGames erwartet einen Spielernamen und generiert Spiele-Empfehlungen.
//
// Die Funktion sucht nach Spielen, die der Spieler oft gespielt hat.
// Für diese Spiele werden Spiele mit gleichem Genre gesucht,
// die von anderen Spielern häufig gespielt wurden.
func (db *GameDb) RecommendGames(playerName string) []*game.Game {
	recommendedGames := []*game.Game{}

	// Spieler finden
	var player *player.Player
	for _, p := range db.Players {
		if p.Name == playerName {
			player = p
			break
		}
	}
	if player == nil {
		return recommendedGames
	}

	// Oft gespielte Titel des Spielers
	playedTitles := player.PlayedGames(db.MinHoursForRecommendation)
	if len(playedTitles) == 0 {
		return recommendedGames
	}

	// Genres der gespielten Spiele sammeln
	genres := map[string]bool{}
	playerGames := map[string]bool{}
	for _, title := range playedTitles {
		playerGames[title] = true
		for _, g := range db.Games {
			if g.Title == title {
				genres[g.Genre] = true
			}
		}
	}

	// Spiele anderer Spieler mit gleichem Genre zählen
	candidateCount := map[string]int{}
	for _, other := range db.Players {
		if other.Name == playerName {
			continue
		}
		for _, title := range other.PlayedGames(db.MinHoursForRecommendation) {
			if !playerGames[title] {
				candidateCount[title]++
			}
		}
	}

	// Nur Spiele empfehlen, die genug Spieler gespielt haben und passendes Genre haben
	seen := map[string]bool{}
	for title, count := range candidateCount {
		if count < db.MinPlayersForRecommendation || seen[title] {
			continue
		}
		for _, g := range db.Games {
			if g.Title == title && genres[g.Genre] {
				recommendedGames = append(recommendedGames, g)
				seen[title] = true
			}
		}
	}

	return recommendedGames
}
