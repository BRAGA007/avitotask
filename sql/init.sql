create table if not exists reservations
(
    user_id    int          null,
    service_id int          null,
    order_id   int auto_increment
        primary key,
    cost       int          null,
    status     varchar(255) null
);
create table if not exists revenues
(
    user_id    int      null,
    service_id int      null,
    order_id   int auto_increment
        primary key,
    amount     int      null,
    created_at datetime null
);
create table if not exists transactions
(
    user_id     int          null,
    created_at  datetime     null,
    description varchar(255) null,
    amount      int          null
);
create table if not exists users
(
    id      int auto_increment
        primary key,
    balance int null
);
engine = InnoDB
default charset = utf8;

