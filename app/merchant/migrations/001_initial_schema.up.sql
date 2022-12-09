-- object: public.role | type: TYPE --
DROP TYPE IF EXISTS public.role CASCADE;
CREATE TYPE public.role AS
ENUM ('base_partner','partner','admin','guest');
-- ddl-end --
ALTER TYPE public.role OWNER TO postgres;
-- ddl-end --

-- object: public.partners | type: TABLE --
DROP TABLE IF EXISTS public.partners CASCADE;
CREATE TABLE public.partners (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	create_time timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	full_name text NOT NULL DEFAULT 'new partner',
	url text NOT NULL,
	api_token uuid NOT NULL DEFAULT gen_random_uuid(),
	role public.role NOT NULL DEFAULT 'guest',
	CONSTRAINT partners_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN public.partners.id IS E'Partner unique id';
-- ddl-end --
COMMENT ON COLUMN public.partners.create_time IS E'Timestamp of creation';
-- ddl-end --
COMMENT ON COLUMN public.partners.full_name IS E'Full name';
-- ddl-end --
COMMENT ON COLUMN public.partners.url IS E'Url';
-- ddl-end --
COMMENT ON COLUMN public.partners.api_token IS E'Partner API token';
-- ddl-end --
COMMENT ON COLUMN public.partners.role IS E'Role';
-- ddl-end --
ALTER TABLE public.partners OWNER TO postgres;
-- ddl-end --

-- object: public.shops | type: TABLE --
DROP TABLE IF EXISTS public.shops CASCADE;
CREATE TABLE public.shops (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	create_time timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	full_name text NOT NULL DEFAULT 'new shop',
	merchant_id uuid NOT NULL,
	login text NOT NULL,
	password text NOT NULL,
	url text NOT NULL,
	CONSTRAINT shops_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN public.shops.id IS E'Shop unique id';
-- ddl-end --
COMMENT ON COLUMN public.shops.create_time IS E'Timestamp of creation';
-- ddl-end --
COMMENT ON COLUMN public.shops.full_name IS E'Shop full name';
-- ddl-end --
COMMENT ON COLUMN public.shops.merchant_id IS E'Parent merchant id';
-- ddl-end --
COMMENT ON COLUMN public.shops.login IS E'Login phrase';
-- ddl-end --
COMMENT ON COLUMN public.shops.password IS E'Password';
-- ddl-end --
COMMENT ON COLUMN public.shops.url IS E'Url';
-- ddl-end --
ALTER TABLE public.shops OWNER TO postgres;
-- ddl-end --

-- object: public.merchants | type: TABLE --
DROP TABLE IF EXISTS public.merchants CASCADE;
CREATE TABLE public.merchants (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	create_time timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	full_name text NOT NULL,
	url text NOT NULL,
	partner_id uuid NOT NULL,
	CONSTRAINT merchants_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN public.merchants.id IS E'Merchant unique index';
-- ddl-end --
COMMENT ON COLUMN public.merchants.create_time IS E'Tiemestamp of creation';
-- ddl-end --
COMMENT ON COLUMN public.merchants.full_name IS E'Full name';
-- ddl-end --
COMMENT ON COLUMN public.merchants.url IS E'Url';
-- ddl-end --
COMMENT ON COLUMN public.merchants.partner_id IS E'Parent partner id';
-- ddl-end --
ALTER TABLE public.merchants OWNER TO postgres;
-- ddl-end --

-- object: public.terminals | type: TABLE --
DROP TABLE IF EXISTS public.terminals CASCADE;
CREATE TABLE public.terminals (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	create_time timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	full_name text NOT NULL,
	shop_id uuid NOT NULL,
	login text NOT NULL,
	password text NOT NULL,
	url text NOT NULL,
	token text NOT NULL,
	CONSTRAINT terminals_pk PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN public.terminals.id IS E'Terminal unique id';
-- ddl-end --
COMMENT ON COLUMN public.terminals.create_time IS E'Timestamp of creation';
-- ddl-end --
COMMENT ON COLUMN public.terminals.full_name IS E'Terminal full name';
-- ddl-end --
COMMENT ON COLUMN public.terminals.shop_id IS E'Parent shop id';
-- ddl-end --
COMMENT ON COLUMN public.terminals.login IS E'Login phrase';
-- ddl-end --
COMMENT ON COLUMN public.terminals.password IS E'Password';
-- ddl-end --
COMMENT ON COLUMN public.terminals.url IS E'Url';
-- ddl-end --
ALTER TABLE public.terminals OWNER TO postgres;
-- ddl-end --

-- object: public.schema_migrations | type: TABLE --
DROP TABLE IF EXISTS public.schema_migrations CASCADE;
CREATE TABLE public.schema_migrations (
	version bigint NOT NULL,
	dirty boolean NOT NULL DEFAULT false,
	CONSTRAINT schema_migrations_pkey PRIMARY KEY (version)
);
-- ddl-end --
ALTER TABLE public.schema_migrations OWNER TO postgres;
-- ddl-end --

-- object: shops_merchant_fk | type: CONSTRAINT --
ALTER TABLE public.shops DROP CONSTRAINT IF EXISTS shops_merchant_fk CASCADE;
ALTER TABLE public.shops ADD CONSTRAINT shops_merchant_fk FOREIGN KEY (merchant_id)
REFERENCES public.merchants (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: merchants_partner_fk | type: CONSTRAINT --
ALTER TABLE public.merchants DROP CONSTRAINT IF EXISTS merchants_partner_fk CASCADE;
ALTER TABLE public.merchants ADD CONSTRAINT merchants_partner_fk FOREIGN KEY (partner_id)
REFERENCES public.partners (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --

-- object: terminal_shop_fk | type: CONSTRAINT --
ALTER TABLE public.terminals DROP CONSTRAINT IF EXISTS terminal_shop_fk CASCADE;
ALTER TABLE public.terminals ADD CONSTRAINT terminal_shop_fk FOREIGN KEY (shop_id)
REFERENCES public.shops (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

INSERT INTO public.partners(
	id, create_time, full_name, url, api_token, role)
	VALUES ('b6f97e90-29b4-4c99-a5a5-dede07629932', '2022-12-06 17:45:15.913909+00', 'partner_1', '', 'f7c92f94-a3ac-40f8-a74f-c7d03247a698', 'base_partner');
	
INSERT INTO public.merchants(
	id, create_time, full_name, url, partner_id)
	VALUES ('57967ad6-8b75-4362-9d79-c4a0f4b42d84', '2022-12-06 17:45:34.011838+00', 'merchant_1', '', 'b6f97e90-29b4-4c99-a5a5-dede07629932');
	
INSERT INTO public.shops(
	id, create_time, full_name, merchant_id, login, password, url)
	VALUES ('4f63e896-d776-4511-9b4c-d98753a22fdb', '2022-12-06 17:45:52.115621+00', 'shop_1', '57967ad6-8b75-4362-9d79-c4a0f4b42d84', '', '', '');

INSERT INTO public.terminals(
	id, create_time, full_name, shop_id, login, password, url, token)
	VALUES ('0cf51482-7a1c-48c9-9d41-a6053da3ac89', '2022-12-06 17:52:56.482736+00', 'terminal_1', '4f63e896-d776-4511-9b4c-d98753a22fdb', 'user', '1234', '', '');
	
INSERT INTO public.schema_migrations(
	version, dirty)
	VALUES (1, false);
	
	