package catalog

import (
	"fmt"
	"github.com/Hakura77/DinnerBot/defs"
)

var MealCatalog = map[string]*defs.MealChoice{}

func StoreMeal(in *defs.MealChoice) (err error) {
	if _, ok := MealCatalog[in.Name]; ok {
		// meal is already stored, duplicates disallowed
		return fmt.Errorf("StoreMeal error: Duplicate meal entry for name %s", in.Name)
	}
	MealCatalog[in.Name] = in
	return nil
}


func GetMeal(name string) (meal *defs.MealChoice, err error) {

}