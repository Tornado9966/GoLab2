package menu

import (
	"encoding/json"
	"github.com/Tornado9966/menu-example/server/tools"
	"log"
	"net/http"
)

// Dishes HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of dishes HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListDishes(store, rw)
		} else if r.Method == "POST" {
			handleOrderCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListDishes(store *Store, rw http.ResponseWriter) {
	res, err := store.ListDishes()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleOrderCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var of OrderForm
	if err := json.NewDecoder(r.Body).Decode(&of); err != nil {
		log.Printf("Error decoding order input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	res, err := store.CreateOrder(of.Ids, of.TableNum)
	if err == nil {
		tools.WriteJsonOk(rw, res)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}
