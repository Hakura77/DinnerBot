package main

import (
	"fmt"
	"github.com/Hakura77/DinnerBot/catalog"
	"github.com/Hakura77/DinnerBot/defs"
	"regexp"
)

const (
	commandOutOfFood   string = "out"
	commandPickMeal           = "pickmeal"
	commandGotFood            = "got"
	commandWhatNeeded         = "foodneeded"
	commandHelp               = "help"
	commandReloadMeals        = "reload_meals_list"
)

func getCommands() []string {
	return []string{commandOutOfFood, commandPickMeal, commandHelp, commandGotFood, commandWhatNeeded}
}

func Help(arguments string) string {
	switch arguments {
	case "commands":
		out := fmt.Sprintf("Currently supported commands are as follows:\n")
		for _, command := range getCommands() {
			out += fmt.Sprintf(" - /%s\n", command)
		}
		return out

	case commandGotFood:
		return fmt.Sprintf("the /%s command accepts two arguments, a food name and a quantity, in order. the food value is then stored and set with a value of the provided number", commandGotFood)

	case commandPickMeal:
		return fmt.Sprintf("the /%s command accepts no arguments, and picks out dinner for you from the stored meals", commandPickMeal)

	case commandOutOfFood:
		return fmt.Sprintf("the /%s command accepts one argument, which is considered a food name, and then marks the food as out in the stored database", commandOutOfFood)
	case commandWhatNeeded:
		return fmt.Sprintf("the /%s command accepts no arguments, and tells you what foods have been marked as out", commandOutOfFood)
	case "":
		// do nothing
	default:
		return fmt.Sprintf("Unrecognized argument %s for /help command", arguments)
	}

	return fmt.Sprintf("Welcome to @Hakuface's Dinner Picking bot. For more information on what I can do, try /help commands, or /help <command> for detailed information")
}

var foodNameValid = regexp.MustCompile(`\A[a-zA-Z0-9 \-\_/\\]+\z`)

func OutOfFood(arguments string) error {
	if !foodNameValid.MatchString(arguments) {
		return fmt.Errorf("argument for out didn't match, food name supports alphanumeric and \\ / - _ special characters only")
	}

	catalog.OutOf(arguments)
	return nil
}

func GotFood(arguments string) (response string) {
	if !foodNameValid.MatchString(arguments) {
		return fmt.Sprintf("argument for got didn't match, food name supports alphanumeric and \\ / - _ special characters only")
	}

	if err := catalog.StoreIngredient(&defs.Ingredient{
		Name:           arguments,
		StoredQuantity: 1,
	}); err != nil {
		return fmt.Sprintf("Problem storing ingredient %s: %s", arguments, err.Error())
	}
	return fmt.Sprintf("Successfully marked %s as in the cupboard", arguments)
}

func FoodNeeded() (response string) {
	ingredientsOut := catalog.IngredientsNeeded()
	if len(ingredientsOut) == 0 {
		return "No Ingredients marked as out of stock"
	}
	response = "Ingredients needed:\n"
	for _, ingredient := range ingredientsOut {
		response += fmt.Sprintf(" - %s: %s\n", ingredient.Name, ingredient.Description)
	}
	return response
}

func PickMeal() (response string) {
	meal, err := catalog.GetRandomMeal()
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Meal Chosen:\n%s", meal.String())

}
