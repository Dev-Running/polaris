package utils

func ExtractString(data func() (string, bool)) string {
	date, ok := data()

	if ok {

		return date
	}

	return ""

}
