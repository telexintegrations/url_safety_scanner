package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// IntegrationSpec for Telex integration
type IntegrationSpec struct {
	Data struct {
		Author string `json:"author"`
		Date   struct {
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		} `json:"date"`
		Descriptions struct {
			AppDescription  string `json:"app_description"`
			AppLogo         string `json:"app_logo"`
			AppName         string `json:"app_name"`
			AppURL          string `json:"app_url"`
			BackgroundColor string `json:"background_color"`
		} `json:"descriptions"`

		IntegrationCategory string   `json:"integration_category"`
		IntegrationType     string   `json:"integration_type"`
		IsActive            bool     `json:"is_active"`
		KeyFeatures         []string `json:"key_features"`
		Permissions         struct {
			Events []string `json:"events"`
		} `json:"permissions"`
		Settings  []Setting `json:"settings"`
		TargetURL string    `json:"target_url"`
		TickURL   string    `json:"tick_url"`
		Website   string    `json:"website"`
	} `json:"data"`
}

type Setting struct {
	Default  bool   `json:"default"`
	Label    string `json:"label"`
	Required bool   `json:"required"`
	Type     string `json:"type"`
}

func GetIntegrationSpec(c *gin.Context) {
	spec := IntegrationSpec{}
	spec.Data.Author = "Joseph Adewunmi"
	spec.Data.Date.CreatedAt = time.Now().Format("2006-01-02")
	spec.Data.Date.UpdatedAt = time.Now().Format("2006-01-02")
	spec.Data.Descriptions.AppName = "Real-Time Scam & Phishing URL Detector"
	spec.Data.Descriptions.AppDescription = "Scans messages for URLs and classifies them as safe or suspicious."
	spec.Data.Descriptions.AppLogo = "https://picsum.photos/200/300"
	spec.Data.Descriptions.AppURL = "https://url-safety-scanner-eg1x.onrender.com/scan-url"
	spec.Data.IntegrationCategory = "Security & Compliance"
	spec.Data.Descriptions.BackgroundColor = "#ffffff"
	spec.Data.IntegrationType = "modifier"
	spec.Data.IsActive = true
	spec.Data.Settings = []Setting{
		{
			Default:  true,
			Label:    "Enable URL Scanning",
			Required: false,
			Type:     "text",
		},
	}
	spec.Data.Permissions.Events = []string{"Receive messages from Telex channels", "Scan for URLs", "Send scan results"}
	spec.Data.KeyFeatures = []string{"Receive messages from Telex channels", "Scan for URLs", "Send scan results"}
	spec.Data.TargetURL = "https://url-safety-scanner-eg1x.onrender.com/scan-url"
	spec.Data.TickURL = "https://url-safety-scanner-eg1x.onrender.com/scan-url"
	spec.Data.Website = "https://telex.im"

	c.JSON(http.StatusOK, spec)
}
