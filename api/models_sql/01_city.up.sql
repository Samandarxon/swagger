CREATE TABLE city(
    "id" VARCHAR PRIMARY KEY,
    "guid" UUID,
    "title" VARCHAR(48),
    "country_id" UUID,
    "city_code" VARCHAR(120),
    "latitude" VARCHAR(120),
    "longitude" VARCHAR(120),
    "offset" VARCHAR(120),
    "timezone_id" UUID,
    "country_name" VARCHAR(128),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);




-- jq -c '.[]' city.json >> your_data.json

-- CREATE TABLE temp(data JSONB);

\copy temp(data) from ./api/mock_data/your_cities.json;

DELETE FROM temp WHERE length(data ->> 'country_id') = 2;

INSERT INTO city (
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
)
SELECT
  data ->'_id'->>'$oid',
  CAST(data ->> 'guid' AS UUID),
  data ->> 'title',
  CAST(data ->> 'country_id' AS UUID),
  data ->> 'city_code',
  data ->> 'latitude',
  data ->> 'longitude',
  data ->> 'offset',
  CAST(data ->> 'timezone_id' AS UUID),
  data ->> 'country_name',
  CAST(data->'createdAt'->>'$date' AS TIMESTAMP),
  CAST(data->'updatedAt'->>'$date' AS TIMESTAMP)
FROM
  temp;

DROP TABLE temp;


