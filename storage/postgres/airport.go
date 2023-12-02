package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"

	"github.com/google/uuid"
)

type AirportRepo struct {
	db *sql.DB
}

func NewAirportRepo(db *sql.DB) *AirportRepo {
	return &AirportRepo{
		db: db,
	}
}

func (a *AirportRepo) Create(req models.CreateAirport) (*models.Airport, error) {
	var id = uuid.New().String()
	var query = `
	INSERT INTO airport(
			"id",
			"title",
			"country_id",
			"city_id",
			"latitude",
			"longitude",
			"radius",
			"image",
			"address",
			"timezone_id",
			"country",
			"city",
			"search_text",
			"code",
			"product_count",
			"gmt",
			"updated_at"
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,NOW())
`
	fmt.Println(query, req)
	_, err := a.db.Exec(query,
		id,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Latitude,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Address,
		req.TimezoneId,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
	)
	if err != nil {
		return &models.Airport{}, err
	}
	return a.GetById(models.AirportPrimaryKey{Id: id})
}

func (c *AirportRepo) GetById(req models.AirportPrimaryKey) (*models.Airport, error) {

	var (
		airport = models.Airport{}
		query   = `
		SELECT 
			"id",
			"title",
			"country_id",
			"city_id",
			"latitude",
			"longitude",
			"radius",
			"image",
			"address",
			"timezone_id",
			"country",
			"city",
			"search_text",
			"code",
			"product_count",
			"gmt",
			"created_at",
			"updated_at"
		FROM airport WHERE id=$1`
	)
	var (
		Id           sql.NullString
		Title        sql.NullString
		CountryId    sql.NullString
		CityId       sql.NullString
		Latitude     sql.NullString
		Longitude    sql.NullString
		Radius       sql.NullString
		Image        sql.NullString
		Address      sql.NullString
		TimezoneId   sql.NullString
		Country      sql.NullString
		City         sql.NullString
		SearchText   sql.NullString
		Code         sql.NullString
		ProductCount sql.NullInt64
		Gmt          sql.NullString
		CreatedAt    sql.NullString
		UpdatedAt    sql.NullString
	)
	fmt.Println(query)
	resp := c.db.QueryRow(query, req.Id)
	fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&Title,
		&CountryId,
		&CityId,
		&Latitude,
		&Longitude,
		&Radius,
		&Image,
		&Address,
		&TimezoneId,
		&Country,
		&City,
		&SearchText,
		&Code,
		&ProductCount,
		&Gmt,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	airport = models.Airport{
		Id:           Id.String,
		Title:        Title.String,
		CountryId:    CountryId.String,
		CityId:       CityId.String,
		Latitude:     Latitude.String,
		Longitude:    Longitude.String,
		Radius:       Radius.String,
		Image:        Image.String,
		Address:      Address.String,
		TimezoneId:   TimezoneId.String,
		Country:      Country.String,
		City:         City.String,
		SearchText:   SearchText.String,
		Code:         Code.String,
		ProductCount: int(ProductCount.Int64),
		Gmt:          Gmt.String,
		CreatedAt:    CreatedAt.String,
		UpdatedAt:    UpdatedAt.String,
	}
	return &airport, nil
}

func (c *AirportRepo) GetList(req models.GetListAirportRequest) (*models.GetListAirportResponse, error) {
	var (
		airports = models.GetListAirportResponse{}

		Id           sql.NullString
		Title        sql.NullString
		CountryId    sql.NullString
		CityId       sql.NullString
		Latitude     sql.NullString
		Longitude    sql.NullString
		Radius       sql.NullString
		Image        sql.NullString
		Address      sql.NullString
		TimezoneId   sql.NullString
		Country      sql.NullString
		City         sql.NullString
		SearchText   sql.NullString
		Code         sql.NullString
		ProductCount sql.NullInt64
		Gmt          sql.NullString
		CreatedAt    sql.NullString
		UpdatedAt    sql.NullString

		where  = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		query  = `
				SELECT 
					COUNT(*) OVER(),
					"id",
					"title",
					"country_id",
					"city_id",
					"latitude",
					"longitude",
					"radius",
					"image",
					"address",
					"timezone_id",
					"country",
					"city",
					"search_text",
					"code",
					"product_count",
					"gmt",
					"created_at",
					"updated_at"
				FROM airport`
	)
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += where + offset + limit

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}
	fmt.Println(query)
	for rows.Next() {
		var airport models.Airport

		err := rows.Scan(
			&airports.Count,
			&Id,
			&Title,
			&CountryId,
			&CityId,
			&Latitude,
			&Longitude,
			&Radius,
			&Image,
			&Address,
			&TimezoneId,
			&Country,
			&City,
			&SearchText,
			&Code,
			&ProductCount,
			&Gmt,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		airport = models.Airport{
			Id:           Id.String,
			Title:        Title.String,
			CountryId:    CountryId.String,
			CityId:       CityId.String,
			Latitude:     Latitude.String,
			Longitude:    Longitude.String,
			Radius:       Radius.String,
			Image:        Image.String,
			Address:      Address.String,
			TimezoneId:   TimezoneId.String,
			Country:      Country.String,
			City:         City.String,
			SearchText:   SearchText.String,
			Code:         Code.String,
			ProductCount: int(ProductCount.Int64),
			Gmt:          Gmt.String,
			CreatedAt:    CreatedAt.String,
			UpdatedAt:    UpdatedAt.String,
		}
		fmt.Printf("%+v\n", airport)
		airports.Airports = append(airports.Airports, airport)
	}
	defer rows.Close()
	return &airports, nil
}

func (c *AirportRepo) Update(req models.UpdateAirport) (*models.Airport, error) {
	var (
		query = `
			UPDATE airport SET
				"title"=$2,
				"country_id"=$3,
				"city_id"=$4,
				"latitude"=$5,
				"longitude"=$6,
				"radius"=$7,
				"image"=$8,
				"address"=$9,
				"timezone_id"=$10,
				"country"=$11,
				"city"=$12,
				"search_text"=$13,
				"code"=$14,
				"product_count"=$15,
				"gmt"=$16,
				updated_at = NOW() 
			WHERE id = $1`
	)
	_, err := c.db.Exec(query,
		req.Id,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Latitude,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Address,
		req.TimezoneId,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
	)
	if err != nil {
		return nil, err
	}

	return c.GetById(models.AirportPrimaryKey{Id: req.Id})
}

func (c *AirportRepo) Delete(req models.AirportPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM airport WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}
