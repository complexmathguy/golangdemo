package model

import (
    "time"
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
     Name            string
    DateOfBirth            time.Time
    Height            int64
    IsProfessional            bool
    LeagueId    uint
    GameId    uint
}

//==============================================================
// League Declaration
//==============================================================	
type League struct {
    gorm.Model
     Name            string
    Players             []Player
}

//==============================================================
// Tournament Declaration
//==============================================================	
type Tournament struct {
    gorm.Model
     Name            string
    Matchups             []Matchup
    Type              TournamentType
}

//==============================================================
// Matchup Declaration
//==============================================================	
type Matchup struct {
    gorm.Model
     Name            string
    Games             []Game
    TournamentId    uint
    GameId    uint
}

//==============================================================
// Game Declaration
//==============================================================	
type Game struct {
    gorm.Model
     Frames            int32
    Matchup            Matchup
    Player            Player
    MatchupId    uint
}

