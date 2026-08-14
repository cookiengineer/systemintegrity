package controllers

import "github.com/cookiengineer/systemintegrity/structs"
import "strings"

// CollectSteps is the amount of collect actions invoked inside actions.Collect.
// The Welcome view uses it to fill the progress bar, one fraction per step.
const CollectSteps = 12

type Progress struct {
	Step      string
	Completed int
	Total     int
	Fraction  float64
}

func CurrentStep(messages []structs.ConsoleMessage) string {

	stack := make([]string, 0)

	for m := 0; m < len(messages); m++ {

		message := messages[m]

		if message.Method == "Group" {
			stack = append(stack, message.Value)
		} else if message.Method == "GroupEnd" {

			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}

		}

	}

	if len(stack) > 0 {
		return stack[len(stack)-1]
	}

	return "Finished"

}

func isCollectStep(name string) bool {
	return strings.HasPrefix(name, "actions/Collect") && name != "actions/Collect"
}

func ProgressOf(messages []structs.ConsoleMessage) Progress {

	completed := 0

	for m := 0; m < len(messages); m++ {

		message := messages[m]

		if message.Method == "Group" && isCollectStep(message.Value) {
			completed++
		}

	}

	fraction := float64(completed) / float64(CollectSteps)

	if fraction > 1.0 {
		fraction = 1.0
	}

	return Progress{
		Step:      CurrentStep(messages),
		Completed: completed,
		Total:     CollectSteps,
		Fraction:  fraction,
	}

}
