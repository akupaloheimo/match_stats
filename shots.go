package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"match_stats/shotpb"

	"google.golang.org/protobuf/encoding/protojson"
)

type MatchStats struct {
	PlayerID   		 string  `json:"player_id"`
	TotalShots 		 int     `json:"total_shots"`
	ShotsIn    		 int     `json:"shots_in"`
	ShotsFault 		 int	 `json:"shots_fault"`
	TotalRallies 	 int  	 `json:"total_rallies"`
	ServeIn 		 int 	 `json:"serve_in"`
	ServeFault 		 int 	 `json:"serve_fault"`
	ServePrecentage  float64 `json:"serve_percentage"`
	Aces 			 int 	 `json:"aces"`
	DoubleFaults 	 int 	 `json:"double_faults"`
	Score 			 int 	 `json:"score"`
	ShortestRally    int     `json:"shortest_rally"`
    LongestRally     int     `json:"longest_rally"`
    TotalRallyLength int     `json:"total_rally_length"`
    AverageRallyLength float64 `json:"average_rally_length"`
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
	rallyLengths := make(map[int64]int)

	for i, shot := range shots {
		if _, exists := stats[shot.PlayerId]; !exists {
			stats[shot.PlayerId] = &MatchStats{PlayerID: shot.PlayerId}
		}

		playerStats := stats[shot.PlayerId]
		playerStats.TotalShots++
		rallyLengths[shot.RallyIndex]++
		
		if shot.OrderInRally == 1 && shot.ShotOutcome == shotpb.Outcome_IN {
			if i+1 >= len(shots) || shots[i+1].RallyIndex != shot.RallyIndex {
				playerStats.Aces++
				playerStats.Score++
			}
		}

		if shot.OrderInRally == 1 {
			if shot.ShotOutcome == shotpb.Outcome_IN {
				playerStats.ServeIn++
			} else if shot.ShotOutcome == shotpb.Outcome_OUT || shot.ShotOutcome == shotpb.Outcome_NET {
				playerStats.ServeFault++
			}
		} 
		
		if shot.OrderInRally == 1 && i > 0 && shots[i-1].PlayerId == shot.PlayerId && shots[i-1].OrderInRally == 1 && (shots[i-1].ShotOutcome == shotpb.Outcome_OUT || shots[i-1].ShotOutcome == shotpb.Outcome_NET) {
			playerStats.TotalRallies--
			if shot.ShotOutcome == shotpb.Outcome_IN {
				playerStats.ServeIn++
			} else if shot.ShotOutcome == shotpb.Outcome_OUT || shot.ShotOutcome == shotpb.Outcome_NET {
				playerStats.ServeFault++
				playerStats.DoubleFaults++
				stats[shots[i+1].PlayerId].Score++
			}
		}

		switch shot.ShotOutcome {
		case shotpb.Outcome_IN:
			playerStats.ShotsIn++
		case shotpb.Outcome_OUT:
			playerStats.ShotsFault++
		case shotpb.Outcome_NET:
			playerStats.ShotsFault++
		}

		if !ralliesSeen[shot.RallyIndex] {
			playerStats.TotalRallies++
			ralliesSeen[shot.RallyIndex] = true
		}

        if i > 0 && shot.RallyIndex != shots[i-1].RallyIndex {
            if shot.PlayerId != shots[i-1].PlayerId && (shots[i-1].ShotOutcome == shotpb.Outcome_OUT || shots[i-1].ShotOutcome == shotpb.Outcome_NET) {
				playerStats.Score++
            }

			if shots[i-1].ShotOutcome == shotpb.Outcome_IN && shots[i-1].OrderInRally != 1 {
				stats[shots[i-1].PlayerId].Score++
			}
        }
	}

	for _, playerStats := range stats {
		totalServes := playerStats.ServeIn + playerStats.ServeFault
		if totalServes > 0 {
			playerStats.ServePrecentage = float64(playerStats.ServeIn) / float64(totalServes) * 100
		} else {
			playerStats.ServePrecentage = 0
		}
	}

	// Calculate shortest, longest, and average rally lengths
    for _, length := range rallyLengths {
        for _, playerStats := range stats {
            if playerStats.ShortestRally == -1 || length < playerStats.ShortestRally {
                playerStats.ShortestRally = length
            }
            if length > playerStats.LongestRally {
                playerStats.LongestRally = length
            }
            playerStats.TotalRallyLength += length
        }
    }

    for _, playerStats := range stats {
        if playerStats.TotalRallies > 0 {
            playerStats.AverageRallyLength = float64(playerStats.TotalRallyLength) / float64(playerStats.TotalRallies)
        } else {
            playerStats.AverageRallyLength = 0
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
