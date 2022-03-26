package config

func PanicMaybe(err error) {
	if err != nil {
		panic(err)
	}
}
