package messages

import "fmt"

const (
	OkMessage = "Flujo sin errores"
	FailMessage = "Han ocurrido errores"
)

func HelloWorldStatic () {
	fmt.Println("Hola mundo por default")
}

func HelloWorldWithParameter(value string) string{

	return "Hola mundo con el parámetro "+value
}

func PresetMessage (value string) string{
	presetString:="Mensaje de la ejecución: "
	switch value {
	case "ok":
		return presetString+OkMessage
	case "fail":
		return  presetString+FailMessage
	default:
		return "Algo salió mal :/ "
	}

}