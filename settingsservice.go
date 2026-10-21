package main

// Settings is the startup configuration supplied by App Builder.
type Settings struct {
	Project  string `json:"project"`
	Theme    string `json:"theme"`
	Language string `json:"language"`
}

// SettingsService makes configuration available app-wide through Wails bindings.
// The theming and localization features can call GetSettings to read it.
type SettingsService struct {
	settings Settings
}

func (s *SettingsService) GetSettings() Settings {
	return s.settings
}
