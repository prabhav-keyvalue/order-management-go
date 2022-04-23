
CREATE TABLE IF NOT EXISTS test.order(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) WITH TIME ZONE DEFAULT NULL,
    customer_id uuid NOT NULL,
    total_quantity integer NOT NULL,
    total_price numeric(12,3) NOT NULL,
    CONSTRAINT "PK_order_id" PRIMARY KEY (id),
    CONSTRAINT "FK_order_customer_id" FOREIGN KEY (customer_id) REFERENCES test.customer(id)  ON DELETE NO ACTION ON UPDATE NO ACTION
);
