package appmessage

func CheckErroMessage(err error) {
	if err != nil {
		panic(err)
	}
}

func SucessMessage() string {
	return "Operação realizada com sucesso!"
}
