package dao

import ( 
    "golangdemo/api/model"
    "golangdemo/api/utils"
    "fmt"
    "strings"
)



//----------------------------------------------------------------------------
// function initialization
//----------------------------------------------------------------------------
func init() {

	//----------------------------------------------------------------------------
  	// Migrate the schema
  	//----------------------------------------------------------------------------
  	err := utils.GetDB().AutoMigrate(&model.Game{}).Error
  	
  	if  (err == nil ) {
  		fmt.Println( strings.ToTitle( "Finished AutoMigrate on Game" ) )
  	} else {
  	    fmt.Println( strings.ToTitle( "Failed to AutoMigrate on Game" ), err )
  	}
}

//----------------------------------------------------------------------------
// CreateGame - creates a new db entry
//----------------------------------------------------------------------------
func CreateGame(obj model.Game)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult	
	var createMsg string		
	var success bool	
		
	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error 

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a Game with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Game", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, createMsg, "CreateGame", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetGame - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetGame(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult	
	var getMsg string		
	var success bool	

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Game	
	
	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Game with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier
	
	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Game using ID=%v", id )  
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Game using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetGame", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// GetAllGame - returns all 
//----------------------------------------------------------------------------
func GetAllGame()(requestResult utils.RequestResult){	 
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string		
	var success bool	
	var objs []model.Game

	//----------------------------------------------------------------------------
	// Request the ORM to find all Game
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Game" )  
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Game", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, getAllMsg, "GetAllGame", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateGame - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateGame(obj model.Game)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string		
	var success bool
		
	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a Game using ID=%v", obj.ID )  
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Game using ID=%v", obj.ID )
		success = false
	}
	
	requestResult = utils.RequestResult{success, updateMsg, "UpdateGame", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// DeleteGame - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteGame(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string		
	var success bool
	
	//----------------------------------------------------------------------------
	// Obtain the Game with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetGame(id)	
	
	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Game so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Game)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete
		
		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Game using ID=%v", id )  
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Game using ID=%v", id )
			success = false
		}	
		
		requestResult = utils.RequestResult{success, deleteMsg, "DeleteGame", requestResult.Data}
		
	}
	
	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Matchup on a Game
//----------------------------------------------------------------------------
func AssignMatchup( gameId uint64, matchupId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Game with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGame(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Game so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		GameObj,_ := parentRequestResult.Data. (model.Game)	
		
		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var MatchupObj model.Matchup	
	
		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Matchup with a 
		// matching matchupId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&MatchupObj, matchupId).Error // find first using identifier
	
		if childRequestResult == nil {					
			//----------------------------------------------------------------------------
			// assign the Matchup	to the Game	
			//----------------------------------------------------------------------------					
			GameObj.Matchup = MatchupObj

			//----------------------------------------------------------------------------      		
			// save the Game
			//----------------------------------------------------------------------------			
			return UpdateGame(GameObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Matchup", matchupId )
			return utils.RequestResult{false, msg, "assignMatchup", MatchupObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Matchup on a Game
//----------------------------------------------------------------------------
func UnassignMatchup(gameId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Game with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGame(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Game so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		GameObj,_ := parentRequestResult.Data. (model.Game)	
		
		//----------------------------------------------------------------------------      		
		// assign an empty Matchup to the Matchup
		//----------------------------------------------------------------------------      				
		GameObj.Matchup = model.Matchup{}
		
		//----------------------------------------------------------------------------      		
		// save the Game
		//----------------------------------------------------------------------------			
		return UpdateGame(GameObj)
	
	} else {
		return parentRequestResult
	}

}
	
//----------------------------------------------------------------------------
// assigns a Player on a Game
//----------------------------------------------------------------------------
func AssignPlayer( gameId uint64, playerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Game with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGame(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Game so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		GameObj,_ := parentRequestResult.Data. (model.Game)	
		
		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var PlayerObj model.Player	
	
		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Player with a 
		// matching playerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&PlayerObj, playerId).Error // find first using identifier
	
		if childRequestResult == nil {					
			//----------------------------------------------------------------------------
			// assign the Player	to the Game	
			//----------------------------------------------------------------------------					
			GameObj.Player = PlayerObj

			//----------------------------------------------------------------------------      		
			// save the Game
			//----------------------------------------------------------------------------			
			return UpdateGame(GameObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Player", playerId )
			return utils.RequestResult{false, msg, "assignPlayer", PlayerObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Player on a Game
//----------------------------------------------------------------------------
func UnassignPlayer(gameId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Game with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGame(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Game so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		GameObj,_ := parentRequestResult.Data. (model.Game)	
		
		//----------------------------------------------------------------------------      		
		// assign an empty Player to the Player
		//----------------------------------------------------------------------------      				
		GameObj.Player = model.Player{}
		
		//----------------------------------------------------------------------------      		
		// save the Game
		//----------------------------------------------------------------------------			
		return UpdateGame(GameObj)
	
	} else {
		return parentRequestResult
	}

}
	

