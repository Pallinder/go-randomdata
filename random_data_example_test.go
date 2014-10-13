package randomdata

import (
	"fmt"
)

func ExampleRandomdata() {

	// Print a male first name
	fmt.Println(FirstName(Male))

	// Print a female first name
	fmt.Println(FirstName(Female))

	// Print a last name
	fmt.Println(LastName())

	// Print a male name
	fmt.Println(FullName(Male))

	// Print a female name
	fmt.Println(FullName(Female))

	// Print a name with random gender
	fmt.Println(FullName(RandomGender))

	// Print a random email
	fmt.Println(Email())

	// Print a country with full text representation
	fmt.Println(Country(FullCountry))

	// Print a country using ISO 3166-1 alpha-3
	fmt.Println(Country(ThreeCharCountry))

	// Print a country using ISO 3166-1 alpha-2
	fmt.Println(Country(TwoCharCountry))

	// Print the name of a random city
	fmt.Println(City())

	// Print the name of a random american state
	fmt.Println(State(Large))

	// Print the name of a random american state using two letters
	fmt.Println(State(Small))

	// Print a random number >= 10 and <= 20
	fmt.Println(Number(10, 20))

	// Print a number >= 0 and <= 20
	fmt.Println(Number(20))

	// Print a paragraph
	fmt.Println(Paragraph())

	// Print a random bool
	fmt.Println(Boolean())

	// Print a random postalcode from Sweden
	fmt.Println(PostalCode("SE"))

	// Print a random american sounding street name
	fmt.Println(Street())

	// Print a random american address
	fmt.Println(Address())

	// Print a random string of numbers
	fmt.Println(StringNumber(2, "-"))

	// Print a random IPv4 address
	fmt.Println(IpV4Address())

}
