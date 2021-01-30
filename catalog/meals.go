package catalog

import (
	"fmt"
	"github.com/Hakura77/DinnerBot/defs"
	"math/rand"
	"time"
)

func init() {
	randomGen = rand.New(rand.NewSource(time.Now().UnixNano()))
}


var randomGen *rand.Rand


var mealCatalog = map[string]*defs.MealChoice{}

func StoreMeal(in *defs.MealChoice) (err error) {
	for _, ingredient := range in.IngredientRequirements {
		if ingredient.Quantity == 0 {
			ingredient.Quantity = 1
		}
	}
	if _, ok := mealCatalog[in.Name]; ok {
		// meal is already stored, duplicates disallowed
		return fmt.Errorf("StoreMeal error: Duplicate meal entry for name %s", in.Name)
	}
	mealCatalog[in.Name] = in
	return nil
}


func GetMeal(name string) (meal *defs.MealChoice, err error) {

	//todo check ingredient library


	if meal, ok := mealCatalog[name]; ok {
		return meal, nil
	}
	return nil, fmt.Errorf("GetMeal error: no meal found for name %s", name)

}


func GetRandomMeal() (meal *defs.MealChoice, err error) {

	if len(mealCatalog) == 0 {
		return nil, fmt.Errorf("GetRandomMeal error: No meals defined")
	}


	keys := make([]string, 0, len(mealCatalog))
	for key := range mealCatalog {
		keys = append(keys, key)
	}

	if len(keys) == 1 {
		return mealCatalog[keys[0]], nil
	}

	return mealCatalog[keys[randomGen.Intn(len(keys)-1)]], nil
}