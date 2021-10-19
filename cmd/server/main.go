package main
import (
	"fmt"
)
// App - the struct which contains things like pointers 
// to database connection
type App struct {

}
// Run - sets up our application
func (app *App) Run() error{
	fmt.Println("Setting up Our App")
	return nil
}

func main(){
	fmt.Println("api stack")
	app := App{}
	if err := app.Run() ; err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
