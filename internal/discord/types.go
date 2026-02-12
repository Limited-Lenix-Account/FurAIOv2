package discord

import (
	"net/url"
	"time"
)

type PaDiscordData struct {
	WebhookURL string

	ProductName  string
	ProductImage string
	ProductVol   string
	Subtotal     string
	OrderNumber  string

	ProfileEmail string
	ProfileName  string
	Fulfillment  string

	CheckoutTime float64
	TaskTime     float64
	Mode         string

	DeclineReason string
}

type SuccessfulHotelData struct {
	WebhookURL string

	EventName string
	StartDate string
	EndDate   string

	HotelName  string
	HotelImage string
	BlockName  string

	TotalCost    string
	ProfileEmail string
	AckNumber    string
	OrderLink    url.URL

	CheckoutTime float64
}

type HotelMonitorData struct {
	HotelName string

	BlockImage string
	BlockName  string
	BlockDates string
	BlockCost  string
	Quantity   string
}

type LiquorMonitorData struct {
	Name   string
	Price  string
	Method string // Pickup or shipping (if applicable)

	LastModified time.Time
	PageURL      string
}

type WebhookMessage struct {
	Content string          `json:"content,omitempty"`
	Embeds  []*WebhookEmbed `json:"embeds,omitempty"`
}

type WebhookEmbed struct {
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	Fields      []*WebhookField `json:"fields,omitempty"`
	Color       int             `json:"color,omitempty"`
	Author      Author          `json:"author,omitempty"`
}

type WebhookField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
