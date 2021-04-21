package controller

import (
    MatchupDAO "golangdemo/api/dao"
    "golangdemo/api/model"
    "golangdemo/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MatchupDAO for database creation
//----------------------------------------------------------------------------
func CreateMatchup(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Matchup model
	//----------------------------------------------------------------------------
	data := model.Matchup{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Matchup model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Matchup data access object to create
	//----------------------------------------------------------------------------
	requestResult := MatchupDAO.CreateMatchup( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MatchupDAO to find the relevant Matchup
//----------------------------------------------------------------------------
func GetMatchup(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]
	
	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	
	//----------------------------------------------------------------------------
	// Delegate to the Matchup data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MatchupDAO.GetMatchup(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MatchupDAO for database read of all Matchups
//----------------------------------------------------------------------------
func GetAllMatchup(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Matchup data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MatchupDAO.GetAllMatchup()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MatchupDAO for database save
//----------------------------------------------------------------------------
func UpdateMatchup(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Matchup model
	//----------------------------------------------------------------------------
	var data = model.Matchup{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Matchup model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Matchup data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MatchupDAO.UpdateMatchup(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MatchupDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMatchup(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]

	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	//----------------------------------------------------------------------------
	// Delegate to the Matchup data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MatchupDAO.DeleteMatchup(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more gamesIds as a Games to a Matchup
	//----------------------------------------------------------------------------
func AddGames(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["id"], 10, 64)
	gamesIds,_ := vars["gamesIds"]

	//----------------------------------------------------------------------------		
	// Delegate to the Matchup DAO
	//----------------------------------------------------------------------------		
	requestResult := MatchupDAO.AddGames(gameId, gamesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}			
	
	//----------------------------------------------------------------------------
	// removes one or more gamesIds as a Games from a Matchup
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------						
func RemoveGames(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["id"], 10, 64)
	gamesIds,_ := vars["gamesIds"]

	//----------------------------------------------------------------------------		
	// Delegate to the Matchup DAO
	//----------------------------------------------------------------------------		
	requestResult := MatchupDAO.RemoveGames(gameId, gamesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
