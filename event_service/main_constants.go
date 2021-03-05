package main

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "events"
)

const (
	enableHeader    = "Access-Control-Expose-Headers"
	firstPageHeader = "First-Page"
	lastPageHeader  = "Last-Page"
)

const (
	serviceURL = "http://localhost:8002"
	secretKey  = "k8@0y%m^4-)ltn%8frs&e6^%dus1)6%s3&_u436h04)hjd6v#o"
	jwtHeader  = "Authorization"
)

const createEventsTable = `CREATE TABLE events
(
    id serial,
	active boolean NOT NULL,
    created_on text NOT NULL,
	user_id integer NOT NULL,
	product_id integer NOT NULL,
    name text NOT NULL,
	category text NOT NULL,
	"from" text NOT NULL,
	"to" text NOT NULL,
    place text NOT NULL,
	description text NOT NULL,
	CONSTRAINT events_pkey PRIMARY KEY (id)
)`

const createImagesTable = `CREATE TABLE images
(
    id serial,
	path text NOT NULL,
	event_ref integer,
	CONSTRAINT images_pkey PRIMARY KEY (id),
	CONSTRAINT foreign_key_constaint FOREIGN KEY (event_ref)
	REFERENCES events (id)
	ON UPDATE SET NULL
	ON DELETE SET NULL
)`
