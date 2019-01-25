package main

import (
	"awesomeProject/firstApplication/src/fakeElements"
	"awesomeProject/firstApplication/src/messages"
	"fmt"
)
type Persona struct {
	id int32
	nombre string
	apellido string
}


func main() {
	var users []fakeElements.User
	var tasks []fakeElements.Task
	users = fakeElements.GetArrayUser()
	tasks = fakeElements.GetArrayTask()
	fmt.Println(users)
	fmt.Println(tasks)




	//triggerRoutines()
	//personasCargadas:= arrayDeCosas();
	//createStruct()

	//fmt.Println(personasCargadas)
	//arrayDeCosas()
	//crearUnSlideQueEsUnArrayDinamico()

	/*for i,v := range personasCargadas {
		fmt.Println(i)
		fmt.Println(v)
	}*/
}

func linkTaskToUser(task *fakeElements.Task, user *fakeElements.User) {
	
}

func crearUnSlideQueEsUnArrayDinamico() {
	type Perro struct {
		raza string
		size string
	}

	perros:= []Perro{
		{"Doberman","large"},
		{"Mestizo","medium"},
		{"Toy Minimal Dog series","small"},
	}

	fmt.Printf("%+v",perros)

	type Pez struct {
		nombreVulgar string
		nombreCientifico string
	}

	peces:= []Pez{
		{"sábalo","prochilodus lineatus"},
		{"raya","rayitus peligrussus"},
	}

	fmt.Printf("%+v",peces)

}

func arrayDeCosas () []Persona{
	var personas [8]Persona

	personas[0].id = 1
	personas[0].nombre = "Marie"
	personas[0].apellido = "Curie"

	personas[1].id = 2
	personas[1].nombre = "José"
	personas[1].apellido = "Von Neumman"

	personas[2].id = 3
	personas[2].nombre = "Luis"
	personas[2].apellido = "Olivetti"

	personas[3].id = 4
	personas[3].nombre = "Albert"
	personas[3].apellido = "Einstein"

	personas[4].id = 5
	personas[4].nombre = "Priscila Guadalupe"
	personas[4].apellido = "Diaz"

	personas[5].id = 6
	personas[5].nombre = "Claribel"
	personas[5].apellido = "Diaz Ferreyra"

	personas[6].id = 7
	personas[6].nombre = "Ariel"
	personas[6].apellido = "Kühn"

	personas[7].id = 8
	personas[7].nombre = "Lucia"
	personas[7].apellido = "Cacciabué Grecca"

	//fmt.Println(personas)

	//vamos a usar slice

	var sliceElement []Persona = personas[4:6]
	sliceElement = append(sliceElement,Persona{8,"Silvana","Reales"},
										Persona{9,"Cristian Ceferino","Lacuadra"},
										Persona{10,"Jose Luis","Migueles"},
										Persona{11,"Felipe","Stellato"})

	//fmt.Println(sliceElement)

	return sliceElement
}

func modifyData(persona *Persona) {
	//Acá vos trabajás con punteros y debés desreferenciar, es decir, sacar el value que estás apuntando
	persona.nombre = "Marie"
	persona.apellido = "Curie"
}


func createStruct () {
	persona:=Persona {1,"Juan Martin", "Del Potro"}
	fmt.Printf("%+v",persona)
	//aca vos tenés que pasarle la dirección de memoria para que persona (con *) tome el value de la dirección &persona
	modifyData(&persona);
	fmt.Printf("%+v",persona)
	//fmt.Println("id: "+fmt.Sprint(persona.id) + " Nombre: "+persona.nombre+" Apellido: " +persona.apellido)
}

func localFunction (parameter string) {
	fmt.Println("Esta es una función local. Hola "+parameter)

}

func add(x,y int) int {
	return x + y
}

func split(value int) (x,y int) {
	x = value * 4/9
	y = value - x
	return
}


func generateWelcomeString (parameter string) string {
	return "Bienvenido "+parameter
}

func mathematicFunction (x,y float64) float64 {
	status:="ok"
	defer fmt.Println(messages.PresetMessage(status))
	if x/y > 0.00000000001 {
		status="failed"
		panic("Entró en pánico")
	}
	return x/y
}

func cosasQueFuiProbando () {
	/*fmt.Println("--- Vamos a llamar a una función importada desde un paquete ---")
	messages.HelloWorldStatic();
	fmt.Println("--- Vamos a llamar a una función interna ---")
	localFunction("Ariel")*/
	//fmt.Println(generateWelcomeString("Ariel"))
	//fmt.Println(add(10,65))
	//fmt.Println(messages.HelloWorldWithParameter("Ariel Ramirez"))
	/*fmt.Println(split(17))
	var a, b ,c bool = true,true,true
	e := "mi cadena de caracteres y palabras con ancentuación"
	fmt.Println(a,b,c,e)
	var alfa int = 50
	var beta float32= float32(alfa)
	fmt.Printf("T%\n",beta)
	const ConstantValue = "this is a constant value"

	fmt.Println("Valor de la constante " +ConstantValue)

	const (
		BigNumber = 1 << 100 //mueve 100 posiciones el 1 hacia adelante (shifting)
		SmallNumber = BigNumber >> 99 //arrastra a BigNumber 99 posiciones para atrás
	)

	fmt.Println(BigNumber * 0.1)
	fmt.Println(SmallNumber *10 +1)*/
	/*
	var a,b bool = true, false
	if a == b {
		fmt.Println("Son iguales")
	} else {
		fmt.Println("Son distintos")
	}

	value:= runtime.GOOS;

	switch value {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.",value)
	}
*/
}

func triggerRoutines () {
	go printGenerator(0,200)
	printGenerator(3,200)
	//go printGenerator(1,200)
	garbageFunction()

}

func printGenerator (value int, repetitions int) {
	for i:=0 ; i<repetitions;i++ {
		fmt.Println(value)
	}
}

func garbageFunction () {
	fmt.Println("Dummy")
}
