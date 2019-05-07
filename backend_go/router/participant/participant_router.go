package participant

import (
	"encoding/json"
	"net/http"
	"time"

	. "github.com/erick-adl/globo-bbb-wall/backend_go/aws/sqs"
	. "github.com/erick-adl/globo-bbb-wall/backend_go/models"
	"github.com/gorilla/mux"
)

var (
	start        = time.Time{}
	participants = ListOfParticipant{}
)

func pooling() {
	for {

		if time.Since(start).Seconds() > 10 {
			// fmt.Println("30s elapsed...", time.Since(start))
			start = time.Now()
			participants.Time = int(time.Since(start).Hours())
			response, _ := json.Marshal(participants)
			SendMessage(string(response))
			// fmt.Print(string(response))
			for i := range participants.Participants {
				participants.Participants[i].Votes = 0
			}
		}
	}
}
func init() {

	participants.Participants = append(participants.Participants, Participant{Name: "Joao"})
	participants.Participants = append(participants.Participants, Participant{Name: "Maria"})

	start = time.Now()
	go pooling()

}

func AddHandler(r *mux.Router) {
	r.HandleFunc("/participants/vote", Vote).Methods("POST")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Vote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var participant Participant
	if err := json.NewDecoder(r.Body).Decode(&participant); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	for i := range participants.Participants {
		if participants.Participants[i].Name == participant.Name {
			participants.Participants[i].Votes++
			respondWithJson(w, http.StatusCreated, map[string]string{"success": "registered vote"})
			return
		}
	}
	respondWithError(w, http.StatusBadRequest, "Participant not found")
}
