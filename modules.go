package main

type ModuleSetting struct {
	InternalName string              `json:"internalName"`
	ShowName     string              `json:"showName"`
	Description  string              `json:"description"`
	Type         string              `json:"type"`
	Data         []ModuleSettingData `json:"data"`
}

type ModuleSettingData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Module struct {
	InternalName string          `json:"internalName"`
	ShowName     string          `json:"showName"`
	Description  string          `json:"description"`
	Settings     []ModuleSetting `json:"settings"`
}

// needed - firestore doesnt like arrays
type allModules struct {
	Modules []Module `json:"modules"`
}

func getAllModules() allModules {
	var modules []Module
	var m Module

	// RSS-FEED
	m = Module{
		InternalName: "rss-feed",
		ShowName:     "RSS/Atom feed",
		Description:  "Get the first x entries of any RSS/Atom feed",
		Settings: []ModuleSetting{
			{
				InternalName: "URL",
				ShowName:     "URL",
				Description:  "The URL of your feed",
				Type:         "str",
			},
			{
				InternalName: "count",
				ShowName:     "Count",
				Description:  "How many entries should be send.",
				Type:         "int",
			},
		},
	}
	modules = append(modules, m)

	// QOTD
	m = Module{
		InternalName: "qotd",
		ShowName:     "Quote of the day",
		Description:  "Get a deep quote to think about",
	}

	return allModules{Modules: modules}
}
