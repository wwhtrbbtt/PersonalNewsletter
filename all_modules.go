package main

func GetAllModules() Config {
	var c Config

	/// rss feed /// (count, url)
	var rss ModuleConfig
	rss.Name = "rss-feed"
	AddSetting(&rss, "URL", "http://feeds.bbci.co.uk/news/technology/rss.xml")
	AddSetting(&rss, "count", "5")
	c.Modules = append(c.Modules, rss)

	return c
}

func AddSetting(m *ModuleConfig, name, value string) {
	var setting ModuleConfigSetting
	setting.Name = name
	setting.Value = value
	m.Settings = append(m.Settings, setting)
}
