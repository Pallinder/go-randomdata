package randomdata

import (
	"strings"
	"testing"
)

func TestFirstName(t *testing.T) {
	t.Parallel()
	t.Log("TestFirstName")
	firstNameMale := FirstName(Male)
	firstNameFemale := FirstName(Female)
	randomName := FirstName(RandomGender)

	if !findInSlice(firstNamesMale, firstNameMale) {
		t.Error("firstNameMale empty or not in male names")
	}

	if !findInSlice(firstNamesFemale, firstNameFemale) {
		t.Error("firstNameFemale empty or not in female names")
	}

	if randomName == "" {
		t.Error("randomName empty")
	}

}

func TestLastName(t *testing.T) {
	t.Parallel()
	t.Log("TestLastName")
	lastName := LastName()

	if !findInSlice(lastNames, lastName) {
		t.Error("lastName empty or not in slice")
	}
}

func TestFullName(t *testing.T) {
	t.Parallel()
	t.Log("TestFullName")

	fullNameMale := FullName(Male)
	fullNameFemale := FullName(Female)
	fullNameRandom := FullName(RandomGender)

	maleSplit := strings.Fields(fullNameMale)
	femaleSplit := strings.Fields(fullNameFemale)
	randomSplit := strings.Fields(fullNameRandom)

	if len(maleSplit) == 0 {
		t.Error("Failed on full name male")
	}

	if !findInSlice(firstNamesMale, maleSplit[0]) {
		t.Error("Couldnt find maleSplit first name in firstNamesMale")
	}

	if !findInSlice(lastNames, maleSplit[1]) {
		t.Error("Couldnt find maleSplit last name in lastNames")
	}

	if len(femaleSplit) == 0 {
		t.Error("Failed on full name female")
	}

	if !findInSlice(firstNamesFemale, femaleSplit[0]) {
		t.Error("Couldnt find femaleSplit first name in firstNamesFemale")
	}

	if !findInSlice(lastNames, femaleSplit[1]) {
		t.Error("Couldnt find femaleSplit last name in lastNames")
	}

	if len(randomSplit) == 0 {
		t.Error("Failed on full name random")
	}

	if !findInSlice(firstNamesMale, randomSplit[0]) && !findInSlice(firstNamesFemale, randomSplit[0]) {
		t.Error("Couldnt find randomSplit first name in either firstNamesMale or firstNamesFemale")
	}

}

func TestEmail(t *testing.T) {
	t.Parallel()
	t.Log("TestEmail")
	email := Email()

	if email == "" {
		t.Error("Failed to generate email with content")
	}

}

func TestCountry(t *testing.T) {
	t.Parallel()
	t.Log("TestCountry")
	countryFull := Country(FullCountry)
	countryTwo := Country(TwoCharCountry)
	countryThree := Country(ThreeCharCountry)

	if len(countryThree) < 3 {
		t.Error("countryThree < 3 chars")
	}

	if !findInSlice(countries, countryFull) {
		t.Error("Couldnt find country in countries")
	}

	if !findInSlice(countriesTwoChars, countryTwo) {
		t.Error("Couldnt find country with two chars in countriesTwoChars")
	}

	if !findInSlice(countriesThreeChars, countryThree) {
		t.Error("Couldnt find country with three chars in countriesThreeChars")
	}
}

func TestCity(t *testing.T) {
	t.Parallel()
	t.Log("TestCity")
	city := City()

	if !findInSlice(cities, city) {
		t.Error("Couldnt find city in cities")
	}
}

func TestParagraph(t *testing.T) {
	t.Parallel()
	t.Log("TestParagraph")
	paragraph := Paragraph()

	if !findInSlice(paragraphs, paragraph) {
		t.Error("Couldnt find paragraph in paragraphs")
	}
}

func findInSlice(source []string, toFind string) bool {
	for _, text := range source {
		if text == toFind {
			return true
		}
	}
	return false
}
