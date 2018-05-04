// Package randomdata implements a bunch of simple ways to generate (pseudo) random data
package randomdata

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	Male         int = 0
	Female       int = 1
	RandomGender int = 2
)

const (
	Small int = 0
	Large int = 1
)

const (
	FullCountry      = 0
	TwoCharCountry   = 1
	ThreeCharCountry = 2
)

const (
	DateInputLayout  = "2006-01-02"
	DateOutputLayout = "Monday 2 Jan 2006"
)

type jsonContent struct {
	Adjectives          []string `json:"adjectives"`
	Nouns               []string `json:"nouns"`
	FirstNamesFemale    []string `json:"firstNamesFemale"`
	FirstNamesMale      []string `json:"firstNamesMale"`
	LastNames           []string `json:"lastNames"`
	Domains             []string `json:"domains"`
	People              []string `json:"people"`
	StreetTypes         []string `json:"streetTypes"` // Taken from https://github.com/tomharris/random_data/blob/master/lib/random_data/locations.rb
	Paragraphs          []string `json:"paragraphs"`  // Taken from feedbooks.com and www.gutenberg.org
	Countries           []string `json:"countries"`   // Fetched from the world bank at http://siteresources.worldbank.org/DATASTATISTICS/Resources/CLASS.XLS
	CountriesThreeChars []string `json:"countriesThreeChars"`
	CountriesTwoChars   []string `json:"countriesTwoChars"`
	Currencies          []string `json:"currencies"` //https://github.com/OpenBookPrices/country-data
	Cities              []string `json:"cities"`
	States              []string `json:"states"`
	StatesSmall         []string `json:"statesSmall"`
	Days                []string `json:"days"`
	Months              []string `json:"months"`
	FemaleTitles        []string `json:"femaleTitles"`
	MaleTitles          []string `json:"maleTitles"`
	Timezones           []string `json:"timezones"`           // https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	UserAgents          []string `json:"userAgents"`          // http://techpatterns.com/downloads/firefox/useragentswitcher.xml
	CountryCallingCodes []string `json:"countryCallingCodes"` // from https://github.com/datasets/country-codes/blob/master/data/country-codes.csv
}

var jsonData = jsonContent{}
var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	jsonData = jsonContent{}

	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		log.Fatal(err)
	}
}

// Returns a random part of a slice
func randomFrom(source []string) string {
	return source[r.Intn(len(source))]
}

// Title returns a random title, gender decides the gender of the name
func Title(gender int) string {
	var title = ""
	switch gender {
	case Male:
		title = randomFrom(jsonData.MaleTitles)
		break
	case Female:
		title = randomFrom(jsonData.FemaleTitles)
		break
	default:
		title = FirstName(r.Intn(2))
		break
	}
	return title
}

// FirstName returns a random first name, gender decides the gender of the name
func FirstName(gender int) string {
	var name = ""
	switch gender {
	case Male:
		name = randomFrom(jsonData.FirstNamesMale)
		break
	case Female:
		name = randomFrom(jsonData.FirstNamesFemale)
		break
	default:
		name = FirstName(rand.Intn(2))
		break
	}
	return name
}

// LastName returns a random last name
func LastName() string {
	return randomFrom(jsonData.LastNames)
}

// FullName returns a combination of FirstName LastName randomized, gender decides the gender of the name
func FullName(gender int) string {
	return FirstName(gender) + " " + LastName()
}

// Email returns a random email
func Email() string {
	return strings.ToLower(FirstName(RandomGender)+LastName()) + StringNumberExt(1, "", 3) + "@" + randomFrom(jsonData.Domains)
}

// Country returns a random country, countryStyle decides what kind of format the returned country will have
func Country(countryStyle int64) string {
	country := ""
	switch countryStyle {

	default:

	case FullCountry:
		country = randomFrom(jsonData.Countries)
		break

	case TwoCharCountry:
		country = randomFrom(jsonData.CountriesTwoChars)
		break

	case ThreeCharCountry:
		country = randomFrom(jsonData.CountriesThreeChars)
		break
	}
	return country
}

// Currency returns a random currency under ISO 4217 format
func Currency() string {
	return randomFrom(jsonData.Currencies)
}

// City returns a random city
func City() string {
	return randomFrom(jsonData.Cities)
}

// State returns a random american state
func State(typeOfState int) string {
	if typeOfState == Small {
		return randomFrom(jsonData.StatesSmall)
	}
	return randomFrom(jsonData.States)
}

// Street returns a random fake street name
func Street() string {
	return fmt.Sprintf("%s %s", randomFrom(jsonData.People), randomFrom(jsonData.StreetTypes))
}

// Address returns an american style address
func Address() string {
	return fmt.Sprintf("%d %s,\n%s, %s, %s", Number(100), Street(), City(), State(Small), PostalCode("US"))
}

// Paragraph returns a random paragraph
func Paragraph() string {
	return randomFrom(jsonData.Paragraphs)
}

