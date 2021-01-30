package main

import (
	"fmt"
	"github.com/Hakura77/DinnerBot/catalog"
	"github.com/Hakura77/DinnerBot/defs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const mealFilePath = "./storage/meals.yaml"

func LoadMeals() (err error) {
	mealFile, err := ioutil.ReadFile(mealFilePath)
	if err != nil {
		return  fmt.Errorf("unable to read %s: %w", mealFilePath, err)
	}


	meals := make([]*defs.MealChoice, 0)
	if err := yaml.Unmarshal(mealFile, &meals); err != nil {
		return fmt.Errorf("error unmarshalling meals file: %w", err)
	}
	for _, meal := range meals {
		err := catalog.StoreMeal(meal)
		if err != nil {
			return fmt.Errorf("error storing meal %s: %w", meal.Name, err)
		}
	}
	return nil

}
