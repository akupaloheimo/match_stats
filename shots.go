package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"match_stats/shotpb"
)

type MatchStats struct {
	PlayerID   string `json:"player_id"`
	TotalShots int    `json:"total_shots"`
	ShotsIn    int    `json:"shots_in"`
	ShotsOut   int    `json:"shots_out"`
	ShotsNet   int    `json:"shots_net"`
	TotalRallies int  `json:"total_rallies"`
}

func main() {
	file, err := os.Open("shots.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var shotList shotpb.ShotList
	if err := protojson.Unmarshal(data, &shotList); err != nil {
		log.Fatalf("Failed to parse shots: %v", err)
	}

	stats := calculateStats(shotList.Shots)

	outputFile := "match_stats.json"
	saveStatsAsJSON(stats, outputFile)

	fmt.Printf("Match stats saved to %s\n", outputFile)
}

func calculateStats(shots []*shotpb.Shot) map[string]*MatchStats {
	stats := make(map[string]*MatchStats)
	ralliesSeen := make(map[int64]bool)

	for _, shot := range shots {
		if _, exists := stats[shot.PlayerId]; !exists {
			stats[shot.PlayerId] = &MatchStats{PlayerID: shot.PlayerId}
		}

		playerStats := stats[shot.PlayerId]
		playerStats.TotalShots++

		switch shot.ShotOutcome {
		case shotpb.Outcome_IN:
			playerStats.ShotsIn++
		case shotpb.Outcome_OUT:
			playerStats.ShotsOut++
		case shotpb.Outcome_NET:
			playerStats.ShotsNet++
		}

		if !ralliesSeen[shot.RallyIndex] {
			playerStats.TotalRallies++
			ralliesSeen[shot.RallyIndex] = true
		}
	}

	return stats
}

func saveStatsAsJSON(stats map[string]*MatchStats, filename string) {
	statsList := make([]*MatchStats, 0, len(stats))
	for _, stat := range stats {
		statsList = append(statsList, stat)
	}

	data, err := json.MarshalIndent(statsList, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal stats to JSON: %v", err)
	}

	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		log.Fatalf("Failed to write stats to file: %v", err)
	}
}
