package sender

type Feed struct {
	Time         string   `json:"time"`         // ex. "10:00"
	Feedname     string   `json:"feedName"`     // ex. "My email feed"
	Greetingname string   `json:"greetingName"` // ex. "peet"
	Email        string   `json:"email"`        // ex. peet@peet.ws
	Modules      []Module `json:"modules"`      // collection of modules
}
