package ui

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	log "github.com/adem02/epse/internal/utils/logutils"
)

func GetInput(prompt survey.Prompt, response interface{}, validator survey.Validator) {
	if err := survey.AskOne(prompt, response, survey.WithValidator(validator)); err != nil {
		log.Logger{}.Error(fmt.Errorf("\n  interruption détectée. Fermeture...\n"))
		os.Exit(1)
	}
}
