create table aiservices
(
    uid         uuid primary key default gen_random_uuid(),
    title       varchar unique   not null,
    description text,
    price       double precision not null
);


create table users
(
    uid      uuid primary key default gen_random_uuid(),
    username varchar unique not null
);

create table statistics
(
    uid           uuid primary key default gen_random_uuid(),
    user_uid      uuid references users (uid)      not null,
    aiservice_uid uuid references aiservices (uid) not null,
    amount        double precision                 not null,
    created_at    timestamp        default now()   not null
);