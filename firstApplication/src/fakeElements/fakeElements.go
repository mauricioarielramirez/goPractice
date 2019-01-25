package fakeElements

type User struct {
	id int
	name string
	alias string
}

type Task struct {
	id   int
	name string
	description string
	owner User
	share []User
}

func GetArrayUser () []User {
	var usuarios= []User{{1,"Ariel Ramirez","marramirez"}}
	users:=usuarios[0:1]
	users = append(users, User{2,"Mariano Diaz","mdiaz"},
	User{3,"Priscila Grinóvero","pgrinovero"},
	User{4,"Stefania Orzuza","sorzuza"},
	User{5,"Johana Moreira","jmoreira"})

	return users
}

func GetArrayTask () []Task {
	var tareas= []Task{{1,"Mantenimiento de backlog","Operación de orden bajo demanda",User{},[]User{}}}
	tasks:=tareas[0:1]
	tasks = append(tasks,Task{2,"Reinicio de planilla de bugs","Dejar en cero el backlog de bug y aplicar las métricas correspondientes",User{},[]User{}},
		Task{3,"Redimensionar escala de burndown","De acuerdo al progreso mensual, volver a dimensionar la escala del gráfico.",User{},[]User{}})
	return tasks
}


