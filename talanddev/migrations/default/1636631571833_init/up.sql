CREATE TABLE history( id serial NOT NULL, amount integer NOT NULL, currency_from text NOT NULL, currency_to text NOT NULL, result integer NOT NULL, PRIMARY KEY(id));