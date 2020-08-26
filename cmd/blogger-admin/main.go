// This program performs administrative tasks for the garage sale service.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mavengopher/gouserview/internal/platform/conf"
	"github.com/mavengopher/gouserview/internal/platform/database"
	"github.com/mavengopher/gouserview/internal/schema"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error: shutting down: %s", err)
		os.Exit(1)
	}
}

func run() error {

	// =========================================================================
	// Configuration

	var cfg struct {
		DB struct {
			URI  string `conf:"default:mongodb://localhost:5432"`
			Name string `conf:"default:eBlogger"`
		}
		Args conf.Args
	}

	if err := conf.Parse(os.Args[1:], "BLOGS", &cfg); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("BLOGS", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating usage")
			}
			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "error: parsing config")
	}

	// Initialize dependencies.
	db, ctx, err := database.Open(database.Config{
		URI:  cfg.DB.URI,
		Name: cfg.DB.Name,
	})
	if err != nil {
		return errors.Wrap(err, "connecting to db")
	}
	defer db.Client().Disconnect(ctx)

	switch cfg.Args.Num(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			return errors.Wrap(err, "applying migrations")
		}
		fmt.Println("Migrations complete")

	case "seed":
		if err := schema.Seed(db); err != nil {
			return errors.Wrap(err, "seeding database")
		}
		fmt.Println("Seed data complete")
	}

	return nil
}
