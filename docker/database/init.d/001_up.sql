BEGIN;

-- public.operations definition

CREATE TABLE public.operations (
	id uuid NOT NULL,
	"name" varchar(50) NOT NULL,
	"cost" int8 NOT NULL,
	CONSTRAINT operations_pkey PRIMARY KEY (id)
);


-- public.users definition

CREATE TABLE public.users (
	id uuid NOT NULL,
	username varchar(50) NOT NULL,
	"password" text NOT NULL,
	status int2 NOT NULL,
	CONSTRAINT uni_users_username UNIQUE (username),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

-- Seed public.users table

INSERT INTO public.users (
	id,
	username,
	"password",
	status)
VALUES (
	'b4588814-dee9-4074-b094-ae14d510bc02'::uuid,
	'sir-geronimo',
	'$2a$10$XZfE3Wxjl1.PGKT3pRZGneC/oXcm1KcGBe2OjPGntJq7fXZUkwGoK',
	1
);


-- public.records definition

CREATE TABLE public.records (
	id uuid NOT NULL,
	operation_id uuid NOT NULL,
	user_id uuid NOT NULL,
	amount int8 NOT NULL,
	user_balance int8 NOT NULL,
	operation_response text NOT NULL,
	"date" timestamptz DEFAULT now() NOT NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT records_pkey PRIMARY KEY (id),
	CONSTRAINT fk_operations_records FOREIGN KEY (operation_id) REFERENCES public.operations(id),
	CONSTRAINT fk_records_user FOREIGN KEY (user_id) REFERENCES public.users(id)
);

COMMIT;
