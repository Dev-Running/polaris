package utils

import "time"

func ExtractData(data func() (time.Time, bool)) string {
	date, ok := data()

	if ok {

		return date.String()
	}

	return ""

}
