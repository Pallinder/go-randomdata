// Package randomdata implements a bunch of simple ways to generate (pseudo) random data
package randomdata

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// PostalCode yields a random postal/zip code for the given 2-letter country code.
// Supported formats obtained from: http://www.geopostcodes.com/GeoPC_Postal_codes_formats
func PostalCode(countrycode string) string {
	switch strings.ToUpper(countrycode) {
	case "PE":
		return RandomDigits(2)

	case "FO", "IS", "LS", "MG", "OM", "PG":
		return RandomDigits(3)

	case "AF", "AL", "AM", "AU", "AT", "BD", "BE", "BG", "CV", "CY", "DK", "SV",
		"ET", "GE", "GL", "GW", "HT", "HU", "LB", "LR", "LI", "LU", "MK", "MD",
		"NZ", "NE", "NO", "PY", "PH", "ZA", "SJ", "CH", "TN", "VE":
		return RandomDigits(4)

	case "DZ", "BA", "KH", "CO", "CR", "HR", "CU", "DO", "EG", "EE", "FI",
		"FR", "GF", "PF", "DE", "GR", "GP", "GT", "HN", "ID", "IR", "IQ", "IL", "IT", "JO",
		"KE", "KV", "KW", "LA", "LY", "MY", "MH", "MQ", "YT", "MX", "FM", "MC", "MN", "ME",
		"MA", "MZ", "MM", "NP", "NC", "PK", "PW", "PR", "RE", "BL", "MF", "PM", "WS", "SM",
		"SA", "RS", "ES", "LK", "SD", "TH", "TR", "UA", "US", "UY", "VA", "WF", "EH", "ZM":
		return RandomDigits(5)

	case "BY", "CN", "EC", "IN", "KZ", "KG", "NG", "RO", "RU", "SG", "TJ", "TM", "UZ", "VN":
		return RandomDigits(6)

	case "CL":
		return RandomDigits(7)

	case "SZ":
		return RandomLetters(1) + RandomDigits(3)

	case "BM":
		return RandomLetters(2) + RandomDigits(2)

	case "AD":
		return RandomLetters(2) + RandomDigits(3)

	case "BN", "AZ", "VG":
		return RandomLetters(2) + RandomDigits(4)

	case "BB":
		return RandomLetters(2) + RandomDigits(5)

	case "MT":
		return RandomLetters(3) + RandomDigits(4)

	case "JM":
		return "JM" + RandomLetters(3) + RandomDigits(2)

	case "AR":
		return RandomLetters(1) + RandomDigits(4) + RandomLetters(3)

	case "CA":
		return RandomLetters(1) + RandomDigits(1) + RandomLetters(1) + RandomDigits(1) + RandomLetters(1) + RandomDigits(1)

	case "FK", "TC":
		return RandomLetters(4) + RandomDigits(1) + RandomLetters(2)

	case "GG", "IM", "JE", "GB":
		return RandomLetters(2) + RandomDigits(2) + RandomLetters(2)

	case "NL":
		return RandomDigits(4) + RandomLetters(2)

	case "BR":
		return RandomDigits(5) + "-" + RandomDigits(3)

	case "KY":
		return RandomLetters(2) + RandomDigits(1) + "-" + RandomDigits(4)

	case "JP":
		return RandomDigits(3) + "-" + RandomDigits(4)

	case "LV", "SI":
		return RandomLetters(2) + "-" + RandomDigits(4)

	case "LT":
		return RandomLetters(2) + "-" + RandomDigits(5)

	case "SE":
		return "SE" + "-" + RandomLetters(3) + " " + RandomDigits(2)

	case "MV":
		return RandomDigits(2) + "-" + RandomDigits(2)

	case "NI":
		return RandomDigits(3) + "-" + RandomDigits(3) + "-" + RandomDigits(1)

	case "PL":
		return RandomDigits(2) + "-" + RandomDigits(3)

	case "PT":
		return RandomDigits(4) + "-" + RandomDigits(3)

	case "KR":
		return RandomDigits(3) + "-" + RandomDigits(3)
	}

	return ""
}

// RandomLetters generates a string of N random leters (A-Z).
func RandomLetters(letters int) string {
	list := make([]byte, letters)

	for i := range list {
		list[i] = byte(rand.Intn('Z'-'A') + 'A')
	}

	return string(list)
}

// RandomDigits generates a string of N random digits, padded with zeros if necessary.
func RandomDigits(digits int) string {
	max := int(math.Pow10(digits)) - 1
	num := rand.Intn(max)
	format := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(format, num)
}
