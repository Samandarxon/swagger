package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"

	"github.com/google/uuid"
)

type CountryRepo struct {
	db *sql.DB
}

func NewCountryRepo(db *sql.DB) *CountryRepo {
	return &CountryRepo{
		db: db,
	}
}

func (a *CountryRepo) Create(req models.CreateCountry) (*models.Country, error) {
	var id = uuid.New().String()
	var query = `
	INSERT INTO country(
			"id",
			"guid",
			"title",
			"code",
			"continent",
			"updated_at"
) VALUES ($1,$2,$3,$4,$5,NOW())
`
	fmt.Println(query, req)
	_, err := a.db.Exec(query,
		id,
		req.Guid,
		req.Title,
		req.Code,
		req.Continent,
	)
	if err != nil {
		return &models.Country{}, err
	}
	return a.GetById(models.CountryPrimaryKey{Id: id})
}

func (c *CountryRepo) GetById(req models.CountryPrimaryKey) (*models.Country, error) {

	var (
		country = models.Country{}
		query   = `
		SELECT 
			"id",
			"guid",
			"title",
			"code",
			"continent",
			"created_at",
			"updated_at"
		FROM country WHERE id=$1`
	)
	var (
		Id        sql.NullString
		Guid      sql.NullString
		Title     sql.NullString
		Code      sql.NullString
		Continent sql.NullString
		CreatedAt sql.NullString
		UpdatedAt sql.NullString
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
		&Code,
		&Continent,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	country = models.Country{
		Id:        Id.String,
		Guid:      Guid.String,
		Title:     Title.String,
		Code:      Code.String,
		Continent: Continent.String,
		CreatedAt: CreatedAt.String,
		UpdatedAt: UpdatedAt.String,
	}
	return &country, nil
}

func (c *CountryRepo) GetList(req models.GetListCountryRequest) (*models.GetListCountryResponse, error) {
	var (
		countrys = models.GetListCountryResponse{}
		where    = " WHERE TRUE "
		offset   = " OFFSET 0"
		limit    = " LIMIT 10"
		query    = `
		SELECT 
			COUNT(*) OVER(),
			"id",
			"guid",
			"title",
			"code",
			"continent",
			"created_at",
			"updated_at"
		FROM country
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
		var country models.Country

		rows.Scan(
			&countrys.Count,
			&country.Id,
			&country.Guid,
			&country.Title,
			&country.Code,
			&country.Continent,
			&country.CreatedAt,
			&country.UpdatedAt,
		)
		countrys.Countries = append(countrys.Countries, country)
	}
	rows.Close()
	return &countrys, nil
}

func (c *CountryRepo) Update(req models.UpdateCountry) (*models.Country, error) {
	var (
		query = `
			UPDATE country SET 
				"guid" = $2,
				"title" = $3,
				"code" = $4,
				"continent" = $5,
				updated_at = NOW() 
			WHERE id = $1`
	)
	_, err := c.db.Exec(query,
		req.Id,
		req.Guid,
		req.Title,
		req.Code,
		req.Continent,
	)
	if err != nil {
		return nil, err
	}

	return c.GetById(models.CountryPrimaryKey{Id: req.Id})
}

func (c *CountryRepo) Delete(req models.CountryPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM country WHERE id = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}
