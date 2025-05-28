package service

import (
	"league-simulation/entities"
	"league-simulation/repository"
	"league-simulation/utils"

	"math/rand"
	"sort"
)

var teamRepo repository.TeamRepository = &repository.DBTeamRepository{}
var matchRepo repository.MatchRepository = &repository.DBMatchRepository{}

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
func CreateFixture(teamRepo repository.TeamRepository, matchRepo repository.MatchRepository){

	teams := teamRepo.GetAllTeams()
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
            matchRepo.AddMatch(match)
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
            matchRepo.AddMatch(match)
            matchID++
        }
        week++
    }
}


func SimulateWeek(week int, teamRepo repository.TeamRepository, matchRepo repository.MatchRepository)  []entities.Match{
	matches := matchRepo.GetAllMatches()
	// we created this because if the match is played before it won't be replayed
	var thisWeekPlayed []entities.Match

	for i := range matches {
		match := &matches[i]
		if(match.Week == week && !match.Played){
			// we finf the teams by id for really running the functions
			home := teamRepo.FindTeamByID(match.HomeTeamID)
			away := teamRepo.FindTeamByID(match.AwayTeamID)

			// here we go...
			SimulateMatch(match, home, away)
			UpdateLeagueStatistics(match, home, away)
			
			// let's save it to our repository
			matchRepo.UpdateMatch(match)
			thisWeekPlayed = append(thisWeekPlayed, *match)
		}
	}
	return thisWeekPlayed
}


func GetLeagueTable(teamRepo repository.TeamRepository, matchRepo repository.MatchRepository) []entities.Team {

    teams := teamRepo.GetAllTeams()
	matches := matchRepo.GetAllMatches()
	// sorting starts here regarding the priorities, we used this function because i decided it is the most effective and effortless way
    sort.Slice(teams, func(i, j int) bool {
		// priority 1: points
        if teams[i].Points != teams[j].Points {
            return teams[i].Points > teams[j].Points
        }
        // proirity 2: goal difference
        if teams[i].GoalDifference != teams[j].GoalDifference {
            return teams[i].GoalDifference > teams[j].GoalDifference
        }
        // proirity 3: goals succedded
        if teams[i].GoalsFor != teams[j].GoalsFor {
            return teams[i].GoalsFor > teams[j].GoalsFor
        }
        // priority 4: head-to-head GF
		head2head := headToHead(teams[i], teams[j], matches)
		if head2head != 0 {
			return head2head > 0
		}
		// priority 5: alphabetic order
        return teams[i].Name < teams[j].Name
    })
    return teams
}

func headToHead(a, b entities.Team, matches []entities.Match) int {
    pointsA, pointsB := 0, 0
    goalDiffA, goalDiffB := 0, 0

    for _, m := range matches {
        // only consider matches between team a and team b
        if (m.HomeTeamID == a.ID && m.AwayTeamID == b.ID) || (m.HomeTeamID == b.ID && m.AwayTeamID == a.ID) {
            var ga, gb int // ga: goals scored by a, gb: goals scored by b

            // determine which team is home/away in the match
            if m.HomeTeamID == a.ID {
                ga, gb = m.HomeGoals, m.AwayGoals
            } else {
                ga, gb = m.AwayGoals, m.HomeGoals
            }

            // points calculation for head-to-head
            if ga > gb {
                pointsA += 3
            } else if gb > ga {
                pointsB += 3
            } else {
                pointsA++
                pointsB++
            }

            // goal difference calculation for head-to-head
            goalDiffA += ga - gb
            goalDiffB += gb - ga
        }
    }

    // first, compare head-to-head points
    if pointsA != pointsB {
        return pointsA - pointsB
    }
    // then, compare head-to-head goal difference
    if goalDiffA != goalDiffB {
        return goalDiffA - goalDiffB
    }
    return 0
}

func GetCurrentWeek(matchRepo repository.MatchRepository) int {
    matches := matchRepo.GetAllMatches()
    maxWeek := 0
    for _, m := range matches {
        if m.Played && m.Week > maxWeek {
            maxWeek = m.Week
        }
    }
    return maxWeek
}

func GetTotalWeeks(matchRepo repository.MatchRepository) int {
    matches := matchRepo.GetAllMatches()
    maxWeek := 0
    for _, m := range matches {
        if m.Week > maxWeek {
            maxWeek = m.Week
        }
    }
    return maxWeek
}
// this function calculates the probabilities of being a champion for each team
func ChampionProbabilities(currentWeek, totalWeeks int, teamRepo repository.TeamRepository, matchRepo repository.MatchRepository) map[string]float64 {
    const simCount = 1000

    // get all teams and matches from the repository
    teams := teamRepo.GetAllTeams()
    matches := matchRepo.GetAllMatches()
    remainingWeeks := totalWeeks - currentWeek
    result := make(map[string]float64)

    // if the leader has more points than the maximum points remaining, the champion is already known
    sort.Slice(teams, func(i, j int) bool { return teams[i].Points > teams[j].Points })
    leader, runnerUp := teams[0], teams[1]
    maxLeft := remainingWeeks * 3
    if leader.Points-runnerUp.Points > maxLeft {
        result[leader.Name] = 1.0
        for _, t := range teams[1:] { result[t.Name] = 0.0 }
        return result
    }

    // here comes the Monte Carlo simulation part
    champCounts := make(map[int]int)
    for sim := 0; sim < simCount; sim++ {
        // create in-memory repositories for teams and matches, so the real data won't change
        memTeamRepo := repository.NewMemoryTeamRepo(teams)
        memMatchRepo := repository.NewMemoryMatchRepo(matches)
        // play all the remaining weeks using the real service logic
        for w := currentWeek + 1; w <= totalWeeks; w++ {
            SimulateWeek(w, memTeamRepo, memMatchRepo)
        }
        // after the league ends, find the team with the most points (the champion)
        simTeams := memTeamRepo.GetAllTeams()
        sort.Slice(simTeams, func(i, j int) bool { return simTeams[i].Points > simTeams[j].Points })
        champCounts[simTeams[0].ID]++
    }
    // now, convert the champion counts to probability percentages
    for _, t := range teams {
        result[t.Name] = float64(champCounts[t.ID]) / float64(simCount)
    }
    return result
}





