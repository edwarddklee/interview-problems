package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// PROBLEM DESCRIPTION:
// Your goal here is to design an API that allows for hero tracking, much like the Vue problem
// You are to implement an API (for which the skeleton already exists) that has the following capabilities
// - Get      : return a JSON representation of the hero with the name supplied
// - Make     : create a superhero according to the JSON body supplied
// - Calamity : a calamity of the supplied level requires heroes with an equivalent combined powerlevel to address it.
//              Takes a calamity with powerlevel and at least 1 hero. On success return a 200 with json response indicating the calamity has been resolved.
//              Otherwise return a response indicating that the heroes were not up to the task. Addressing a calamity adds 1 point of exhaustion.
// - Rest     : recover 1 point of exhaustion
// - Retire   : retire a superhero, someone may take up his name for the future passing on the title
// - Kill     : a superhero has passed away, his name may not be taken up again.

// On success all endpoints should return a status code 200.

// If a hero reaches an exhaustion level of maxExhaustion then they die.

// You are free to decide what your API endpoints should be called and what shape they should take. You can modify any code in this file however you'd like.

// NOTE: you may want to install postman or another request generating software of your choosing to make testing easier. (api is running on localhost port 8081)

// NOTE the second: the API is receiving asynchronous requests to manage our super friends. As such, your hero access should be thread-safe for writes.
// Even if the operations are extremely lightweight we want to make our application scalable.

// NOTE the third: There are many ways to make whatever package-level tracking you implement thread-safe, feel free to change anything about this file (without changing the functionality of the program) to do so.
// i.e. add package-level maps, add functions that take the hero struct as reference, modify the hero struct, creating access control paradigms
// I highly recommend looking into channels, mutexes, and other golang memory management patterns and pick whatever you're most comfortable with.
// For mad props: a timeout on the memory management process which returns a resource not available.

// Bonus: If you're having fun (this is by no means necessary) you can make the calamity hold the heroes up for a time and delay their unlocking in a go-routine
// example:
// go func(h *hero) {
//     time.Sleep(calamityTime)
//     // release lock on hero
// }(heroPtr)

type hero struct {
	Name       string `json:"Name"`
	PowerLevel int    `json:"PowerLevel"`
	Exhaustion int    `json:"Exhaustion"`
	// TODO: changeme?
}

type villain struct {
	Name       string `json:"Name"`
	PowerLevel int    `json:"PowerLevel"`
}

type testStruct struct {
	Test string
}

//calamity format
type showdown struct {
	VillainName string `json:"villainName"`
	HeroesName  string `json:"heroesName"`
	Result      string `json:"Result"`
}

var lock sync.Mutex
var maxExhaustion = 10
var allHeroes []hero
var allVillains []villain
var killedNames []string

//function to get json of all heroes
func getAllHeroes(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allHeroes)
	time.Sleep(1 * time.Second)
}

