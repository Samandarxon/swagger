CREATE TABLE country(
    "id" VARCHAR PRIMARY KEY
    "guid" UUID ,
    "title" VARCHAR(64),
    "code" VARCHAR(12),
    "continent" VARCHAR(12),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


-- cp -r ../premium_plaza_lounges_plaza_object_builder_service_countries.json  ./mock_data/countries.json


jq -c '.[]' ./mock_data/countries.json > ./mock_data/your_countries.json 

CREATE TABLE temp(data JSONB);

\copy temp(data) from ./api/mock_data/your_countries.json;


  INSERT INTO country(
      "id",
      "guid",
      "title",
      "code",
      "continent",
      "created_at",
      "updated_at"
  )
  SELECT 
    data->'_id' ->> '$oid',
    CAST(data->>'guid' AS  UUID),
    data->>'title',
    data->>'code',
    data->>'continent',
    CAST(data->'createdAt'->>'$date' AS TIMESTAMP),
    CAST(data->'updatedAt'->>'$date' AS TIMESTAMP)
  FROM 
    temp;

DROP TABLE temp;



    