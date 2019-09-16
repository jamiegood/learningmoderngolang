package dinowebportal

import (
	"dino/databaselayer"
	"dino/dinowebportal/dinoTemplate"
	"dino/dinowebportal/dinoapi"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"log"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// DataFeedMessage ...
type DataFeedMessage struct {
	Heartrate     int
	Bloodpressure int
}

//var cookieStore = session.NewCookieStore([]byte("somtheing-very-secret"))
var cookieStore = sessions.NewCookieStore([]byte("fdsfdsfdsfdsfds"))

const (
	INPUTNAME     = "inputname" // inputname this is a comment
	SIGNINSESSION = "signinsession"
	USERNAME      = "username"
)

//RunWebPortal starts running the dino web portal on address addr
func RunWebPortal(dbtype uint8, addr, dbconnection, frontend string) error {

	rand.Seed(time.Now().UTC().UnixNano())

	r := mux.NewRouter()

	db, err := databaselayer.GetDatabaseHandler(dbtype, dbconnection)

	if err != nil {
		fmt.Println(err)

		return err

	}
	fmt.Println("we got this far")

	dinoapi.RunAPIOnRouter(r, db)

	fmt.Println("we got this far")
	r.Path("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		session, err := cookieStore.Get(req, SIGNINSESSION)
		if err != nil {
			return
		}
		fmt.Println(session.Values)
		val, ok := session.Values[USERNAME]

		if !ok {
			dinoTemplate.HandleSignUp(w)
			return
		}

		name, ok := val.(string)
		if !ok {
			dinoTemplate.HandleSignUp(w)
			return
		}
		//

		dinoTemplate.Homepage("Dino Portal", fmt.Sprintf("Welcome %s, where you can find metrics and information", name), w)
	})

	r.Path("/signup/").Methods("POST").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			return
		}
		namelist := req.Form[INPUTNAME]
		session, err := cookieStore.Get(req, SIGNINSESSION)
		if err != nil {
			return
		}

		if len(namelist) == 0 {
			return
		}
		session.Values[USERNAME] = namelist[0]
		session.Save(req, w)
		dinoTemplate.Homepage("Dino Portal", fmt.Sprintf("Welcome %s, where you can find metrics and information", namelist[0]), w)

	})

	r.PathPrefix("/metrics/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		animals, err := db.GetAvailableDynos()
		if err != nil {
			return
		}
		dinoTemplate.HandleMetrics(animals, w)
	})

	r.PathPrefix("/info/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		animals, err := db.GetAvailableDynos()
		if err != nil {
			return
		}
		dinoTemplate.HandleInfo(animals, w)
	})

	fileserver := http.FileServer(http.Dir(frontend))
	r.Path("/dinodatafeed").HandlerFunc(dinoDataFeedHandler)

	r.PathPrefix("/").Handler(fileserver)
	return http.ListenAndServe(addr, r)
}

func dinoDataFeedHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("inside datafeed")

	conn, err := upgrader.Upgrade(w, r, nil)
	fmt.Println("inside datafeed")

	if err != nil {
		log.Println("Could not establish websocket connection, error", err)
		fmt.Println("Could not establish websocket connection, error")
		fmt.Println(err)

		return
	}
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println("Could not read message from websocket, error", err)
			return
		}

		active := true
		if messageType == websocket.CloseMessage {
			log.Println("closing websocket... ")
			active = false
			break
		}

		go func(dino string) {

			for active {
				time.Sleep(1 * time.Second)
				msg := &DataFeedMessage{rand.Intn(300) + 1, rand.Intn(1000) + 300}
				//msg := dino + strconv.Itoa(rand.Intn(300)+1)
				databytes, err := json.Marshal(msg)
				if err != nil {
					log.Println("Could not convert data to JSON, error", databytes)
					return
				}
				if err = conn.WriteMessage(messageType, databytes); err != nil {
					log.Println("Could not write message to websocket, error", err)
					return
				}
			}
		}(string(p))
	}
}
