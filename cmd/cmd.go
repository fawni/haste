package cmd

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

func Execute() (content string, instance string, raw bool) {
	content = ""
	text := &survey.Multiline{
		Message: "enter text to upload",
	}
	err := survey.AskOne(text, &content, survey.WithValidator(survey.Required))
	if err == terminal.InterruptErr {
		log.Fatalln("interrupted")
	}

	instance = ""
	url := &survey.Select{
		Message: "choose an instance:",
		Options: []string{"https://p.x4.pm", "https://hastebin.cc", "other"},
	}
	survey.AskOne(url, &instance)

	if instance == "other" {
		custom := &survey.Input{
			Message: "enter custom instance:",
		}
		err := survey.AskOne(custom, &instance, survey.WithValidator(survey.Required))
		if err == terminal.InterruptErr {
			log.Fatalln("interrupted")
		}
	}

	raw = false
	prompt := &survey.Confirm{
		Message: "return raw file?",
	}
	survey.AskOne(prompt, &raw)

	return
}
