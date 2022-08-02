package appmessage

func CheckErroMessage(err error) {
	if err != nil {
		panic(err)
	}
}
