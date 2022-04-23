CREATE TABLE IF NOT EXISTS test.customer(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) WITH TIME ZONE DEFAULT null,
    name character varying NOT NULL,
    phone varchar(30) NOT NULL,
    email varchar(255) NOT NULL,
    street character varying NOT NULL,
    city character varying NOT NULL,
    state character varying NOT NULL,
    profile_pic character varying NOT NULL,
    CONSTRAINT "PK_customer_id" PRIMARY KEY (id)
);