package controller

import (
    TournamentDAO "golangdemo/api/dao"
    "golangdemo/api/model"
    "golangdemo/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TournamentDAO for database creation
//----------------------------------------------------------------------------
func CreateTournament(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Tournament model
	//----------------------------------------------------------------------------
	data := model.Tournament{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Tournament model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Tournament data access object to create
	//----------------------------------------------------------------------------
	requestResult := TournamentDAO.CreateTournament( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TournamentDAO to find the relevant Tournament
//----------------------------------------------------------------------------
func GetTournament(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Tournament data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TournamentDAO.GetTournament(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TournamentDAO for database read of all Tournaments
//----------------------------------------------------------------------------
func GetAllTournament(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Tournament data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TournamentDAO.GetAllTournament()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TournamentDAO for database save
//----------------------------------------------------------------------------
func UpdateTournament(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Tournament model
	//----------------------------------------------------------------------------
	var data = model.Tournament{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Tournament model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Tournament data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TournamentDAO.UpdateTournament(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TournamentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTournament(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Tournament data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TournamentDAO.DeleteTournament(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more matchupsIds as a Matchups to a Tournament
	//----------------------------------------------------------------------------
func AddMatchups(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	matchupsIds,_ := vars["matchupsIds"]

	//----------------------------------------------------------------------------		
	// Delegate to the Tournament DAO
	//----------------------------------------------------------------------------		
	requestResult := TournamentDAO.AddMatchups(gameId, matchupsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}			
	
	//----------------------------------------------------------------------------
	// removes one or more matchupsIds as a Matchups from a Tournament
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------						
func RemoveMatchups(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	matchupsIds,_ := vars["matchupsIds"]

	//----------------------------------------------------------------------------		
	// Delegate to the Tournament DAO
	//----------------------------------------------------------------------------		
	requestResult := TournamentDAO.RemoveMatchups(gameId, matchupsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
