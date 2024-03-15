CREATE TABLE
    users (
        id serial not null unique,
        name varchar(255) not null,
        username varchar(255) not null unique,
        password_hash varchar(255) not null
    );

CREATE TABLE
    accounts (
        id serial not null unique,
        owner_id int references users(id) on delete cascade not null,
        currency varchar(255) not null,
        balance int not null default 0
    );

CREATE TABLE 
    crypto_cur (
        id serial not null unique,
        account_id int references accounts(id) on delete cascade not null,
        coin varchar(255) not null,
        amount int not null default 0
    );

CREATE TABLE 
    exchange_operations (
        id serial not null unique,
        account_id int references accounts(id) on delete cascade not null,
        currency varchar(255) not null,
        coin varchar(255) not null,
        amount int not null,
        purch_price int not null,
        desire_time timestamp not null,
        status boolean not null default false
    );