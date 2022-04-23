CREATE TABLE IF NOT EXISTS test.order_item(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) WITH TIME ZONE DEFAULT NULL,
    order_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity integer NOT NULL,
    price numeric(12,3) NOT NULL,
    row_total numeric(12,3) NOT NULL,
    CONSTRAINT "PK_order_item_id" PRIMARY KEY(id),
    CONSTRAINT "FK_order_item_order_id" FOREIGN KEY (order_id) REFERENCES test.order(id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    CONSTRAINT "FK_order_item_product_id" FOREIGN KEY (product_id) REFERENCES test.product(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);