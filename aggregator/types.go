package aggregator

type Button struct {
	Text string `json:"text"` // Button text
	URL  string `json:"url"`  // Button URL
}

// Basically the same as a button, but a Module can contain multiple. They are below the text.
type Link struct {
	Text string `json:"text"` // Link text
	URL  string `json:"url"`  // Link URL
}

type Image struct {
	Text string `json:"text"`   // Image text
	URL  string `json:"source"` // Image source URL
	X    int64  `json:"x"`      // Image X size
	Y    int64  `json:"y"`      // Image Y size
}

// The output of a module. Several of these make up an email
type Module struct {
	Caption string `json:"caption"`            // Title of the module
	Image   Image  `json:"image"`              // Image of the module - optional
	Text    string `json:"text"`               // Text of the module - optional
	Links   []Link `json:"links"`              // Collection of links - optional (ex. RSS feed)
	Button  Button `json:"button",omitonempty` // Button of the module - optional
}
