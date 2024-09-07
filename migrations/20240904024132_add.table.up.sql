CREATE SCHEMA IF NOT EXISTS numeris;

CREATE TABLE numeris."user" (
    "id" TEXT NOT NULL unique,
    "name" TEXT NOT NULL,
    "username" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "ref" TEXT NOT NULL,
    "outlet_id" TEXT NOT NULL,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

CREATE FUNCTION numeris.UPDATED_AT_COLUMN() RETURNS TRIGGER 
AS $$ BEGIN NEW.updated_at = now();
	RETURN NEW;
	END;
	$$ language 'plpgsql';

ALTER TABLE numeris."user" ADD PRIMARY KEY ("id");

CREATE INDEX "user_name_index" ON numeris."user" ("name");

ALTER TABLE numeris."user" ADD CONSTRAINT "user_ref_unique" UNIQUE ("ref");

ALTER TABLE numeris."user"
ADD CONSTRAINT "user_username_unique" UNIQUE ("username");

CREATE INDEX "user_outlet_id_index" ON numeris."user" ("outlet_id");

CREATE TABLE numeris."customer" (
    "id" TEXT NOT NULL unique,
    "name" TEXT NOT NULL,
    "mobile_no" TEXT NOT NULL,
    "address" TEXT NOT NULL,
    "outlet_id" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "created_at" timestamptz not null default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."customer" ADD PRIMARY KEY ("id");

CREATE INDEX "customer_name_index" ON numeris."customer" ("name");

CREATE INDEX "customer_mobile_no_index" ON numeris."customer" ("mobile_no");

CREATE INDEX "customer_address_index" ON numeris."customer" ("address");

CREATE INDEX "customer_outlet_id_index" ON numeris."customer" ("outlet_id");

CREATE INDEX "customer_email_index" ON numeris."customer" ("email");

CREATE TABLE numeris."variant" (
    "id" TEXT NOT NULL unique,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "item_id" TEXT NOT NULL,
    "outlet_id" TEXT NOT NULL,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."variant" ADD PRIMARY KEY ("id");

CREATE INDEX "variant_name_index" ON numeris."variant" ("name");

CREATE INDEX "variant_item_id_index" ON numeris."variant" ("item_id");

CREATE INDEX "variant_outlet_id_index" ON numeris."variant" ("outlet_id");

CREATE TABLE numeris."currency" (
    "id" TEXT NOT NULL unique,
    "name" TEXT NOT NULL,
    "code" TEXT NOT NULL unique,
    "symbol" TEXT NOT NULL
);

ALTER TABLE numeris."currency" ADD PRIMARY KEY ("id");

CREATE TABLE numeris."outlet" (
    "id" TEXT NOT NULL unique,
    "name" TEXT NOT NULL,
    "is_default" BOOLEAN NOT NULL DEFAULT FALSE,
    "address" TEXT NOT NULL,
    "org_id" TEXT NOT NULL,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."outlet" ADD PRIMARY KEY ("id");

CREATE UNIQUE INDEX "outlet_org_id_is_default_index" ON numeris."outlet" ("org_id", "is_default");

CREATE INDEX "outlet_name_index" ON numeris."outlet" ("name");

CREATE INDEX "outlet_org_id_index" ON numeris."outlet" ("org_id");

CREATE TABLE numeris."currency_measure" (
    "id" TEXT NOT NULL,
    "currency_id" TEXT NOT NULL,
    "measure_id" TEXT NOT NULL,
    "price" DECIMAL(8, 2) NOT NULL
);

ALTER TABLE numeris."currency_measure" ADD PRIMARY KEY ("id");

create type numeris.measure_entity as enum('variant', 'invoice');

CREATE TABLE numeris."measure" (
    "id" TEXT NOT NULL unique,
    "entity" numeris.measure_entity NOT NULL,
    "entity_id" TEXT NOT NULL,
    "unit" TEXT NOT NULL,
    "quantity" DECIMAL(8, 2) NOT NULL,
    "is_active" BOOLEAN DEFAULT TRUE,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."measure" ADD PRIMARY KEY ("id");

CREATE INDEX "measure_entity_index" ON numeris."measure" ("entity");

CREATE INDEX "measure_entity_id_index" ON numeris."measure" ("entity_id");

CREATE TABLE numeris."item" (
    "id" TEXT NOT NULL unique,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "category_id" TEXT NOT NULL,
    "sku" TEXT NOT NULL,
    "outlet_id" TEXT NOT NULL,
    "is_deleted" BOOLEAN NOT NULL,
    "created_by" TEXT NOT NULL,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."item" ADD PRIMARY KEY ("id");

CREATE INDEX "item_name_index" ON numeris."item" ("name");

CREATE INDEX "item_category_id_index" ON numeris."item" ("category_id");

ALTER TABLE numeris."item" ADD CONSTRAINT "item_sku_unique" UNIQUE ("sku");

create type numeris.note_entity as enum(
    'organization',
    'outlet',
    'user',
    'customer',
    'item',
    'measure',
    'variant',
    'invoice',
    'invoice_boq',
    'inventory',
    'brand',
    'category'
);

CREATE TABLE numeris."note" (
    "id" TEXT NOT NULL unique,
    "entity" numeris.note_entity NOT NULL,
    "entity_id" TEXT NOT NULL,
    "note" TEXT NOT NULL,
    "command" jsonb NOT NULL,
    "created_by" TEXT NOT NULL,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."note" ADD PRIMARY KEY ("id");

CREATE INDEX "note_entity_index" ON numeris."note" ("entity");

CREATE INDEX "note_entity_id_index" ON numeris."note" ("entity_id");

CREATE TABLE numeris."category" (
    "id" TEXT NOT NULL unique,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "outlet_id" TEXT NOT NULL,
    "is_deleted" BOOLEAN NOT NULL,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."category" ADD PRIMARY KEY ("id");

CREATE INDEX "category_outlet_id_index" ON numeris."category" ("outlet_id");

CREATE TABLE numeris."organisation" (
    "id" TEXT NOT NULL unique,
    "name" TEXT NOT NULL,
    "reference" TEXT NOT NULL,
    "address" TEXT NOT NULL,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."organisation" ADD PRIMARY KEY ("id");

ALTER TABLE numeris."organisation"
ADD CONSTRAINT "organisation_name_unique" UNIQUE ("name");

ALTER TABLE numeris."organisation"
ADD CONSTRAINT "organisation_reference_unique" UNIQUE ("reference");

CREATE TABLE numeris."invoice_boq" (
    "id" TEXT NOT NULL unique, 
    "variant_id" TEXT NOT NULL,
    "invoice_id" TEXT NOT NULL,
    "measure_id" TEXT NOT NULL references numeris."measure" ("id"),
    "quantity" DECIMAL(8, 2) NOT NULL,
    "unit_price" DECIMAL(8, 2) NOT NULL,
    "total" DECIMAL(8, 2) NOT NULL,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."invoice_boq" ADD PRIMARY KEY ("id");

CREATE TABLE numeris."bank_account" (
    "id" TEXT NOT NULL unique, 
    "name" TEXT NOT NULL,
    "outlet_id" TEXT NOT NULL,
    "account_no" TEXT NOT NULL,
    "routing_no" TEXT NOT NULL,
    "account_type" TEXT NOT NULL,
    "bank_name" TEXT NOT NULL,
    "currency_id" text not null references numeris.currency (id),
    "is_active" BOOLEAN DEFAULT TRUE,
    "is_deleted" BOOLEAN NOT NULL,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."bank_account" ADD PRIMARY KEY ("id");

create type numeris.invoice_status as enum(
    'draft',
    'sent',
    'paid',
    'rejected',
    'partial',
    'unpaid',
    'overdue',
    'pending',
    'active'
);

CREATE TABLE numeris."invoice" (
    "id" TEXT NOT NULL unique,
    "name" TEXT NOT NULL,
    "ref" TEXT NOT NULL,
    "currency_id" TEXT NOT NULL,
    "outlet_id" TEXT NOT NULL references numeris.outlet (id),
    "customer_id" TEXT NOT NULL references numeris.customer (id),
    "total" DECIMAL(8, 2) NOT NULL,
    "due_date" DATE,
    "sub_total" DECIMAL(8, 2) NOT NULL,
    "discount" DECIMAL(8, 2),
    "created_at" timestamptz not null default now(),
    "updated_at" timestamptz not null default current_timestamp,
    "reminder" jsonb NOT NULL,
    "bank_account_id" TEXT NOT NULL references numeris.bank_account (id),
    "status" numeris.invoice_status NOT NULL,
    "created_by" TEXT
);

ALTER TABLE numeris."invoice" ADD PRIMARY KEY ("id");

CREATE TABLE numeris."inventory" (
    "id" TEXT NOT NULL unique,
    "variant_id" TEXT NOT NULL,
    "measure_id" TEXT NOT NULL references numeris."measure" ("id"),
    "opening_stock" DECIMAL(8, 2) NOT NULL DEFAULT 0,
    "added_stock" DECIMAL(8, 2) NOT NULL DEFAULT 0,
    "issued_stock" DECIMAL(8, 2) NOT NULL DEFAULT 0,
    "stock_balance" DECIMAL(8, 2) NOT NULL DEFAULT 0,
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default current_timestamp
);

ALTER TABLE numeris."inventory" ADD PRIMARY KEY ("id");

CREATE INDEX "inventory_variant_id_index" ON numeris."inventory" ("variant_id");

create type numeris.batch_state as enum(
    'none',
    'added',
    'sold',
    'waste',
    'returned'
);

create type numeris.batch_entity as enum('invoice', 'inventory');

create table numeris.batch (
    id text not null unique,
    variant_id text not null references "numeris".variant (id) on delete cascade,
    "measure_id" TEXT NOT NULL references numeris."measure" ("id"), 
    entity_id text,
    entity numeris.batch_entity not null,
    state numeris.batch_state not null,
    stock numeric not null check (stock >= 0),
    created_at timestamptz default now(),
    updated_at timestamptz default current_timestamp,
    primary key (id, variant_id)
);

create index batch_variant_id on numeris.batch using btree (variant_id);

create trigger update_batch_updated_at before
update on numeris.batch for each row
execute procedure numeris.updated_at_column ();

ALTER TABLE numeris."bank_account"
ADD CONSTRAINT "bank_account_outlet_id_foreign" FOREIGN KEY ("outlet_id") REFERENCES numeris."outlet" ("id");

ALTER TABLE numeris."variant"
ADD CONSTRAINT "variant_outlet_id_foreign" FOREIGN KEY ("outlet_id") REFERENCES numeris."outlet" ("id");

ALTER TABLE numeris."customer"
ADD CONSTRAINT "customer_outlet_id_foreign" FOREIGN KEY ("outlet_id") REFERENCES numeris."outlet" ("id");

ALTER TABLE numeris."currency_measure"
ADD CONSTRAINT "currency_measure_measure_id_foreign" FOREIGN KEY ("measure_id") REFERENCES numeris."measure" ("id");

ALTER TABLE numeris."invoice_boq"
ADD CONSTRAINT "invoice_boq_variant_id_foreign" FOREIGN KEY ("variant_id") REFERENCES numeris."variant" ("id");

ALTER TABLE numeris."invoice_boq"
ADD CONSTRAINT "invoice_boq_invoice_id_foreign" FOREIGN KEY ("invoice_id") REFERENCES numeris."invoice" ("id");

ALTER TABLE numeris."user"
ADD CONSTRAINT "user_outlet_id_foreign" FOREIGN KEY ("outlet_id") REFERENCES numeris."outlet" ("id");

ALTER TABLE numeris."item"
ADD CONSTRAINT "item_category_id_foreign" FOREIGN KEY ("category_id") REFERENCES numeris."category" ("id");

ALTER TABLE numeris."outlet"
ADD CONSTRAINT "outlet_org_id_foreign" FOREIGN KEY ("org_id") REFERENCES numeris."organisation" ("id");

ALTER TABLE numeris."inventory"
ADD CONSTRAINT "inventory_variant_id_foreign" FOREIGN KEY ("variant_id") REFERENCES numeris."variant" ("id");

ALTER TABLE numeris."currency_measure"
ADD CONSTRAINT "currency_measure_currency_id_foreign" FOREIGN KEY ("currency_id") REFERENCES numeris."currency" ("id");

ALTER TABLE numeris."variant"
ADD CONSTRAINT "variant_item_id_foreign" FOREIGN KEY ("item_id") REFERENCES numeris."item" ("id");

ALTER TABLE numeris."invoice"
ADD CONSTRAINT "invoice_currency_id_foreign" FOREIGN KEY ("currency_id") REFERENCES numeris."currency" ("id");

ALTER TABLE numeris."measure"
ADD CONSTRAINT "measure_entity_id_variant_foreign" FOREIGN KEY ("entity_id") REFERENCES numeris."variant" ("id");

ALTER TABLE numeris."note"
ADD CONSTRAINT "note_created_by_foreign" FOREIGN KEY ("created_by") REFERENCES numeris."user" ("id");