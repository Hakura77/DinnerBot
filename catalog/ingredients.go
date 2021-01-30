package catalog

import (
	"github.com/Hakura77/DinnerBot/defs"
)

var ingredientCatalog = map[string]*defs.Ingredient{}

func StoreIngredient(in *defs.Ingredient) (err error) {
	if _, ok := ingredientCatalog[in.Name]; ok {
		// ingredient is already stored, duplicates disallowed
		ingredientCatalog[in.Name].StoredQuantity = in.StoredQuantity
	}
	ingredientCatalog[in.Name] = in
	return nil
}

func OutOf(name string) {
	if entry, ok := ingredientCatalog[name]; ok {
		entry.StoredQuantity = 0
		return
	}
	ingredientCatalog[name] = &defs.Ingredient{
		Name:           name,
		StoredQuantity: 0,
	}
	return

}

func IngredientsNeeded() (out []*defs.Ingredient) {
	out = make([]*defs.Ingredient, 0)
	for _, ingredient := range ingredientCatalog {
		if ingredient.StoredQuantity == 0 {
			out = append(out, ingredient)
		}
	}
	return
}

//HaveIngredient returns true if either a. ingredient is not stored in catalog or b. Ingredient is stored in catalog and quantity is >0
func HaveIngredient(name string) bool {
	if ingredient, ok := ingredientCatalog[name]; ok {
		if ingredient.StoredQuantity > 0 {
			return true
		}
		return false
	}
	return true
}

func HaveEnoughIngredient(name string, quantity int) bool {
	if ingredient, ok := ingredientCatalog[name]; ok {
		if ingredient.StoredQuantity > quantity {
			return true
		}
		return false
	}
	return true
}

func CanCook(in *defs.MealChoice) bool {
	for _, requirements := range in.IngredientRequirements {
		if !HaveEnoughIngredient(requirements.Name, requirements.Quantity) {
			return false
		}
	}
	return true
}
