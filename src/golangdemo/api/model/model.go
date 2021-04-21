package model

import (
//    "time"
    "gorm.io/gorm"
)


//==============================================================  
// TournamentType Declaration
//==============================================================
type TournamentType int
const (
    Pro TournamentType = iota
	Amateur
)


//==============================================================
// Player Declaration
//==============================================================	
type Player struct {
    gorm.Model
     name            string
    dateOfBirth            string
    height            string
    isProfessional            string
    LeagueId    uint
    GameId    uint
}

//==============================================================
// League Declaration
//==============================================================	
type League struct {
    gorm.Model
     name            string
    Players             []Player
}

//==============================================================
// Tournament Declaration
//==============================================================	
type Tournament struct {
    gorm.Model
     name            string
    Matchups             []Matchup
    Type              TournamentType
}

//==============================================================
// Matchup Declaration
//==============================================================	
type Matchup struct {
    gorm.Model
     name            string
    Games             []Game
    TournamentId    uint
    GameId    uint
}

//==============================================================
// Game Declaration
//==============================================================	
type Game struct {
    gorm.Model
     frames            string
    Matchup            Matchup
    Player            Player
    MatchupId    uint
}

