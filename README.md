go-random-data
==============

randomdata is a tiny help suite for generating random data such as 
* first names (male or female)
* last names
* full names (male or female) 
* country names (full name or iso 3166.1 alpha-2 or alpha-3)
* random email address
* city names
* American state names
* random numbers (in an interval)
* random paragraphs 
* random bool values

### Installation
```go get github.com/Pallinder/go-random-data```

### Documentation
http://godoc.org/github.com/Pallinder/go-randomdata

### Usage
```go

package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

func main() {

	// Print a male first name
	fmt.Println(randomdata.FirstName(randomdata.Male))

	// Print a female first name
	fmt.Println(randomdata.FirstName(randomdata.Female))

	// Print a last name
	fmt.Println(randomdata.LastName())

	// Print a male name
	fmt.Println(randomdata.FullName(randomdata.Male))

	// Print a female name
	fmt.Println(randomdata.FullName(randomdata.Female))

	// Print a name with random gender
	fmt.Println(randomdata.FullName(randomdata.RandomGender))

	// Print an email
	fmt.Println(randomdata.Email())

	// Print a country with full text representation
	fmt.Println(randomdata.Country(randomdata.FullCountry))

	// Print a country using ISO 3166-1 alpha-2
	fmt.Println(randomdata.Country(randomdata.TwoCharCountry))

	// Print a country using ISO 3166-1 alpha-3
	fmt.Println(randomdata.Country(randomdata.ThreeCharCountry))

	// Print the name of a random city
	fmt.Println(randomdata.City())

	// Print the name of a random american state
	fmt.Println(randomdata.State())

	// Print a random number >= 10 and <= 20
	fmt.Println(randomdata.Number(10, 20))

	// Print a number >= 0 and <= 20
	fmt.Println(randomdata.Number(20))

	// Print a bool
	fmt.Println(randomdata.Boolean())
	
	// Print a paragraph
	fmt.Println(randomdata.Paragraph())
}

```

### Todo
* add ability to generate complete addresses
