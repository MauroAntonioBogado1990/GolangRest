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
			fmt.Fprintln(w, "Inserte values valids")
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
	fmt.Println(w, "Welcome to the Api");
}
//definimos la funcion de busqueda de tarea
func getTask(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
    //en el casod de existencia de error 
	if err != nil {
		fmt.Println(w, "Invalid ID")

	}
    //si esta todo OK
	for _, task := range tasks{
		if task.ID == taskID {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
	
}
}

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
	log.Fatal(http.ListenAndServe(":3000", router))
	}