// Number returns a random number, if only one integer is supplied it is treated as the max value to return
// if a second argument is supplied it returns a number between (and including) the two numbers
func Number(numberRange ...int) int {
	nr := 0
	if len(numberRange) > 1 {
		nr = 1
		nr = r.Intn(numberRange[1]-numberRange[0]) + numberRange[0]
	} else {
		nr = r.Intn(numberRange[0])
	}
	return nr
}

func Decimal(numberRange ...int) float64 {
	nr := 0.0
	if len(numberRange) > 1 {
		nr = 1.0
		nr = r.Float64()*(float64(numberRange[1])-float64(numberRange[0])) + float64(numberRange[0])
	} else {
		nr = r.Float64() * float64(numberRange[0])
	}

	if len(numberRange) > 2 {
		sf := strconv.FormatFloat(nr, 'f', numberRange[2], 64)
		nr, _ = strconv.ParseFloat(sf, 64)
	}
	return nr
}

func StringNumberExt(numberPairs int, separator string, numberOfDigits int) string {
	numberString := ""

	for i := 0; i < numberPairs; i++ {
		for d := 0; d < numberOfDigits; d++ {
			numberString += fmt.Sprintf("%d", Number(0, 9))
		}

		if i+1 != numberPairs {
			numberString += separator
		}
	}

	return numberString
}

// StringNumber returns a random number as a string
func StringNumber(numberPairs int, separator string) string {
	return StringNumberExt(numberPairs, separator, 2)
}

// StringSample returns a random string from a list of strings
func StringSample(stringList ...string) string {
	str := ""
	if len(stringList) > 0 {
		str = stringList[Number(0, len(stringList))]
	}
	return str
}

func Boolean() bool {
	nr := r.Intn(2)
	return nr != 0
}

// Noun returns a random noun
func Noun() string {
	return randomFrom(jsonData.Nouns)
}

// Adjective returns a random adjective
func Adjective() string {
	return randomFrom(jsonData.Adjectives)
}

func uppercaseFirstLetter(word string) string {
	a := []rune(word)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

func lowercaseFirstLetter(word string) string {
	a := []rune(word)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

// SillyName returns a silly name, useful for randomizing naming of things
func SillyName() string {
	return uppercaseFirstLetter(Noun()) + Adjective()
}

// IpV4Address returns a valid IPv4 address as string
func IpV4Address() string {
	blocks := []string{}
	for i := 0; i < 4; i++ {
		number := r.Intn(255)
		blocks = append(blocks, strconv.Itoa(number))
	}

	return strings.Join(blocks, ".")
}

// IpV6Address returns a valid IPv6 address as net.IP
func IpV6Address() string {
	var ip net.IP
	for i := 0; i < net.IPv6len; i++ {
		number := uint8(r.Intn(255))
		ip = append(ip, number)
	}
	return ip.String()
}

// MacAddress returns an mac address string
func MacAddress() string {
	blocks := []string{}
	for i := 0; i < 6; i++ {
		number := fmt.Sprintf("%02x", r.Intn(255))
		blocks = append(blocks, number)
	}

	return strings.Join(blocks, ":")
}

// Day returns random day
func Day() string {
	return randomFrom(jsonData.Days)
}

// Month returns random month
func Month() string {
	return randomFrom(jsonData.Months)
}

// FullDate returns full date
func FullDate() string {
	timestamp := time.Now()
	day := Day()
	month := Month()
	year := timestamp.Year()
	fullDate := day + " " + strconv.Itoa(Number(1, 30)) + " " + month[0:3] + " " + strconv.Itoa(year)
	return fullDate
}

// FullDateInRange returns a date string within a given range, given in the format "2006-01-02".
// If no argument is supplied it will return the result of randomdata.FullDate().
// If only one argument is supplied it is treated as the max date to return.
// If a second argument is supplied it returns a date between (and including) the two dates.
// Returned date is in format "Monday 2 Jan 2006".
func FullDateInRange(dateRange ...string) string {
	var (
		min        time.Time
		max        time.Time
		duration   int
		dateString string
	)
	if len(dateRange) == 1 {
		max, _ = time.Parse(DateInputLayout, dateRange[0])
	} else if len(dateRange) == 2 {
		min, _ = time.Parse(DateInputLayout, dateRange[0])
		max, _ = time.Parse(DateInputLayout, dateRange[1])
	}
	if !max.IsZero() && max.After(min) {
		duration = Number(int(max.Sub(min))) * -1
		dateString = max.Add(time.Duration(duration)).Format(DateOutputLayout)
	} else if !max.IsZero() && !max.After(min) {
		dateString = max.Format(DateOutputLayout)
	} else {
		dateString = FullDate()
	}
	return dateString
}

func Timezone() string {
	return randomFrom(jsonData.Timezones)
}

func UserAgentString() string {
	return randomFrom(jsonData.UserAgents)
}

func PhoneNumber() string {
	str := randomFrom(jsonData.CountryCallingCodes) + " "

	str += Digits(r.Intn(3) + 1)

	for {
		// max 15 chars
		remaining := 15 - (len(str) - strings.Count(str, " "))
		if remaining < 2 {
			return "+" + str
		}
		str += " " + Digits(r.Intn(remaining-1)+1)
	}
}
