package service

import (
    "league-simulation/entities"
    "league-simulation/repository"
	"league-simulation/utils"
    "math/rand"
)


// there is a basic algorithm according to teams' strengths
func SimulateGoals(strength int) int {
    maxGoals := strength / 20 + 2  
    return rand.Intn(maxGoals)
}

// here, there is a match simulation
func SimulateMatch(match *entities.Match, homeTeam *entities.Team, awayTeam *entities.Team)  {
	if homeTeam == nil || awayTeam == nil || match == nil{
		return
	}
	match.HomeGoals = SimulateGoals(homeTeam.Strength)
	match.AwayGoals = SimulateGoals(awayTeam.Strength)
	match.Played = true
}



// as in real life, the statistics are updated here
func UpdateLeagueStatistics(match *entities.Match, homeTeam, awayTeam *entities.Team){

	// if the teams and the match are not found, return
	if homeTeam == nil || awayTeam == nil || match == nil {
        return
    }

	homeTeam.MatchesPlayed++
	awayTeam.MatchesPlayed++

	// GF and GA are updated here
	homeTeam.GoalsFor += match.HomeGoals
	awayTeam.GoalsFor += match.AwayGoals
	homeTeam.GoalsAgainst += match.AwayGoals
	awayTeam.GoalsAgainst += match.HomeGoals

	// GD is updated here
	homeTeam.GoalDifference = homeTeam.GoalsFor - homeTeam.GoalsAgainst
	awayTeam.GoalDifference = awayTeam.GoalsFor - awayTeam.GoalsAgainst

	// score adjustment
	if match.HomeGoals > match.AwayGoals {
		homeTeam.MatchesWon++
		awayTeam.MatchesLost++
		homeTeam.Points += 3
	} else if match.AwayGoals > match.HomeGoals {
		homeTeam.MatchesLost++
		awayTeam.MatchesWon++
		awayTeam.Points += 3
	} else {
		homeTeam.MatchesDrawn++
		awayTeam.MatchesDrawn++
		homeTeam.Points += 1
		awayTeam.Points += 1
	}
}



// the whole fixture is created here
func CreateFixture(){

	teams := repository.GetAllTeams()
	numberOfTeams := len(teams)
	if numberOfTeams != 4 {
    	panic("CreateFixture currently only supports 4 teams.")
	}
	teams = utils.ShuffleTeams(teams)
	matchID := 1
	week := 1

	// create a round robin for each leg as in real life football tournaments and leagues
	var rounds = [][][2] int {
		{{0,1},{2,3}},
		{{0,2},{1,3}},
		{{0,3},{1,2}},
	}
	
	// create the first leg
	for _, matches := range rounds {
        for _, pair := range matches {
            match := entities.Match{
                ID:       matchID,
                Week:     week,
                HomeTeamID: teams[pair[0]].ID,
                AwayTeamID: teams[pair[1]].ID,
                Played:   false,
            }
            repository.AddMatch(match)
            matchID++
        }
        week++
    }

	// create the return matches
	for _, matches := range rounds {
        for _, pair := range matches {
            match := entities.Match{
                ID:       matchID,
                Week:     week,
                HomeTeamID: teams[pair[1]].ID,
                AwayTeamID: teams[pair[0]].ID,
                Played:   false,
            }
            repository.AddMatch(match)
            matchID++
        }
        week++
    }
}


func SimulateWeek(week int) []entities.Match{
	matches := repository.GetAllMatches()
	// we created this because if the match is played before it won't be replayed
	var thisWeekPlayed []entities.Match

	for i := range matches {
		match := &matches[i]
		if(match.Week == week && !match.Played){
			// we finf the teams by id for really running the functions
			home := repository.FindTeamByID(match.HomeTeamID)
			away := repository.FindTeamByID(match.AwayTeamID)

			// here we go...
			SimulateMatch(match, home, away)
			UpdateLeagueStatistics(match, home, away)
			
			// let's save it to our repository
			repository.UpdateMatch(match)
			thisWeekPlayed = append(thisWeekPlayed, *match)
		}
	}
	return thisWeekPlayed
}





