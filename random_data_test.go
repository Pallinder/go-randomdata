package randomdata

import (
	"bytes"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestRandomStringDigits(t *testing.T) {
	t.Log("TestRandomStringDigits")

	if len(StringNumber(2, "-")) != 5 {
		t.Fatal("Wrong length returned")
	}

	if len(StringNumber(2, "")) != 4 {
		t.Fatal("Wrong length returned")
	}

	if len(StringNumberExt(3, "/", 3)) != 11 {
		t.Fatal("Wrong length returned")
	}

	if len(StringNumberExt(3, "", 3)) != 9 {
		t.Fatal("Wrong length returned")
	}
}

func TestFirstName(t *testing.T) {
	t.Log("TestFirstName")
	firstNameMale := FirstName(Male)
	firstNameFemale := FirstName(Female)
	randomName := FirstName(RandomGender)

	if !findInSlice(jsonData.FirstNamesMale, firstNameMale) {
		t.Error("firstNameMale empty or not in male names")
	}

	if !findInSlice(jsonData.FirstNamesFemale, firstNameFemale) {
		t.Error("firstNameFemale empty or not in female names")
	}

	if randomName == "" {
		t.Error("randomName empty")
	}

}

func TestLastName(t *testing.T) {
	t.Log("TestLastName")
	lastName := LastName()

	if !findInSlice(jsonData.LastNames, lastName) {
		t.Error("lastName empty or not in slice")
	}
}

func TestFullName(t *testing.T) {
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

	if !findInSlice(jsonData.FirstNamesMale, maleSplit[0]) {
		t.Error("Couldnt find maleSplit first name in firstNamesMale")
	}

	if !findInSlice(jsonData.LastNames, maleSplit[1]) {
		t.Error("Couldnt find maleSplit last name in lastNames")
	}

	if len(femaleSplit) == 0 {
		t.Error("Failed on full name female")
	}

	if !findInSlice(jsonData.FirstNamesFemale, femaleSplit[0]) {
		t.Error("Couldnt find femaleSplit first name in firstNamesFemale")
	}

	if !findInSlice(jsonData.LastNames, femaleSplit[1]) {
		t.Error("Couldnt find femaleSplit last name in lastNames")
	}

	if len(randomSplit) == 0 {
		t.Error("Failed on full name random")
	}

	if !findInSlice(jsonData.FirstNamesMale, randomSplit[0]) && !findInSlice(jsonData.FirstNamesFemale, randomSplit[0]) {
		t.Error("Couldnt find randomSplit first name in either firstNamesMale or firstNamesFemale")
	}

}

func TestEmail(t *testing.T) {
	t.Log("TestEmail")
	email := Email()

	if email == "" {
		t.Error("Failed to generate email with content")
	}

}

func TestCountry(t *testing.T) {
	t.Log("TestCountry")
	countryFull := Country(FullCountry)
	countryTwo := Country(TwoCharCountry)
	countryThree := Country(ThreeCharCountry)

	if len(countryThree) < 3 {
		t.Error("countryThree < 3 chars")
	}

	if !findInSlice(jsonData.Countries, countryFull) {
		t.Error("Couldnt find country in countries")
	}

	if !findInSlice(jsonData.CountriesTwoChars, countryTwo) {
		t.Error("Couldnt find country with two chars in countriesTwoChars")
	}

	if !findInSlice(jsonData.CountriesThreeChars, countryThree) {
		t.Error("Couldnt find country with three chars in countriesThreeChars")
	}
}

func TestCurrency(t *testing.T) {
	t.Log("TestCurrency")
	if !findInSlice(jsonData.Currencies, Currency()) {
		t.Error("Could not find currency in currencies")
	}
}

func TestCity(t *testing.T) {
	t.Log("TestCity")
	city := City()

	if !findInSlice(jsonData.Cities, city) {
		t.Error("Couldnt find city in cities")
	}
}

func TestParagraph(t *testing.T) {
	t.Log("TestParagraph")
	paragraph := Paragraph()

	if !findInSlice(jsonData.Paragraphs, paragraph) {
		t.Error("Couldnt find paragraph in paragraphs")
	}
}

func TestBool(t *testing.T) {
	t.Log("TestBool")
	booleanVal := Boolean()
	if booleanVal != true && booleanVal != false {
		t.Error("Bool was wrong format")
	}
}

func TestState(t *testing.T) {
	t.Log("TestState")
	stateValSmall := State(Small)
	stateValLarge := State(Large)

	if !findInSlice(jsonData.StatesSmall, stateValSmall) {
		t.Error("Couldnt find small state name in states")
	}

	if !findInSlice(jsonData.States, stateValLarge) {
		t.Error("Couldnt find state name in states")
	}

}

func TestNoun(t *testing.T) {
	if len(jsonData.Nouns) == 0 {
		t.Error("Nouns is empty")
	}

	noun := Noun()

	if !findInSlice(jsonData.Nouns, noun) {
		t.Error("Couldnt find noun in json data")
	}
}

func TestAdjective(t *testing.T) {
	if len(jsonData.Adjectives) == 0 {
		t.Error("Adjectives array is empty")
	}

	adjective := Adjective()

	if !findInSlice(jsonData.Adjectives, adjective) {
		t.Error("Couldnt find noun in json data")
	}
}

func TestSillyName(t *testing.T) {
	sillyName := SillyName()

	if len(sillyName) == 0 {
		t.Error("Couldnt generate a silly name")
	}
}

func TestIpV4Address(t *testing.T) {
	ipAddress := IpV4Address()

	ipBlocks := strings.Split(ipAddress, ".")

	if len(ipBlocks) < 0 || len(ipBlocks) > 4 {
		t.Error("Invalid generated IP address")
	}

	for _, blockString := range ipBlocks {
		blockNumber, err := strconv.Atoi(blockString)

		if err != nil {
			t.Error("Error while testing IpV4Address(): " + err.Error())
		}

		if blockNumber < 0 || blockNumber > 255 {
			t.Error("Invalid generated IP address")
		}
	}
}

func TestIpV6Address(t *testing.T) {
	ipAddress := net.ParseIP(IpV6Address())

	if len(ipAddress) != net.IPv6len {
		t.Errorf("Invalid generated IPv6 address %v", ipAddress)
	}
	roundTripIP := net.ParseIP(ipAddress.String())
	if roundTripIP == nil || !bytes.Equal(ipAddress, roundTripIP) {
		t.Errorf("Invalid generated IPv6 address %v", ipAddress)
	}
}

func TestMacAddress(t *testing.T) {
	t.Log("MacAddress")

	mac := MacAddress()
	if len(mac) != 17 {
		t.Errorf("Invalid generated Mac address %v", mac)
	}

	if !regexp.MustCompile(`([0-9a-fa-f]{2}[:-]){5}([0-9a-fa-f]{2})`).MatchString(mac) {
		t.Errorf("Invalid generated Mac address %v", mac)
	}
}

func TestDecimal(t *testing.T) {
	d := Decimal(2, 4, 3)
	if !(d >= 2 && d <= 4) {
		t.Error("Invalid generate range")
	}

	ds := strings.Split(strconv.FormatFloat(d, 'f', 3, 64), ".")
	if len(ds[1]) != 3 {
		t.Error("Invalid floating point")
	}
}

func TestDay(t *testing.T) {
	t.Log("TestDay")
	day := Day()

	if !findInSlice(jsonData.Days, day) {
		t.Error("Couldnt find day in days")
	}
}

func TestMonth(t *testing.T) {
	t.Log("TestMonth")
	month := Month()

	if !findInSlice(jsonData.Months, month) {
		t.Error("Couldnt find month in months")
	}
}

func TestStringSample(t *testing.T) {
	t.Log("TestStringSample")
	list := []string{"str1", "str2", "str3"}
	str := StringSample(list...)
	if reflect.TypeOf(str).String() != "string" {
		t.Error("Didn't get a string object")
	}
	if !findInSlice(list, str) {
		t.Error("Didn't get string from sample list")
	}
}

func TestStringSampleEmptyList(t *testing.T) {
	t.Log("TestStringSample")
	str := StringSample()
	if reflect.TypeOf(str).String() != "string" {
		t.Error("Didn't get a string object")
	}
	if str != "" {
		t.Error("Didn't get empty string for empty sample list")
	}
}

func TestFullDate(t *testing.T) {
	t.Log("TestFullDate")
	fulldateOne := FullDate()
	fulldateTwo := FullDate()

	if fulldateOne == fulldateTwo {
		t.Error("Invalid random full date")
	}
}

func TestFullDateInRangeNoArgs(t *testing.T) {
	t.Log("TestFullDateInRangeNoArgs")
	fullDate := FullDateInRange()
	_, err := time.Parse(DateOutputLayout, fullDate)

	if err != nil {
		t.Error("Didn't get valid date format.")
	}
}

func TestFullDateInRangeOneArg(t *testing.T) {
	t.Log("TestFullDateInRangeOneArg")
	maxDate, _ := time.Parse(DateInputLayout, "2016-12-31")
	for i := 0; i < 100000; i++ {
		fullDate := FullDateInRange("2016-12-31")
		d, err := time.Parse(DateOutputLayout, fullDate)

		if err != nil {
			t.Error("Didn't get valid date format.")
		}

		if d.After(maxDate) {
			t.Error("Random date didn't match specified max date.")
		}
	}
}

func TestFullDateInRangeTwoArgs(t *testing.T) {
	t.Log("TestFullDateInRangeTwoArgs")
	minDate, _ := time.Parse(DateInputLayout, "2016-01-01")
	maxDate, _ := time.Parse(DateInputLayout, "2016-12-31")
	for i := 0; i < 100000; i++ {
		fullDate := FullDateInRange("2016-01-01", "2016-12-31")
		d, err := time.Parse(DateOutputLayout, fullDate)

		if err != nil {
			t.Error("Didn't get valid date format.")
		}

		if d.After(maxDate) {
			t.Error("Random date didn't match specified max date.")
		}

		if d.Before(minDate) {
			t.Error("Random date didn't match specified min date.")
		}
	}
}

func TestFullDateInRangeSwappedArgs(t *testing.T) {
	t.Log("TestFullDateInRangeSwappedArgs")
	wrongMaxDate, _ := time.Parse(DateInputLayout, "2016-01-01")
	fullDate := FullDateInRange("2016-12-31", "2016-01-01")
	d, err := time.Parse(DateOutputLayout, fullDate)

	if err != nil {
		t.Error("Didn't get valid date format.")
	}

	if d != wrongMaxDate {
		t.Error("Didn't return min date.")
	}
}

func TestTimezone(t *testing.T) {
	t.Log("TestTimezone")
	timezone := Timezone()

	if !findInSlice(jsonData.Timezones, timezone) {
		t.Errorf("Couldnt find timezone in timezones: %v", timezone)
	}
}

func TestUserAgentString(t *testing.T) {
	t.Log("UserAgentString")

	ua := UserAgentString()
	if len(ua) == 0 {
		t.Error("Empty User Agent String")
	}

	if !regexp.MustCompile(`^[a-zA-Z]+\/[0-9]+.[0-9]+\ \(.*\).*$`).MatchString(ua) {
		t.Errorf("Invalid generated User Agent String: %v", ua)
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

func TestPhoneNumbers(t *testing.T) {
	CheckPhoneNumber(PhoneNumber(), t)
}

func CheckPhoneNumber(str string, t *testing.T) bool {
	if (len(str) - strings.Count(str, " ")) > 16 {
		t.Error("phone number too long")
		return false
	}

	matched, err := regexp.MatchString("\\+\\d{1,3}\\s\\d{1,3}", str)

	if err != nil {
		t.Errorf("error matching %v", err)
		return false
	}

	if !matched {
		t.Error("phone number did not match expectations")
		return false
	}

	return true
}
