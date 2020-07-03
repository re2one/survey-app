
package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	// "github.com/manakuro/golang-clean-architecture/config"
	"survey-app-backend/frameworks/persistence"
	"survey-app-backend/frameworks/router"
	"survey-app-backend/frameworks/registry"
)

func main() {
	// config.ReadConfig()

	db := persistence.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + "8081")
	//config.C.Server.Address)
	if err := e.Start(":" + "8081"); err != nil {
	// config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}