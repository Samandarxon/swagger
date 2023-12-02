package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"

	"github.com/google/uuid"
)

type CityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) *CityRepo {
	return &CityRepo{
		db: db,
	}
}

func (a *CityRepo) Create(req models.CreateCity) (*models.City, error) {

	var id = uuid.New().String()
	var query = `
	INSERT INTO city(
		"id",
		"guid",
		"title",
		"country_id",
		"city_code",
		"latitude",
		"longitude",
		"offset",
		"timezone_id",
		"country_name",
		updated_at
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW())
`
	fmt.Println(query, req)
	_, err := a.db.Exec(query,
		id,
		req.Guid,
		req.Title,
		req.CountryId,
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		req.TimezoneId,
		req.CountryName,
	)
	if err != nil {
		return &models.City{}, err
	}
	return a.GetById(models.CityPrimaryKey{Id: id})
}

func (c *CityRepo) GetById(req models.CityPrimaryKey) (*models.City, error) {

	var (
		city  = models.City{}
		query = `
		SELECT 
			"id",
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			created_at,
			updated_at 
		FROM city WHERE id=$1`
	)
	var (
		Id          sql.NullString
		Guid        sql.NullString
		Title       sql.NullString
		CountryId   sql.NullString
		CityCode    sql.NullString
		Latitude    sql.NullString
		Longitude   sql.NullString
		Offset      sql.NullString
		TimezoneId  sql.NullString
		CountryName sql.NullString
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
	)
	fmt.Println(query)
	resp := c.db.QueryRow(query, req.Id)
	fmt.Println("*********************", resp)

	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&Id,
		&Guid,
		&Title,
		&CountryId,
		&CityCode,
		&Latitude,
		&Longitude,
		&Offset,
		&TimezoneId,
		&CountryName,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	city = models.City{
		Id:          Id.String,
		Guid:        Guid.String,
		Title:       Title.String,
		CountryId:   CountryId.String,
		CityCode:    CityCode.String,
		Latitude:    Latitude.String,
		Longitude:   Longitude.String,
		Offset:      Offset.String,
		TimezoneId:  TimezoneId.String,
		CountryName: CountryName.String,
		CreatedAt:   CreatedAt.String,
		UpdatedAt:   UpdatedAt.String,
	}
	return &city, nil
}

func (c *CityRepo) GetList(req models.GetListCityRequest) (*models.GetListCityResponse, error) {
	var (
		citys  = models.GetListCityResponse{}
		where  = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		guid   sql.NullString
		query  = `
		SELECT 
			COUNT(*) OVER(),
			"id",
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			"created_at",
			"updated_at"
		FROM city
		`
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
		var city models.City
		err = rows.Scan(
			&citys.Count,
			&city.Id,
			&guid,
			&city.Title,
			&city.CountryId,
			&city.CityCode,
			&city.Latitude,
			&city.Longitude,
			&city.Offset,
			&city.TimezoneId,
			&city.CountryName,
			&city.CreatedAt,
			&city.UpdatedAt,
		)
		city.Guid = guid.String
		if err != nil {
			return nil, err
		}
		fmt.Println(city.Guid)
		citys.Cities = append(citys.Cities, city)
	}
	defer rows.Close()
	return &citys, nil
}

func (c *CityRepo) Update(req models.UpdateCity) (*models.City, error) {
	var (
		query = `
			UPDATE city SET 
				guid = $2,
				title = $3,
				country_id = $4,
				city_code = $5,
				latitude = $6,
				longitude = $7,
				"offset" = $8,
				timezone_id = $9,
				country_name = $10,
				updated_at = NOW() 
			WHERE id = $1`
	)

	_, err := c.db.Exec(query,
		req.Id,
		req.Guid,
		req.Title,
		req.CountryId,
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		req.TimezoneId,
		req.CountryName,
	)

	fmt.Println("*******************", query)
	if err != nil {
		return nil, err
	}

	return c.GetById(models.CityPrimaryKey{Id: req.Id})
}

func (c *CityRepo) Delete(req models.CityPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM city WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}
