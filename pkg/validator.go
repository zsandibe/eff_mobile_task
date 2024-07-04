package pkg

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

func ValidatePassport(series, number string) error {
	if series != "" {
		matchSeries, _ := regexp.MatchString(`^\d{4}$`, series)
		if !matchSeries {
			return fmt.Errorf("invalid passport series: %s", series)
		}
	}

	if number != "" {
		matchNumber, _ := regexp.MatchString(`^\d{6}$`, number)
		if !matchNumber {
			return fmt.Errorf("invalid passport number format: %s", number)
		}

		if number < "000101" || number > "999999" {
			return fmt.Errorf("passport number out of allowed range: %s", number)
		}
	}

	return nil
}

func ValidatePersonalInfo(name, surname, patronymic, address string) error {
	pattern := `^[А-Я][а-яё]*$`

	if name != "" {
		matchName, err := regexp.MatchString(pattern, name)
		if err != nil {
			return fmt.Errorf("error matching name: %v", err)
		}
		if !matchName {
			return fmt.Errorf("invalid name: %s", name)
		}
	}

	if surname != "" {
		matchSurname, err := regexp.MatchString(pattern, surname)
		if err != nil {
			return fmt.Errorf("error matching surname: %v", err)
		}
		if !matchSurname {
			return fmt.Errorf("invalid surname: %s", surname)
		}
	}

	if patronymic != "" {
		matchPatronymic, _ := regexp.MatchString(`^[А-ЯA-Z][а-яa-zёЁ]{1,}([-][А-ЯA-Z][а-яa-zёЁ]{1,})?$`, patronymic)
		if !matchPatronymic {
			return fmt.Errorf("invalid name: %s", patronymic)
		}
	}

	if address != "" {
		matchAddress, _ := regexp.MatchString(`^[а-яА-Яa-zA-Z0-9\s.,/\\-]{5,}$`, address)
		if !matchAddress {
			return fmt.Errorf("invalid address: %s", address)
		}
	}

	return nil
}

func isASCIIPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			return false
		}
	}
	return true
}

func ValidateStrings(s1, s2 string) bool {

	if !utf8.ValidString(s1) || !utf8.ValidString(s2) {
		return false
	}

	return isASCIIPrintable(s1) && isASCIIPrintable(s2)
}