//get specific hero specificed
func heroGet(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	w.Header().Set("Content-Type", "application/json")
	// TODO: access the global tracking to return the hero object
	var name string
	var ok bool
	if name, ok = mux.Vars(r)["Name"]; !ok {
		log.Println("Hero does not exist in Hero database")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Hero does not exist in Hero database!"))
	}
	for _, item := range allHeroes {
		if item.Name == name {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	time.Sleep(1 * time.Second) //sleep for thread safety

	//_ = name // TODO: something with name
	//w.WriteHeader(http.StatusNotImplemented)
	//json.NewEncoder(w).Encode(&hero{})
}

//create hero (can test using put in Postman)
func heroMake(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	// TODO : create a hero and add to some sort of global tracking
	w.Header().Set("Content-Type", "application/json")
	var newHero hero
	if stringInSlice(newHero.Name, killedNames) {
		log.Println("Hero was killed, name cannot be used")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Hero was killed, name cannot be used!"))
	} else {
		_ = json.NewDecoder(r.Body).Decode(&newHero)
		allHeroes = append(allHeroes, newHero)
		json.NewEncoder(w).Encode(newHero)
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var t testStruct
	err = json.Unmarshal(content, &t)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	r.Body.Close()
	time.Sleep(1 * time.Second)

	//_ = err     // TODO handle read error
	//w.WriteHeader(http.StatusNotImplemented)
}

//step 3 of implementation, addressing calamities
func calamity(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	w.Header().Set("Content-Type", "application/json")
	var fight showdown
	_ = json.NewDecoder(r.Body).Decode(&fight)
	rand.Seed(time.Now().Unix()) //pseudorandom generator
	randVillain := allVillains[rand.Intn(len(allVillains))]
	numHeroes := rand.Intn(len(allHeroes))
	s := make([]string, numHeroes)
	t := make([]int, numHeroes)
	u := make([]hero, numHeroes)
	for i := 0; i < numHeroes; i++ {
		randHero := allHeroes[i]
		s = append(s, randHero.Name)
		t = append(t, randHero.PowerLevel)
		u = append(u, randHero)
	}
	for q := 0; q < len(u); q++ {
		for _, item := range allHeroes {
			if item.Name == s[q] {
				item.Exhaustion++
				heroLimit := item.Exhaustion
				if heroLimit == 10 {
					killedNames = append(killedNames, s[q])
					for x, v := range allHeroes {
						if v == u[q] {
							allHeroes = append(allHeroes[:x], allHeroes[x+1:]...)
							break
						}
					}
				}
			}
		}

	}
	var heroNames string
	var sum int
	//joining hero names together
	for j := 0; j < len(t); j++ {
		heroNames = heroNames + " " + s[j]
		sum = sum + t[j]
	}
	//calamity of villain vs heroes added
	fight.VillainName = randVillain.Name
	fight.HeroesName = heroNames

	//win if heroes power is greater than villains
	if randVillain.PowerLevel < sum {
		heroesWin := "The Heroes emerged victorious!"
		fight.Result = heroesWin
	} else {
		heroesLose := "The Heroes lost :(" //else lose
		fight.Result = heroesLose
	}
	json.NewEncoder(w).Encode(fight)
	time.Sleep(1 * time.Second) //sleep for thread safety
}

//rests exhaustion
func heroRest(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	w.Header().Set("Content-Type", "application/json")
	for i := 0; i < len(allHeroes); i++ {
		allHeroes[i].Exhaustion--
	}
	json.NewEncoder(w).Encode(allHeroes)
	time.Sleep(1 * time.Second)

}

//retires hero of choice by name
func retireHero(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range allHeroes {
		if item.Name == params["name"] {
			allHeroes = append(allHeroes[:i], allHeroes[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(allHeroes)
	time.Sleep(1 * time.Second)
}

//to see if name is part of killed
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
func linkRoutes(r *mux.Router) {
	r.HandleFunc("/hero", heroMake).Methods("POST")
	r.HandleFunc("/hero/{name}", heroGet).Methods("GET")
	r.HandleFunc("/hero", getAllHeroes).Methods("GET")
	r.HandleFunc("/calamity", calamity).Methods("POST")
	r.HandleFunc("/hero/{name}", retireHero).Methods("DELETE")
	r.HandleFunc("/rest", heroRest).Methods("GET")
	// TODO: add more routes
}

func initHeroDB() {
	allHeroes = append(allHeroes, hero{Name: "Groot", PowerLevel: 50, Exhaustion: 9})
	allHeroes = append(allHeroes, hero{Name: "Captain-America", PowerLevel: 60, Exhaustion: 0})
	allHeroes = append(allHeroes, hero{Name: "Thor", PowerLevel: 200, Exhaustion: 0})
	allHeroes = append(allHeroes, hero{Name: "Dr-Strange", PowerLevel: 150, Exhaustion: 0})
	allHeroes = append(allHeroes, hero{Name: "Spiderman", PowerLevel: 30, Exhaustion: 0})
	allHeroes = append(allHeroes, hero{Name: "Black-Widow", PowerLevel: 20, Exhaustion: 0})
	allHeroes = append(allHeroes, hero{Name: "Goku", PowerLevel: 9001, Exhaustion: 0})

}

func initVillainsDB() {
	allVillains = append(allVillains, villain{Name: "Iron-Monger", PowerLevel: 20})
	allVillains = append(allVillains, villain{Name: "Ultron", PowerLevel: 250})
	allVillains = append(allVillains, villain{Name: "Thanos", PowerLevel: 350})
	allVillains = append(allVillains, villain{Name: "Dormammu", PowerLevel: 9000})
	allVillains = append(allVillains, villain{Name: "Galactus", PowerLevel: 500})
	allVillains = append(allVillains, villain{Name: "Killmonger", PowerLevel: 50})

}

func main() {
	// create a router
	router := mux.NewRouter()

	//initial "database" of heroes and villians
	initHeroDB()
	initVillainsDB()

	//rip Iron Man
	killedNames = append(killedNames, "Iron-Man")

	// and a server to listen on port 8081
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8081),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// link the supplied routes
	linkRoutes(router)
	// wait for requests
	log.Fatal(server.ListenAndServe())
}
