package dao

import ( 
    "golangdemo/api/model"
    "golangdemo/api/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PlayerDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePlayer - creates a new db entry
//----------------------------------------------------------------------------
func CreatePlayer(obj model.Player)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Player with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Player", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, createMsg, "CreatePlayer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPlayer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPlayer(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult	
	var getMsg string		
	var success bool	

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Player	
	
	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Player with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier
	
	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Player using ID=%v", id )  
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Player using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPlayer", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPlayer - returns all 
//----------------------------------------------------------------------------
func GetAllPlayer()(requestResult utils.RequestResult){	 
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string		
	var success bool	
	var objs []model.Player

	//----------------------------------------------------------------------------
	// Request the ORM to find all Player
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Player" )  
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Player", result )
		success = false
	}
	
	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPlayer", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePlayer - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePlayer(obj model.Player)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Player using ID=%v", obj.ID )  
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Player using ID=%v", obj.ID )
		success = false
	}
	
	requestResult = utils.RequestResult{success, updateMsg, "UpdatePlayer", obj}
	
	return requestResult
}

//----------------------------------------------------------------------------
// DeletePlayer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePlayer(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string		
	var success bool
	
	//----------------------------------------------------------------------------
	// Obtain the Player with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPlayer(id)	
	
	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Player so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Player)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete
		
		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Player using ID=%v", id )  
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Player using ID=%v", id )
			success = false
		}	
		
		requestResult = utils.RequestResult{success, deleteMsg, "DeletePlayer", requestResult.Data}
		
	}
	
	return requestResult
}



