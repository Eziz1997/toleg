CREATE DATABASE toleg;
\c toleg;
set client_encoding to 'UTF-8';
create extension if not exists "uuid-ossp";

CREATE TABLE merchant (
                          "uuid" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
                          "order_number" numeric(12,0) NOT NULL ,
                          "currency" character varying(250) NOT NULL DEFAULT '',
                          "language" character varying(250) NOT NULL DEFAULT '',
                          "password" character varying(250) NOT NULL DEFAULT '',
                          "return_url" character varying(250) NOT NULL DEFAULT '',
                          "username" character varying(250) NOT NULL DEFAULT '',
                          "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL,
                          "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')  NOT NULL
);

insert into merchant (order_number,currency,language,password,return_url,username)VALUES(0112201700003,'934','ru','Jt3dsRvgTPdfPfA',
'https://mpi.gov.tm/payment/finish.html%3Flogin%3D102211004516%26password%3DJt3dsRvgTPdfPfA','102211004516');