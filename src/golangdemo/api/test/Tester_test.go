package test

import ( 
	"testing"
    dao "golangdemo/api/dao"
	"golangdemo/api/model"
	"golangdemo/api/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
    "time"
)

func init() {
	utils.InitializeEnvironment()
}


func TestPlayerCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Player
	//----------------------------------------------------------------------------
	PlayerObj := model.Player{Name:"test value for Name",DateOfBirth:time.Now(),Height:10,IsProfessional:true}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPlayerRequestResult := dao.CreatePlayer( PlayerObj )
	
	if createPlayerRequestResult.Success == false {
		t.Errorf(createPlayerRequestResult.Msg)
	} else {
		fmt.Println("Check Create Player success...")
	}
	
	createPlayerObj,_ := createPlayerRequestResult.Data. (model.Player)

	// --------------------------------------------------------------
	// Check Player Obj ID
	// --------------------------------------------------------------	
	if createPlayerObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Player" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPlayerRequestResult := dao.GetPlayer( uint64(createPlayerObj.ID) )
	
	if getPlayerRequestResult.Success == false {
		t.Errorf(getPlayerRequestResult.Msg)
	} else {
		fmt.Println("Check Get Player success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPlayerObj,_ := getPlayerRequestResult.Data. (model.Player)
	comparePlayer := cmp.Equal(createPlayerObj.ID, getPlayerObj.ID)
	
	if  comparePlayer == false	{
		t.Errorf( "Created Player object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPlayerRequestResult := dao.GetAllPlayer()

	if getAllPlayerRequestResult.Success == false {
			t.Errorf(getAllPlayerRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Player success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPlayerObj []model.Player = getAllPlayerRequestResult.Data. ([]model.Player)
		
	equalPlayer := cmp.Equal(createPlayerObj.ID, getAllPlayerObj[len(getAllPlayerObj)-1].ID)
		
	if equalPlayer == false {
		t.Errorf( "Created object is not equal to the last entry in Player[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Player
	// --------------------------------------------------------------	
	deletePlayerRequestResult := dao.DeletePlayer(uint64(createPlayerObj.ID))

	if deletePlayerRequestResult.Success == false {
			t.Errorf(deletePlayerRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Player success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPlayerRequestResult = dao.GetPlayer( uint64(createPlayerObj.ID) )
	
	if getPlayerRequestResult.Success == true {
		t.Errorf(getPlayerRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLeagueCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for League
	//----------------------------------------------------------------------------
	LeagueObj := model.League{Name:"test value for Name"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLeagueRequestResult := dao.CreateLeague( LeagueObj )
	
	if createLeagueRequestResult.Success == false {
		t.Errorf(createLeagueRequestResult.Msg)
	} else {
		fmt.Println("Check Create League success...")
	}
	
	createLeagueObj,_ := createLeagueRequestResult.Data. (model.League)

	// --------------------------------------------------------------
	// Check League Obj ID
	// --------------------------------------------------------------	
	if createLeagueObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for League" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLeagueRequestResult := dao.GetLeague( uint64(createLeagueObj.ID) )
	
	if getLeagueRequestResult.Success == false {
		t.Errorf(getLeagueRequestResult.Msg)
	} else {
		fmt.Println("Check Get League success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLeagueObj,_ := getLeagueRequestResult.Data. (model.League)
	compareLeague := cmp.Equal(createLeagueObj.ID, getLeagueObj.ID)
	
	if  compareLeague == false	{
		t.Errorf( "Created League object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLeagueRequestResult := dao.GetAllLeague()

	if getAllLeagueRequestResult.Success == false {
			t.Errorf(getAllLeagueRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll League success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLeagueObj []model.League = getAllLeagueRequestResult.Data. ([]model.League)
		
	equalLeague := cmp.Equal(createLeagueObj.ID, getAllLeagueObj[len(getAllLeagueObj)-1].ID)
		
	if equalLeague == false {
		t.Errorf( "Created object is not equal to the last entry in League[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for League
	// --------------------------------------------------------------	
	deleteLeagueRequestResult := dao.DeleteLeague(uint64(createLeagueObj.ID))

	if deleteLeagueRequestResult.Success == false {
			t.Errorf(deleteLeagueRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion League success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLeagueRequestResult = dao.GetLeague( uint64(createLeagueObj.ID) )
	
	if getLeagueRequestResult.Success == true {
		t.Errorf(getLeagueRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTournamentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Tournament
	//----------------------------------------------------------------------------
	TournamentObj := model.Tournament {Name:"test value for Name",Type:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTournamentRequestResult := dao.CreateTournament( TournamentObj )
	
	if createTournamentRequestResult.Success == false {
		t.Errorf(createTournamentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Tournament success...")
	}
	
	createTournamentObj,_ := createTournamentRequestResult.Data. (model.Tournament)

	// --------------------------------------------------------------
	// Check Tournament Obj ID
	// --------------------------------------------------------------	
	if createTournamentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Tournament" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTournamentRequestResult := dao.GetTournament( uint64(createTournamentObj.ID) )
	
	if getTournamentRequestResult.Success == false {
		t.Errorf(getTournamentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Tournament success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTournamentObj,_ := getTournamentRequestResult.Data. (model.Tournament)
	compareTournament := cmp.Equal(createTournamentObj.ID, getTournamentObj.ID)
	
	if  compareTournament == false	{
		t.Errorf( "Created Tournament object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTournamentRequestResult := dao.GetAllTournament()

	if getAllTournamentRequestResult.Success == false {
			t.Errorf(getAllTournamentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Tournament success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTournamentObj []model.Tournament = getAllTournamentRequestResult.Data. ([]model.Tournament)
		
	equalTournament := cmp.Equal(createTournamentObj.ID, getAllTournamentObj[len(getAllTournamentObj)-1].ID)
		
	if equalTournament == false {
		t.Errorf( "Created object is not equal to the last entry in Tournament[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Tournament
	// --------------------------------------------------------------	
	deleteTournamentRequestResult := dao.DeleteTournament(uint64(createTournamentObj.ID))

	if deleteTournamentRequestResult.Success == false {
			t.Errorf(deleteTournamentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Tournament success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTournamentRequestResult = dao.GetTournament( uint64(createTournamentObj.ID) )
	
	if getTournamentRequestResult.Success == true {
		t.Errorf(getTournamentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMatchupCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Matchup
	//----------------------------------------------------------------------------
	MatchupObj := model.Matchup{Name:"test value for Name"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMatchupRequestResult := dao.CreateMatchup( MatchupObj )
	
	if createMatchupRequestResult.Success == false {
		t.Errorf(createMatchupRequestResult.Msg)
	} else {
		fmt.Println("Check Create Matchup success...")
	}
	
	createMatchupObj,_ := createMatchupRequestResult.Data. (model.Matchup)

	// --------------------------------------------------------------
	// Check Matchup Obj ID
	// --------------------------------------------------------------	
	if createMatchupObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Matchup" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMatchupRequestResult := dao.GetMatchup( uint64(createMatchupObj.ID) )
	
	if getMatchupRequestResult.Success == false {
		t.Errorf(getMatchupRequestResult.Msg)
	} else {
		fmt.Println("Check Get Matchup success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMatchupObj,_ := getMatchupRequestResult.Data. (model.Matchup)
	compareMatchup := cmp.Equal(createMatchupObj.ID, getMatchupObj.ID)
	
	if  compareMatchup == false	{
		t.Errorf( "Created Matchup object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMatchupRequestResult := dao.GetAllMatchup()

	if getAllMatchupRequestResult.Success == false {
			t.Errorf(getAllMatchupRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Matchup success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMatchupObj []model.Matchup = getAllMatchupRequestResult.Data. ([]model.Matchup)
		
	equalMatchup := cmp.Equal(createMatchupObj.ID, getAllMatchupObj[len(getAllMatchupObj)-1].ID)
		
	if equalMatchup == false {
		t.Errorf( "Created object is not equal to the last entry in Matchup[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Matchup
	// --------------------------------------------------------------	
	deleteMatchupRequestResult := dao.DeleteMatchup(uint64(createMatchupObj.ID))

	if deleteMatchupRequestResult.Success == false {
			t.Errorf(deleteMatchupRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Matchup success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMatchupRequestResult = dao.GetMatchup( uint64(createMatchupObj.ID) )
	
	if getMatchupRequestResult.Success == true {
		t.Errorf(getMatchupRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestGameCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Game
	//----------------------------------------------------------------------------
	GameObj := model.Game{Frames:100}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createGameRequestResult := dao.CreateGame( GameObj )
	
	if createGameRequestResult.Success == false {
		t.Errorf(createGameRequestResult.Msg)
	} else {
		fmt.Println("Check Create Game success...")
	}
	
	createGameObj,_ := createGameRequestResult.Data. (model.Game)

	// --------------------------------------------------------------
	// Check Game Obj ID
	// --------------------------------------------------------------	
	if createGameObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Game" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getGameRequestResult := dao.GetGame( uint64(createGameObj.ID) )
	
	if getGameRequestResult.Success == false {
		t.Errorf(getGameRequestResult.Msg)
	} else {
		fmt.Println("Check Get Game success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getGameObj,_ := getGameRequestResult.Data. (model.Game)
	compareGame := cmp.Equal(createGameObj.ID, getGameObj.ID)
	
	if  compareGame == false	{
		t.Errorf( "Created Game object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllGameRequestResult := dao.GetAllGame()

	if getAllGameRequestResult.Success == false {
			t.Errorf(getAllGameRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Game success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllGameObj []model.Game = getAllGameRequestResult.Data. ([]model.Game)
		
	equalGame := cmp.Equal(createGameObj.ID, getAllGameObj[len(getAllGameObj)-1].ID)
		
	if equalGame == false {
		t.Errorf( "Created object is not equal to the last entry in Game[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Game
	// --------------------------------------------------------------	
	deleteGameRequestResult := dao.DeleteGame(uint64(createGameObj.ID))

	if deleteGameRequestResult.Success == false {
			t.Errorf(deleteGameRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Game success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getGameRequestResult = dao.GetGame( uint64(createGameObj.ID) )
	
	if getGameRequestResult.Success == true {
		t.Errorf(getGameRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

