package main

func GetAllModules() Config {
	var c Config

	/// rss feed (count, url) ///
	var rss ModuleConfig
	rss.Name = "rss-feed"
	AddSetting(&rss, "URL", "")
	AddSetting(&rss, "Count", 1)
	c.Modules = append(c.Modules, rss)

	/// QOTD (no settings) ///
	var qotd ModuleConfig
	qotd.Name = "qotd"
	c.Modules = append(c.Modules, rss)

	return c
}

func AddSetting(m *ModuleConfig, name, value interface{}) {
	var setting ModuleConfigSetting
	setting.Value = value
	m.Settings = append(m.Settings, setting)
}
