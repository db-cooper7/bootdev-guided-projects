package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/db-cooper7/bootdev-guided-projects/go-rss-aggregator/internal/config"
	"github.com/db-cooper7/bootdev-guided-projects/go-rss-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error creating database connection: %v", err)
	}
	dbQueries := database.New(db)

	runtimeState := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <command_name> args ...string", os.Args[0])
		return
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.run(&runtimeState, cmd); err != nil {
		log.Fatal(err)
	}

}
