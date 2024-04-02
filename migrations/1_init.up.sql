create table services
(
    uid           uuid primary key default gen_random_uuid(),
    name          varchar unique not null,
    description   text,
    current_price integer        not null
);


create table users
(
    uid      uuid primary key default gen_random_uuid(),
    username varchar unique not null
);

create table stats
(
    uid         uuid primary key default gen_random_uuid(),
    user_uid    uuid references users (uid)    not null,
    service_uid uuid references services (uid) not null,
    amount      integer                        not null,
    created_at  timestamp        default now()
);