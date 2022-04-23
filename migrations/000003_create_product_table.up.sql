CREATE TABLE IF NOT EXISTS test.product(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) WITH TIME ZONE DEFAULT NULL,
    name character varying NOT NULL,
    image character varying NOT NULL,
    description character varying NOT NULL,
    unit_price numeric(13,2) NOT NULL,
    category_id uuid NOT NULL,
    CONSTRAINT "PK_product_id" PRIMARY KEY (id),
    CONSTRAINT "FK_product_category_id" FOREIGN KEY (category_id) REFERENCES test.category(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);