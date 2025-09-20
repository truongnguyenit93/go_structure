package script

import (
	"blog/database"
	"blog/pkg/constants"
	"log"
	"os"
	"strings"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func Commands(injector *do.Injector) bool {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)

	var scriptName string

	migrate, seed, run, scripFlag := false, false, false, false

	for _, arg := range os.Args[1:] {
		switch {
		case arg == "--migrate":
			migrate = true
		case arg == "--seed":
			seed = true
		case arg == "--run":
			run = true
		case strings.HasPrefix(arg, "--script="):
			scriptName = strings.TrimPrefix(arg, "--script=")
			scripFlag = true
		}
	}
	if migrate {
		if err := database.Seeder(db); err != nil {
			log.Fatalf("Failed to run migration: %v", err)
		}

		log.Println("Database seeding completed.")
	}

	if seed {
		if err := database.Seeder(db); err != nil {
			log.Fatalf("Failed to run seeder: %v", err)
		}

		log.Println("Database seeding completed.")
	}
	
	if scripFlag {
		switch scriptName {
		case "migrate":
			if err := database.Migrate(db); err != nil {
				log.Fatalf("Failed to run migration: %v", err)
			}

			log.Println("Database migration completed.")
		case "seed":
			if err := database.Seeder(db); err != nil {
				log.Fatalf("Failed to run seeder: %v", err)
			}

			log.Println("Database seeding completed.")
		default:
			log.Printf("Unknown script: %s", scriptName)
		}
	}

	if run {
		return true
	}

	return false
}