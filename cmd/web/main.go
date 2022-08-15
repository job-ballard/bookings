package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jo-msanii/bookings/internal/config"
	"github.com/jo-msanii/bookings/internal/handlers"
	"github.com/jo-msanii/bookings/internal/render"
)

const portNumber = ":8080"

// 1. Create a variable with the AppConfig Struct
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// 2. The one time we create the template cache, we pull this from the render package
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// 3. We assign the created tc the map of [string] *template file
	app.TemplateCache = tc
	app.UseCache = true

	// so we take app config and set pointer in app config - which is available everywhere
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// 4. We give the NewTemplates func in render pkg access to newly contructed tc struct
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
