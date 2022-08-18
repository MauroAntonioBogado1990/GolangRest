package main

import ("fmt"
	    "net/http"
		"log"
		"github.com/gorilla/mux"
		"encoding/json"
		"io/ioutil"
		"strconv"
	    )

//Definimos una tarea

type task struct {
	ID int `json:"ID"`	 
	Name string `json:"Name"` 
	Price float64 `json:"Price"`
	Sizes string `json:"Sizes"`
	Content string `json:"Content"`
}

//Creamos una arra de tareas, para el caso de acciones futuras
type allTasks []task

var tasks = allTasks {
	{
		ID: 1,
		Name: "Article",
		Price: 0.0,
		Sizes: "M",
		Content: "Pant with draw of color blue",
	},
}

//Definimos la funcion para retornar las tareas
func getTasks(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(tasks)
}
//Definimos la funcion para crear las tareas
func createTask(w http.ResponseWriter, r *http.Request){
	var newTask task;
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
			fmt.Fprintf(w, "Inserte valores validos")
	}

	json.Unmarshal(reqBody, &newTask)
    //sumamos un valor incremental al ID
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

//creamos la funcion de llamada a Indice de pagina 
func indexRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Bienvenido a la API");
}
//definimos la funcion de busqueda de tarea
func getTask(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
    //en el casod de existencia de error 
	if err != nil {
		fmt.Fprintf(w, " ID Invalido")

	}
    //si esta todo OK
	for _, task := range tasks{
		if task.ID == taskID {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
	
}
}
//definimos una funcion para la eliminacion de tareas
func deleteTask(w http.ResponseWriter, r *http.Request) { 
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err!=nil {
		fmt.Fprintf(w, " ID Invalido")	
	}
	//Si esta to OK 
	for i, task := range tasks{
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "Eliminando Tareas de %d", task.ID)

	}
	}
}
//funcione para actualizar tareas
func updateTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	//guaradamos en una variable la tarea a actualizar
	var updateTask task
	
	if err!=nil {
		fmt.Fprintf(w, " ID Invalido")
	}
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Por favor colocar valores validos")
	} 

	json.Unmarshal(reqBody, &updateTask)

	for i, task := range tasks{
			if task.ID == taskID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				updateTask.ID = taskID
				tasks = append(tasks, updateTask)

				fmt.Fprintf(w, "Task %d updated", taskID)							
			}
	}
	
}
//funcion main principal 
func main() {
	//establecemos el modo estricto de enrutar, con la var router
	router := mux.NewRouter().StrictSlash(true)
	//Indice
	router.HandleFunc("/", indexRoute )	
	//Ruta get
	router.HandleFunc("/tasks", getTasks).Methods("GET")//retorna las tareas
	//ruta Post
	router.HandleFunc("/tasks", createTask).Methods("POST")
	//ruta de busqueda por ID
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	//ruta para elimacion de tareas
	router.HandleFunc("/deleteTasks/{id}", deleteTask).Methods("DELETE")
	//ruta para la actualizacion de la tareas
	router.HandleFunc("/updateTask/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
	}