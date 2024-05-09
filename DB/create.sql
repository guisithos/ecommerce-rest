-- CREATE DATABASE IF NOT EXISTS ecommerce-rest;
\c ecommerce-rest;

DROP TABLE IF EXISTS public.products;

CREATE TABLE public.products (
    id serial NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    price numeric(10,2) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

INSERT INTO public.products (name, description, price, created_at, updated_at)
VALUES
    ('Cabernet Sauvignon 2018', 'Rich red wine with blackcurrant notes', 49.99, now(), now()),
    ('Chardonnay Reserve 2020', 'Oak-aged white wine with buttery texture', 39.50, now(), now()),
    ('Merlot 2019', 'Smooth red wine with plum and cherry flavors', 29.95, now(), now()),
    ('Sauvignon Blanc 2021', 'Crisp white wine with citrusy aromas', 34.00, now(), now()),
    ('Pinot Noir 2017', 'Elegant red wine with hints of raspberry and spice', 59.99, now(), now()),
    ('Rosé 2020', 'Refreshing pink wine with strawberry undertones', 27.50, now(), now()),
    ('Prosecco Brut NV', 'Sparkling wine with floral and fruity notes', 39.99, now(), now()),
    ('Malbec Reserve 2018', 'Full-bodied red wine with dark fruit flavors', 45.95, now(), now()),
    ('Riesling 2020', 'Off-dry white wine with peach and apricot aromas', 32.00, now(), now()),
    ('Syrah/Shiraz 2019', 'Bold red wine with blackberry and pepper nuances', 54.75, now(), now());
