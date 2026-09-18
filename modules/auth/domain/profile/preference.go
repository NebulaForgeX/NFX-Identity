package profile

import (
	"nfxidentity/enums"
	"nfxidentity/pkgs/jsonx"

	"gorm.io/datatypes"
)

// Theme mirrors nfx-ui ResolvedThemePreference.
type Theme struct {
	Accent          string `json:"accent"`
	Gray            string `json:"gray"`
	Appearance      string `json:"appearance"`
	Radius          string `json:"radius"`
	Scaling         string `json:"scaling"`
	PanelBackground string `json:"panelBackground"`
	FontFamily      string `json:"fontFamily"`
}

// Preference is the canonical profile preference JSON stored on auth profiles.
// Shape matches nfx-ui ResolvedPreference (language + theme + layoutMode + dashboardBackground).
type Preference struct {
	Language            string `json:"language"`
	Theme               Theme  `json:"theme"`
	LayoutMode          string `json:"layoutMode"`
	DashboardBackground string `json:"dashboardBackground"`
}

func defaultPreference(lang enums.AuthProfileLanguage) Preference {
	if lang == "" {
		lang = enums.AuthProfileLanguageEn
	}
	return Preference{
		Language: string(lang),
		Theme: Theme{
			Accent:          "tomato",
			Gray:            "slate",
			Appearance:      "system",
			Radius:          "medium",
			Scaling:         "100%",
			PanelBackground: "translucent",
			FontFamily:      "system",
		},
		LayoutMode:          "show",
		DashboardBackground: "none",
	}
}

const fallbackJSON = `{"language":"en","theme":{"accent":"tomato","gray":"slate","appearance":"system","radius":"medium","scaling":"100%","panelBackground":"translucent","fontFamily":"system"},"layoutMode":"show","dashboardBackground":"none"}`

func marshalPreference(p Preference) *datatypes.JSON {
	j := jsonx.MarshalOr(p, datatypes.JSON([]byte(fallbackJSON)))
	return &j
}

// Default returns the full initial preference JSON for a new profile.
func Default(lang enums.AuthProfileLanguage) *datatypes.JSON {
	return marshalPreference(defaultPreference(lang))
}
