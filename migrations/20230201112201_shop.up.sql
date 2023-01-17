CREATE TABLE public.partner (
	id bigserial NOT NULL,
	create_time timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	full_name text NOT NULL DEFAULT 'new partner',
	url text NOT NULL,
	api_token text NOT NULL DEFAULT '',
	role int2 NOT NULL DEFAULT 0
);

ALTER TABLE public.partner OWNER TO shops_user;

CREATE TABLE public.shop (
	id bigserial NOT NULL,
	create_time timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	full_name text NOT NULL DEFAULT 'new shop',
	merchant_id bigint NOT NULL,
	login text NOT NULL,
	password text NOT NULL,
	url text NOT NULL,
	settings jsonb NOT NULL DEFAULT '{}'::jsonb,
	CONSTRAINT shop_pk PRIMARY KEY (id)
);

ALTER TABLE public.shop OWNER TO shops_user;

CREATE TABLE public.merchant (
	id bigserial NOT NULL,
	create_time timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	full_name text NOT NULL,
	url text NOT NULL,
	partner_id bigint NOT NULL,
	CONSTRAINT merchants_pk PRIMARY KEY (id)
);

ALTER TABLE public.merchant OWNER TO shops_user;


