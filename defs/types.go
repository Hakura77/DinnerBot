package defs

import "fmt"

type MealChoice struct {
	Name string `yaml:"Name"`
	Description string `yaml:"Description"`
	IngredientRequirements []*IngredientRequirement `yaml:"IngredientRequirements,omitempty"`
}

func (m *MealChoice) NeedsIngredient(in *Ingredient) bool {
	for _, requirement := range m.IngredientRequirements {
		if requirement.Name == in.Name {
			return true
		}
	}
	return false
}

func (m *MealChoice) String() string {
	out := fmt.Sprintf("Name: %s\nDescription: %s\nIngredients:", m.Name, m.Description)
	for _, ing := range m.IngredientRequirements {
		out += fmt.Sprintf("\n - %s", ing.String())
	}
	return out
}


type IngredientRequirement struct {
	Name string `yaml:"Name"`
	Quantity int `yaml:"Quantity" default:1`
}

func (i *IngredientRequirement) MetBy(in *Ingredient) bool {
	if in.Name != i.Name {
		return false
	}
	if i.Quantity > in.StoredQuantity {
		return false
	}
	return true
}
func (i *IngredientRequirement) String() string {
	return fmt.Sprintf("Name: %s, Quantity Needed: %d", i.Name, i.Quantity)
}


type Ingredient struct {
	Name string
	Description string
	UnitSize string
	StoredQuantity int
}

func (m *Ingredient) Matches(in *Ingredient) bool {
	if m.Name != in.Name {
		return false
	}
	return true
}