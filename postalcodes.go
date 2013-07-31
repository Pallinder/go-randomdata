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
		return randomDigits(2)

	case "FO", "IS", "LS", "MG", "OM", "PG":
		return randomDigits(3)

	case "AF", "AL", "AM", "AU", "AT", "BD", "BE", "BG", "CV", "CY", "DK", "SV",
		"ET", "GE", "GL", "GW", "HT", "HU", "LB", "LR", "LI", "LU", "MK", "MD",
		"NZ", "NE", "NO", "PY", "PH", "ZA", "SJ", "CH", "TN", "VE":
		return randomDigits(4)

	case "DZ", "BA", "KH", "CO", "CR", "HR", "CU", "DO", "EG", "EE", "FI",
		"FR", "GF", "PF", "DE", "GR", "GP", "GT", "HN", "ID", "IR", "IQ", "IL", "IT", "JO",
		"KE", "KV", "KW", "LA", "LY", "MY", "MH", "MQ", "YT", "MX", "FM", "MC", "MN", "ME",
		"MA", "MZ", "MM", "NP", "NC", "PK", "PW", "PR", "RE", "BL", "MF", "PM", "WS", "SM",
		"SA", "RS", "ES", "LK", "SD", "TH", "TR", "UA", "US", "UY", "VA", "WF", "EH", "ZM":
		return randomDigits(5)

	case "BY", "CN", "EC", "IN", "KZ", "KG", "NG", "RO", "RU", "SG", "TJ", "TM", "UZ", "VN":
		return randomDigits(6)

	case "CL":
		return randomDigits(7)

	case "SZ":
		return randomLetters(1) + randomDigits(3)

	case "BM":
		return randomLetters(2) + randomDigits(2)

	case "AD":
		return randomLetters(2) + randomDigits(3)

	case "BN", "AZ", "VG":
		return randomLetters(2) + randomDigits(4)

	case "BB":
		return randomLetters(2) + randomDigits(5)

	case "MT":
		return randomLetters(3) + randomDigits(4)

	case "JM":
		return "JM" + randomLetters(3) + randomDigits(2)

	case "AR":
		return randomLetters(1) + randomDigits(4) + randomLetters(3)

	case "CA":
		return randomLetters(1) + randomDigits(1) + randomLetters(1) + randomDigits(1) + randomLetters(1) + randomDigits(1)

	case "FK", "TC":
		return randomLetters(4) + randomDigits(1) + randomLetters(2)

	case "GG", "IM", "JE", "GB":
		return randomLetters(2) + randomDigits(2) + randomLetters(2)

	case "NL":
		return randomDigits(4) + randomLetters(2)

	case "BR":
		return randomDigits(5) + "-" + randomDigits(3)

	case "KY":
		return randomLetters(2) + randomDigits(1) + "-" + randomDigits(4)

	case "JP":
		return randomDigits(3) + "-" + randomDigits(4)

	case "LV", "SI":
		return randomLetters(2) + "-" + randomDigits(4)

	case "LT", "SE":
		return randomLetters(2) + "-" + randomDigits(5)

	case "MV":
		return randomDigits(2) + "-" + randomDigits(2)

	case "NI":
		return randomDigits(3) + "-" + randomDigits(3) + "-" + randomDigits(1)

	case "PL":
		return randomDigits(2) + "-" + randomDigits(3)

	case "PT":
		return randomDigits(4) + "-" + randomDigits(3)

	case "KR":
		return randomDigits(3) + "-" + randomDigits(3)
	}

	return ""
}

// randomLetters generates a string of N random leters (A-Z).
func randomLetters(letters int) string {
	list := make([]byte, letters)

	for i := range list {
		list[i] = byte(rand.Intn('Z'-'A') + 'A')
	}

	return string(list)
}

// randomDigits generates a string of N random digits, padded with zeros if necessary.
func randomDigits(digits int) string {
	max := int(math.Pow10(digits)) - 1
	num := rand.Intn(max)
	format := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(format, num)
}
