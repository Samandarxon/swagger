CREATE TABLE airport(
  "id" VARCHAR PRIMARY KEY,
  "guid" UUID,
  "title" VARCHAR,
  "country_id" VARCHAR,
  "city_id" VARCHAR ,
  "latitude" VARCHAR,
  "longitude" VARCHAR,
  "radius" VARCHAR,
  "image" VARCHAR,
  "adress" VARCHAR,
  "timezone_id" UUID,
  "country" VARCHAR,
  "city" VARCHAR,
  "search_text" VARCHAR,
  "code" VARCHAR,
  "product_count" INT,
  "gmt" VARCHAR,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);

-- cp -r ../premium_plaza_lounges_plaza_object_builder_service_buildings.json ./mock_data/air.json  

jq -c '.[]' ./mock_data/air.json >> ./mock_data/your_air.json
-- CREATE TABLE temp(data JSONB);

\copy temp(data) from ./api/mock_data/your_air.json

    INSERT INTO airport(
        id,  
        guid,  
        title,  
        country_id,  
        city_id,  
        latitude,  
        longitude,  
        radius,  
        image,  
        adress,  
        timezone_id,  
        country,  
        city,  
        search_text,  
        code,  
        product_count,  
        gmt,  
        created_at,  
        updated_at
    )
    SELECT
        data ->'_id'->>'$oid',
        CAST(data ->>'guid' AS UUID),
        data ->>'title',
        data ->>'country_id',
        data ->>'city_id',
        data ->>'latitude',
        data ->>'longitude',
        data ->>'radius',
        data ->>'image',
        data ->>'adress',
        CAST(data ->>'timezone_id' AS UUID),
        data ->>'country',
        data ->>'city',
        data ->>'search_text',
        data ->>'code',
        COALESCE(CAST(data ->>'product_count' AS INT),0),
        data ->>'gmt',
        CAST(data ->'createdAt' ->> '$date' AS TIMESTAMP),
        CAST(data ->'updatedAt' ->> '$date' AS TIMESTAMP)
    FROM
      temp;

-- DROP TABLE temp;