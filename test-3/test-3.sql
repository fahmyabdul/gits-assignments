CREATE TABLE public.t_author (
	id bigserial NOT NULL,
	"name" varchar NULL,
	detail varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT t_author_pk PRIMARY KEY (id)
);

CREATE TABLE public.t_book (
	id bigserial NOT NULL,
	"name" varchar NULL,
	pages int4 NULL,
	publisher_id int4 NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT t_book_pk PRIMARY KEY (id),
	CONSTRAINT t_book_fk FOREIGN KEY (publisher_id) REFERENCES public.t_publisher(id)
);

CREATE TABLE public.t_book_author (
	book_id int4 NOT NULL,
	author_id int4 NOT NULL,
	CONSTRAINT t_book_author_pk PRIMARY KEY (book_id)
);

CREATE TABLE public.t_publisher (
	id bigserial NOT NULL,
	"name" varchar NULL,
	detail varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT t_publisher_pk PRIMARY KEY (id)
);