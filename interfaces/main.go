package main

type dbcontract interface {
	Close()
	InsertUser(username string) error
	SelectingUser(username string) (string, error)
}

type Application struct {
	db dbcontract
}

func NewApplication(db dbcontract) *Application {
	return &Application{db: db}
}

func main() {
	db, err := mysqldb.New()
	app := NewApplication(db)
}