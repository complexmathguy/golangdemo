package controller

import (
    LeagueDAO "golangdemo/api/dao"
    "golangdemo/api/model"
    "golangdemo/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LeagueDAO for database creation
//----------------------------------------------------------------------------
func CreateLeague(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty League model
	//----------------------------------------------------------------------------
	data := model.League{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a League model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the League data access object to create
	//----------------------------------------------------------------------------
	requestResult := LeagueDAO.CreateLeague( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LeagueDAO to find the relevant League
//----------------------------------------------------------------------------
func GetLeague(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the League data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LeagueDAO.GetLeague(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LeagueDAO for database read of all Leagues
//----------------------------------------------------------------------------
func GetAllLeague(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the League data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LeagueDAO.GetAllLeague()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LeagueDAO for database save
//----------------------------------------------------------------------------
func UpdateLeague(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty League model
	//----------------------------------------------------------------------------
	var data = model.League{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a League model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the League data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LeagueDAO.UpdateLeague(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LeagueDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLeague(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the League data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LeagueDAO.DeleteLeague(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more playersIds as a Players to a League
	//----------------------------------------------------------------------------
func AddPlayers(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	playersIds,_ := vars["playersIds"]

	//----------------------------------------------------------------------------		
	// Delegate to the League DAO
	//----------------------------------------------------------------------------		
	requestResult := LeagueDAO.AddPlayers(gameId, playersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}			
	
	//----------------------------------------------------------------------------
	// removes one or more playersIds as a Players from a League
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------						
func RemovePlayers(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------		
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------		
	gameId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	playersIds,_ := vars["playersIds"]

	//----------------------------------------------------------------------------		
	// Delegate to the League DAO
	//----------------------------------------------------------------------------		
	requestResult := LeagueDAO.RemovePlayers(gameId, playersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
