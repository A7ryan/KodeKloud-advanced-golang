package main

func main() {
	app := App{}
	app.Initialise(Dbuser, DbPassword, DbName)
	app.Run("localhost:7000")
}
