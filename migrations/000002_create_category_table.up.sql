CREATE TABLE IF NOT EXISTS test.category(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) WITH TIME ZONE DEFAULT null,
    name character varying NOT NULL,
    parent_category uuid,
    CONSTRAINT "PK_category_id" PRIMARY KEY (id),
    CONSTRAINT "FK_parent_category" FOREIGN key(parent_category) REFERENCES test.category(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);