package main

import (
	"flag"
	"log"

	"github.com/karim-w/gocas/cmd/apps"
	"github.com/karim-w/gocas/cmd/cli"

	"github.com/joho/godotenv"
)

func main() {
	flag.Var(&appFlag, "app", "app(s) to run")
	flag.Parse()

	// ========= Setup Env =========
	godotenv.Load()
	if *envFlag != "" {
		godotenv.Load(*envFlag)
	}

	if len(appFlag) == 0 {
		log.Println("No app to run")
		return
	}

	// ========= Setup Apps =========
	apps := make(map[string]apps.Application)

	for _, app := range appFlag {
		if _, ok := apps[app]; ok {
			continue
		}

		// create app(s)
		switch app {
		case "cli":
			apps["cli"] = cli.CliApp()
		default:
			log.Fatalf("Unknown app: %s", app)
		}
	}

	for _, app := range apps {
		assert(app.Setup())
		app.Close()
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err) // just in case, i have trust issues
	}
}
