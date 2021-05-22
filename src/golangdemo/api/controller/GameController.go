package controller

import (
    GameDAO "golangdemo/api/dao"
    "golangdemo/api/model"
    "golangdemo/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to GameDAO for database creation
//----------------------------------------------------------------------------
func CreateGame(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Game model
	//----------------------------------------------------------------------------
	data := model.Game{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Game model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Game data access object to create
	//----------------------------------------------------------------------------
	requestResult := GameDAO.CreateGame( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to GameDAO to find the relevant Game
//----------------------------------------------------------------------------
func GetGame(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Game data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GameDAO.GetGame(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to GameDAO for database read of all Games
//----------------------------------------------------------------------------
func GetAllGame(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Game data access object to get all
	//----------------------------------------------------------------------------
	requestResult := GameDAO.GetAllGame()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to GameDAO for database save
//----------------------------------------------------------------------------
func UpdateGame(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Game model
	//----------------------------------------------------------------------------
	var data = model.Game{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Game model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Game data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GameDAO.UpdateGame(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to GameDAO for database deletion
//----------------------------------------------------------------------------
func DeleteGame(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Game data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := GameDAO.DeleteGame(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Matchup on a Game
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMatchup(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id params
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	matchupId,_ := strconv.ParseUint( vars["matchupId"], 10, 64)

	//----------------------------------------------------------------------------		
	// Delegate to the Game DAO
	//----------------------------------------------------------------------------		
	requestResult := GameDAO.AssignMatchup(gameId, matchupId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Matchup on a Game
	// delegates to the ORM handler
	//----------------------------------------------------------------------------			
func UnassignMatchup( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id params
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------		
	// Delegate to the Game DAO
	//----------------------------------------------------------------------------		
	requestResult := GameDAO.UnassignMatchup(gameId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
	
	//----------------------------------------------------------------------------
	// assigns a Player on a Game
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlayer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id params
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	playerId,_ := strconv.ParseUint( vars["playerId"], 10, 64)

	//----------------------------------------------------------------------------		
	// Delegate to the Game DAO
	//----------------------------------------------------------------------------		
	requestResult := GameDAO.AssignPlayer(gameId, playerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Player on a Game
	// delegates to the ORM handler
	//----------------------------------------------------------------------------			
func UnassignPlayer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id params
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------		
	// Delegate to the Game DAO
	//----------------------------------------------------------------------------		
	requestResult := GameDAO.UnassignPlayer(gameId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
	

