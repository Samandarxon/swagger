package models

type City struct {
	Id          string `json:"id"`
	Guid        string `json:"guid"`
	Title       string `json:"title"`
	CountryId   string `json:"country_id"`
	CityCode    string `json:"city_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offset      string `json:"offset"`
	TimezoneId  string `json:"timezone_id"`
	CountryName string `json:"country_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateCity struct {
	Guid        string `json:"guid"`
	Title       string `json:"title"`
	CountryId   string `json:"country_id"`
	CityCode    string `json:"city_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offset      string `json:"offset"`
	TimezoneId  string `json:"timezone_id"`
	CountryName string `json:"country_name"`
}

type UpdateCity struct {
	Id          string `json:"-"`
	Guid        string `json:"guid"`
	Title       string `json:"title"`
	CountryId   string `json:"country_id"`
	CityCode    string `json:"city_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offset      string `json:"offset"`
	TimezoneId  string `json:"timezone_id"`
	CountryName string `json:"country_name"`
}

type CityPrimaryKey struct {
	Id string `json:"id"`
}

type GetListCityRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListCityResponse struct {
	Count  int    `json:"count"`
	Cities []City `json:"cities"`
}
