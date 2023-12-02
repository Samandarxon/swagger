package models

type Airport struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	CountryId    string `json:"country_id"`
	CityId       string `json:"city_id"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Radius       string `json:"radius"`
	Image        string `json:"image"`
	Address      string `json:"address"`
	TimezoneId   string `json:"timezone_id"`
	Country      string `json:"country"`
	City         string `json:"city"`
	SearchText   string `json:"search_text"`
	Code         string `json:"code"`
	ProductCount int    `json:"product_count"`
	Gmt          string `json:"gmt"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CreateAirport struct {
	Title        string `json:"title"`
	CountryId    string `json:"country_id"`
	CityId       string `json:"city_id"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Radius       string `json:"radius"`
	Image        string `json:"image"`
	Address      string `json:"address"`
	TimezoneId   string `json:"timezone_id"`
	Country      string `json:"country"`
	City         string `json:"city"`
	SearchText   string `json:"search_text"`
	Code         string `json:"code"`
	ProductCount int    `json:"product_count"`
	Gmt          string `json:"gmt"`
}

type UpdateAirport struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	CountryId    string `json:"country_id"`
	CityId       string `json:"city_id"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Radius       string `json:"radius"`
	Image        string `json:"image"`
	Address      string `json:"address"`
	TimezoneId   string `json:"timezone_id"`
	Country      string `json:"country"`
	City         string `json:"city"`
	SearchText   string `json:"search_text"`
	Code         string `json:"code"`
	ProductCount int    `json:"product_count"`
	Gmt          string `json:"gmt"`
}

type AirportPrimaryKey struct {
	Id string `json:"id"`
}

type GetListAirportRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListAirportResponse struct {
	Count    int       `json:"count"`
	Airports []Airport `json:"airports"`
}
