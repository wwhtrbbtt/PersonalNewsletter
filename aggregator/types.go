package aggregator

type Button struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

// Basically the same as a button, but a Module can contain multiple. They are below the text.
type Link struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type Image struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	X      int64  `json:"x"`
	Y      int64  `json:"y"`
}

// The output of a module. Several of these make up an email
type Module struct {
	Caption string `json:"caption"`
	Image   Image  `json:"image"`
	Text    string `json:"text"`
	Links   []Link `json:"links"`
	Button  Button `json:"button"`
}
