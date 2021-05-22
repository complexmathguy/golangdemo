package dao

import ( 
    "golangdemo/api/model"
    "golangdemo/api/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LeagueDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLeague - creates a new db entry
//----------------------------------------------------------------------------
func CreateLeague(obj model.League)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a League with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a League", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, createMsg, "CreateLeague", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLeague - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLeague(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult	
	var getMsg string		
	var success bool	

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.League	
	
	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a League with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier
	
	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a League using ID=%v", id )  
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a League using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLeague", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLeague - returns all 
//----------------------------------------------------------------------------
func GetAllLeague()(requestResult utils.RequestResult){	 
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string		
	var success bool	
	var objs []model.League

	//----------------------------------------------------------------------------
	// Request the ORM to find all League
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all League" )  
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all League", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLeague", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLeague - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLeague(obj model.League)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a League using ID=%v", obj.ID )  
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a League using ID=%v", obj.ID )
		success = false
	}
	
	requestResult = utils.RequestResult{success, updateMsg, "UpdateLeague", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLeague - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLeague(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string		
	var success bool
	
	//----------------------------------------------------------------------------
	// Obtain the League with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLeague(id)	
	
	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.League so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.League)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete
		
		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a League using ID=%v", id )  
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a League using ID=%v", id )
			success = false
		}	
		
		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLeague", requestResult.Data}
		
	}
	
	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more playersIds as a Players to a League
//----------------------------------------------------------------------------
func AddPlayers( gameId uint64, playersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the League with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeague(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.League so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LeagueObj,_ := parentRequestResult.Data. (model.League)	
				
		// slice the ids on comma with no spaces
		ids := strings.Split( playersIds, ",")

		for _, playersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var PlayerObj model.Player	
	
			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Player 
			// with a matching playersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&PlayerObj , playersId).Error // find first using identifier
	
			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Players using the gorm mechanism
				//----------------------------------------------------------------------------	
				utils.GetDB().Model(&LeagueObj).Association("Players").Append( &PlayerObj )
				
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Players", playersId )
				return utils.RequestResult{false, msg, "unassignPlayers", PlayerObj}
			}
		}
				
		//----------------------------------------------------------------------------      		
		// retrieve the modified League from the gorm
		//----------------------------------------------------------------------------			
		return GetLeague(gameId)

	} else {
		return parentRequestResult
	}
}			
	
//----------------------------------------------------------------------------
// removes one or more playersIds as a Players from a League
//----------------------------------------------------------------------------
func RemovePlayers( gameId uint64, playersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the League with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeague(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.League so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		LeagueObj,_ := parentRequestResult.Data. (model.League)	
				
		// slice the ids on comma with no spaces
		ids := strings.Split( playersIds, ",")

		for _, playersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var PlayerObj model.Player	
	
			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Player 
			// with a matching playersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&PlayerObj , playersId).Error // find first using identifier
		
			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PlayerObj from the Players array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&LeagueObj).Association("Players").Delete( &PlayerObj )
				
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Players", playersId )
				return utils.RequestResult{false, msg, "removePlayers", PlayerObj}
			}
		}

		//----------------------------------------------------------------------------      		
		// retrieve the modified League from the gorm
		//----------------------------------------------------------------------------			
		return GetLeague(gameId)

	} else {
		return parentRequestResult
	}
}
			
