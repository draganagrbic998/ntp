package main

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "ads"
)

const (
	enableHeader    = "Access-Control-Expose-Headers"
	firstPageHeader = "First-Page"
	lastPageHeader  = "Last-Page"
	defaultPageSize = 10
)

const (
	serviceURL = "http://localhost:8001"
	secretKey  = "k8@0y%m^4-)ltn%8frs&e6^%dus1)6%s3&_u436h04)hjd6v#o"
	jwtHeader  = "Authorization"
)

const createAdvertisementsTable = `CREATE TABLE advertisements
(
    id serial,
	active boolean NOT NULL,
    created_on text NOT NULL,
	user_id integer NOT NULL,
    name text NOT NULL,
	category text NOT NULL,
    price text NOT NULL,
	description text NOT NULL,
	CONSTRAINT advertisements_pkey PRIMARY KEY (id)
)`

const createImagesTable = `CREATE TABLE images
(
    id serial,
	path text NOT NULL,
	prod_ref integer,
	CONSTRAINT images_pkey PRIMARY KEY (id),
	CONSTRAINT foreign_key_constaint FOREIGN KEY (prod_ref)
	REFERENCES advertisements (id)
	ON UPDATE SET NULL
	ON DELETE SET NULL
)`
