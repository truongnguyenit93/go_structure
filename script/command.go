package script

import (
	"log"
	"os"
	"strings"
	
	"github.com/samber/do"
	"gorm.io/gorm"
)

func Commands(injector *do.Injector) bool {
	db := do.MustInvokeNamed[*gorm.DB](injector)

	var scriptName string

	migrate := false
	seed := false
	run := false
	scripFlag := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}  
		if arg == "--seed" {
			seed = true
		}
		if arg == "--run" {
			run = true
		}
		if strings.HasPrefix(arg, "--script=") {
			scriptName = strings.TrimPrefix(arg, "--script:")
			scripFlag = true
		}
	}
	if migrate {
		if err 
}