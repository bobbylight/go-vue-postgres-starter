drop schema if exists demo_schema cascade;
create schema demo_schema;
set search_path to demo_schema;

drop table if exists widgets cascade;
create table widgets (
    id varchar(36) primary key,
    name varchar(80) not null,
    price double precision not null
);
create index widgets_name_ix01 on widgets (lower(name));

-- Test data
insert into widgets values
    ('111', 'widget1', '11.11'),
    ('222', 'widget2', '22.22'),
    ('333', 'widget3', '33.33')
;
