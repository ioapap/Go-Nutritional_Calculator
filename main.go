package main

import "fmt"

func main() {

	ns := GetNutritionalScore(NutritionalData{
		// Example used below: Chicken breast
		Energy:              EnergyFromKcal(110),
		Sugars:              SugarGram(0),
		SaturatedFattyAcids: SaturatedFattyAcids(0.33),
		Sodium:              SodiumMilligram(65),
		Fruits:              FruitsPercent(0),
		Fibre:               FibreGram(0),
		Protein:             ProteinGram(23),
	}, Food)
	fmt.Printf("Nutritional Score:%d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
