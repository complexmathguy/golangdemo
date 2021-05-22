package dao

import ( 
    "golangdemo/api/model"
    "golangdemo/api/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MatchupDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMatchup - creates a new db entry
//----------------------------------------------------------------------------
func CreateMatchup(obj model.Matchup)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Matchup with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Matchup", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, createMsg, "CreateMatchup", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMatchup - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMatchup(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult	
	var getMsg string		
	var success bool	

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Matchup	
	
	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Matchup with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier
	
	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Matchup using ID=%v", id )  
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Matchup using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMatchup", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMatchup - returns all 
//----------------------------------------------------------------------------
func GetAllMatchup()(requestResult utils.RequestResult){	 
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string		
	var success bool	
	var objs []model.Matchup

	//----------------------------------------------------------------------------
	// Request the ORM to find all Matchup
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Matchup" )  
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Matchup", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMatchup", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMatchup - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMatchup(obj model.Matchup)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Matchup using ID=%v", obj.ID )  
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Matchup using ID=%v", obj.ID )
		success = false
	}
	
	requestResult = utils.RequestResult{success, updateMsg, "UpdateMatchup", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMatchup - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMatchup(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string		
	var success bool
	
	//----------------------------------------------------------------------------
	// Obtain the Matchup with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMatchup(id)	
	
	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matchup so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Matchup)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete
		
		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Matchup using ID=%v", id )  
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Matchup using ID=%v", id )
			success = false
		}	
		
		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMatchup", requestResult.Data}
		
	}
	
	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more gamesIds as a Games to a Matchup
//----------------------------------------------------------------------------
func AddGames( gameId uint64, gamesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Matchup with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatchup(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matchup so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		MatchupObj,_ := parentRequestResult.Data. (model.Matchup)	
				
		// slice the ids on comma with no spaces
		ids := strings.Split( gamesIds, ",")

		for _, gamesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var GameObj model.Game	
	
			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Game 
			// with a matching gamesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&GameObj , gamesId).Error // find first using identifier
	
			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Games using the gorm mechanism
				//----------------------------------------------------------------------------	
				utils.GetDB().Model(&MatchupObj).Association("Games").Append( &GameObj )
				
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Games", gamesId )
				return utils.RequestResult{false, msg, "unassignGames", GameObj}
			}
		}
				
		//----------------------------------------------------------------------------      		
		// retrieve the modified Matchup from the gorm
		//----------------------------------------------------------------------------			
		return GetMatchup(gameId)

	} else {
		return parentRequestResult
	}
}			
	
//----------------------------------------------------------------------------
// removes one or more gamesIds as a Games from a Matchup
//----------------------------------------------------------------------------
func RemoveGames( gameId uint64, gamesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Matchup with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatchup(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matchup so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		MatchupObj,_ := parentRequestResult.Data. (model.Matchup)	
				
		// slice the ids on comma with no spaces
		ids := strings.Split( gamesIds, ",")

		for _, gamesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var GameObj model.Game	
	
			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Game 
			// with a matching gamesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&GameObj , gamesId).Error // find first using identifier
		
			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove GameObj from the Games array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&MatchupObj).Association("Games").Delete( &GameObj )
				
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Games", gamesId )
				return utils.RequestResult{false, msg, "removeGames", GameObj}
			}
		}

		//----------------------------------------------------------------------------      		
		// retrieve the modified Matchup from the gorm
		//----------------------------------------------------------------------------			
		return GetMatchup(gameId)

	} else {
		return parentRequestResult
	}
}
			
