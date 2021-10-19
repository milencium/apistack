package main
import (
	"fmt"
	"net/http"
	database "github.com/milencium/apistack/internal/database"
	transportHTTP "github.com/milencium/apistack/internal/transport/http"
	"github.com/milencium/apistack/internal/comment"
)
// App - the struct which contains things like pointers 
// to database connection
type App struct {

}
// Run - sets up our application
func (app *App) Run() error{
	fmt.Println("Setting up Our App")

	//Defining database connection
	var err error 
	db, err := database.NewDatabase()
	if err != nil{
		fmt.Println("database error")
		return err
	}
	//Running migrations for database
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	//Ode se kreira new service na razini servera
	commentService := comment.NewService(db)

	//Novo kreirani service se stavlja u new handlera na razini servera
	//Defining server
	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()


	
	if err := http.ListenAndServe(":8080", handler.Router); err != nil{
		fmt.Println("Failed to set up server")
		return err 
	}
	
	return nil

}

func main(){
	fmt.Println("api stack")
	app := App{}
	//Run je metoda na struct App koja se ode poziva
	if err := app.Run() ; err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
