package main

import(
	"fmt"
)

func main(){

	ns := GetNutritionalScore(NutritionalData{
		Energy: EnergyFromKcal(250),
		Sugars: SugarGram(10), 
		StaturatedFattyAcids: StaturatedFattyAcids(2),
		Sodium: SodiumMilligram(250),
		Fruits: FruitsPercent(30),
		Fiber: FiberGram(6),
		Protien: ProtienGram(18),
	}, Food)
	
	fmt.Printf("Nutritional Score:%d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}