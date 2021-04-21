package router

import (
    PlayerController "golangdemo/api/controller"
    LeagueController "golangdemo/api/controller"
    TournamentController "golangdemo/api/controller"
    MatchupController "golangdemo/api/controller"
    GameController "golangdemo/api/controller"
    jsonResponseFormatter "golangdemo/api/response"
    "github.com/gorilla/mux"
    
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()    


	//----------------------------------------------------------------------------
    // Player Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------
    
    router.HandleFunc("/api/Player/{id}", jsonResponseFormatter.FormatToJSON(PlayerController.GetPlayer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Player", jsonResponseFormatter.FormatToJSON(PlayerController.GetAllPlayer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newPlayer", jsonResponseFormatter.FormatToJSON(PlayerController.CreatePlayer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Player/{id}", jsonResponseFormatter.FormatToJSON(PlayerController.UpdatePlayer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deletePlayer/{id}", jsonResponseFormatter.FormatToJSON(PlayerController.DeletePlayer)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
    // League Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------
    
    router.HandleFunc("/api/League/{id}", jsonResponseFormatter.FormatToJSON(LeagueController.GetLeague)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/League", jsonResponseFormatter.FormatToJSON(LeagueController.GetAllLeague)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newLeague", jsonResponseFormatter.FormatToJSON(LeagueController.CreateLeague)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/League/{id}", jsonResponseFormatter.FormatToJSON(LeagueController.UpdateLeague)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deleteLeague/{id}", jsonResponseFormatter.FormatToJSON(LeagueController.DeleteLeague)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
    // Tournament Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------
    
    router.HandleFunc("/api/Tournament/{id}", jsonResponseFormatter.FormatToJSON(TournamentController.GetTournament)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Tournament", jsonResponseFormatter.FormatToJSON(TournamentController.GetAllTournament)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newTournament", jsonResponseFormatter.FormatToJSON(TournamentController.CreateTournament)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Tournament/{id}", jsonResponseFormatter.FormatToJSON(TournamentController.UpdateTournament)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deleteTournament/{id}", jsonResponseFormatter.FormatToJSON(TournamentController.DeleteTournament)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
    // Matchup Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------
    
    router.HandleFunc("/api/Matchup/{id}", jsonResponseFormatter.FormatToJSON(MatchupController.GetMatchup)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Matchup", jsonResponseFormatter.FormatToJSON(MatchupController.GetAllMatchup)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newMatchup", jsonResponseFormatter.FormatToJSON(MatchupController.CreateMatchup)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Matchup/{id}", jsonResponseFormatter.FormatToJSON(MatchupController.UpdateMatchup)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deleteMatchup/{id}", jsonResponseFormatter.FormatToJSON(MatchupController.DeleteMatchup)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
    // Game Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------
    
    router.HandleFunc("/api/Game/{id}", jsonResponseFormatter.FormatToJSON(GameController.GetGame)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Game", jsonResponseFormatter.FormatToJSON(GameController.GetAllGame)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newGame", jsonResponseFormatter.FormatToJSON(GameController.CreateGame)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Game/{id}", jsonResponseFormatter.FormatToJSON(GameController.UpdateGame)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deleteGame/{id}", jsonResponseFormatter.FormatToJSON(GameController.DeleteGame)).Methods("DELETE", "OPTIONS")
     return router
}