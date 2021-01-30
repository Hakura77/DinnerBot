package Main

type MealChoice struct {
	Name string
	Description string
	Ingredients []*Ingredient
}


func (m *MealChoice) HasIngredient(in Ingredient) {
	for _, ingredient := range m.Ingredients {
		if ingredient.
	}



}


type Ingredient struct {
	Name string
	Description string
	UnitSize string
	StoredQuantity int
}

func (m *Ingredient) Matches