# Pokeapi

Pokeapi is a Go package that interacts with pokeapi (https://pokeapi.co/)
Caching implemented



## Usage
In your ``go.mod`` file, ``require github.com/vahaponur/pokeapi <Latest Version>``
You can check latest version from tags
```go
package main

import (
	"fmt"

	pokeapi "github.com/vahaponur/pokeapi"
)

type LocationArea = pokeapi.LocationArea

func main() {
	//location areas are chunck into 20 groups, you can get next 20 by locArea.Next
	locArea := pokeapi.GetLocationArea(pokeapi.DEFAULT_LOC_AREA)
	pokemon := pokeapi.GetPokemonFromName("pikachu")
	areaDetails := pokeapi.GetPokemonsFromLocArea("canalave-city-area")
	fmt.Println("---First 20 Location Area---")
	for i, locAreaR := range locArea.Results {
		fmt.Printf("Location %v: %v\n", i, locAreaR.Name)
	}
	fmt.Println("---Pokemon Example---")
	fmt.Println(pokemon.Name)
	fmt.Println("---Area Details Example---")
	fmt.Println(areaDetails.Name)

}

//You can check api usage and fields from https://pokeapi.co
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.
