package Main

type MealChoice struct {
	Name string
	Description string
	Ingredients []*Ingredient
}


type Ingredient struct {
	Name string
	Description string
	UnitSize string
	StoredQuantity int
}
