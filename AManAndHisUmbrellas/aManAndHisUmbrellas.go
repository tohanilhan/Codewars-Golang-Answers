package kata

func MinUmbrellas(weather []string) int {
	// your code here

	umbrellasNeeded := 0 // number of umbrellas needed

	hasUmbrellaAtHome := false // if he has an umbrella at home
	hasUmbrellaAtWork := false // if he has an umbrella at work
	umbrellaAtHome := 0        // number of umbrellas at home
	umbrellaAtWork := 0        // number of umbrellas at work

	for i, condition := range weather {
		if condition == "rainy" || condition == "thunderstorms" {
			// If it's rainy or thunderstorms, he needs an umbrella, so he takes it from home or work.
			if i%2 == 0 { // if he goes to work
				if hasUmbrellaAtHome || umbrellaAtHome > 0 {
					hasUmbrellaAtHome = false // he takes the umbrella from home to work
					umbrellaAtHome--          // remove one umbrella from home
				} else {
					umbrellasNeeded++ // he needs an umbrella
				}
				hasUmbrellaAtWork = true // he takes the umbrella from home to work
				umbrellaAtWork++         // add one umbrella at work
			} else { // if he goes home
				if hasUmbrellaAtWork || umbrellaAtWork > 0 {
					hasUmbrellaAtWork = false // he takes the umbrella from work to home
					umbrellaAtWork--          // remove one umbrella from work
				} else {
					umbrellasNeeded++ // he needs an umbrella
				}
				hasUmbrellaAtHome = true // he takes the umbrella from work to home
				umbrellaAtHome++         // add one umbrella at home
			}
		}
	}

	return umbrellasNeeded
}
