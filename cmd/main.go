package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

// map of Id to team name for clarity. These IDs are what are used in MLB's APIs.
var teamIds = map[uint8]string{
	108: "Angels",
	109: "Diamondbacks",
	110: "Orioles",
	111: "RedSox",
	112: "Cubs",
	113: "Reds",
	114: "Guardians",
	115: "Rockies",
	116: "Tigers",
	117: "Astros",
	118: "Royals",
	119: "Dodgers",
	120: "Nationals",
	121: "Mets",
	133: "Athletics",
	134: "Pirates",
	135: "Padres",
	136: "Mariners",
	137: "Giants",
	138: "Cardinals",
	139: "Rays",
	140: "Rangers",
	141: "BlueJays",
	142: "Twins",
	143: "Phillies",
	144: "Braves",
	145: "WhiteSox",
	146: "Marlins",
	147: "Yankees",
	158: "Brewers",
}

func RosterEvaluation() (string, error) {
	for key := range teamIds {
		resp, err := http.Get(fmt.Sprintf("https://statsapi.mlb.com/api/v1/teams/%v/roster?rosterType=active", key))
		if err != nil {
			// log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			// log.Fatal(err)
		}

		fmt.Println(string(body))
	}
	return "Hello Lambda", nil
}

func main() {
	lambda.Start(RosterEvaluation)
}
