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
  	err := utils.GetDB().AutoMigrate(&model.Tournament{}).Error
  	
  	if  (err == nil ) {
  		fmt.Println( strings.ToTitle( "Finished AutoMigrate on Tournament" ) )
  	} else {
  	    fmt.Println( strings.ToTitle( "Failed to AutoMigrate on Tournament" ), err )
  	}
}

//----------------------------------------------------------------------------
// CreateTournament - creates a new db entry
//----------------------------------------------------------------------------
func CreateTournament(obj model.Tournament)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Tournament with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Tournament", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, createMsg, "CreateTournament", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTournament - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTournament(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult	
	var getMsg string		
	var success bool	

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Tournament	
	
	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Tournament with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier
	
	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Tournament using ID=%v", id )  
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Tournament using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTournament", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTournament - returns all 
//----------------------------------------------------------------------------
func GetAllTournament()(requestResult utils.RequestResult){	 
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string		
	var success bool	
	var objs []model.Tournament

	//----------------------------------------------------------------------------
	// Request the ORM to find all Tournament
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Tournament" )  
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Tournament", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTournament", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTournament - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTournament(obj model.Tournament)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Tournament using ID=%v", obj.ID )  
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Tournament using ID=%v", obj.ID )
		success = false
	}
	
	requestResult = utils.RequestResult{success, updateMsg, "UpdateTournament", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTournament - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTournament(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string		
	var success bool
	
	//----------------------------------------------------------------------------
	// Obtain the Tournament with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTournament(id)	
	
	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tournament so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Tournament)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete
		
		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Tournament using ID=%v", id )  
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Tournament using ID=%v", id )
			success = false
		}	
		
		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTournament", requestResult.Data}
		
	}
	
	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more matchupsIds as a Matchups to a Tournament
//----------------------------------------------------------------------------
func AddMatchups( gameId uint64, matchupsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Tournament with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTournament(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tournament so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TournamentObj,_ := parentRequestResult.Data. (model.Tournament)	
				
		// slice the ids on comma with no spaces
		ids := strings.Split( matchupsIds, ",")

		for _, matchupsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var MatchupObj model.Matchup	
	
			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Matchup 
			// with a matching matchupsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&MatchupObj , matchupsId).Error // find first using identifier
	
			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Matchups using the gorm mechanism
				//----------------------------------------------------------------------------	
				utils.GetDB().Model(&TournamentObj).Association("Matchups").Append( &MatchupObj )
				
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Matchups", matchupsId )
				return utils.RequestResult{false, msg, "unassignMatchups", MatchupObj}
			}
		}
				
		//----------------------------------------------------------------------------      		
		// retrieve the modified Tournament from the gorm
		//----------------------------------------------------------------------------			
		return GetTournament(gameId)

	} else {
		return parentRequestResult
	}
}			
	
//----------------------------------------------------------------------------
// removes one or more matchupsIds as a Matchups from a Tournament
//----------------------------------------------------------------------------
func RemoveMatchups( gameId uint64, matchupsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Tournament with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTournament(gameId)	
	
	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Tournament so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		TournamentObj,_ := parentRequestResult.Data. (model.Tournament)	
				
		// slice the ids on comma with no spaces
		ids := strings.Split( matchupsIds, ",")

		for _, matchupsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var MatchupObj model.Matchup	
	
			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Matchup 
			// with a matching matchupsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&MatchupObj , matchupsId).Error // find first using identifier
		
			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MatchupObj from the Matchups array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&TournamentObj).Association("Matchups").Delete( &MatchupObj )
				
			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Matchups", matchupsId )
				return utils.RequestResult{false, msg, "removeMatchups", MatchupObj}
			}
		}

		//----------------------------------------------------------------------------      		
		// retrieve the modified Tournament from the gorm
		//----------------------------------------------------------------------------			
		return GetTournament(gameId)

	} else {
		return parentRequestResult
	}
}
			
