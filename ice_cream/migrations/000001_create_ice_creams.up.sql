CREATE TABLE IF NOT EXISTS ice_creams (
	id BYTEA PRIMARY KEY,
	name VARCHAR UNIQUE NOT NULL,
	image_closed VARCHAR NOT NULL,
	image_open VARCHAR NOT NULL,
	description VARCHAR NOT NULL,
	story VARCHAR NOT NULL,
	sourcing_values VARCHAR[],
	ingredients VARCHAR[],
	allergy_info VARCHAR,
	dietary_certifications VARCHAR,
	product_id VARCHAR,
	created_by vARCHAR,
	updated_by vARCHAR,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ice_cream_sourcing_values_idx on ice_creams USING GIN (sourcing_values);
CREATE INDEX ice_cream_ingredients_idx on ice_creams USING GIN (ingredients);