drop schema if exists dwitd cascade;
create schema dwitd;
set search_path to dwitd;

drop table if exists tasks cascade;
create table tasks (
    pk serial primary key,
    id varchar(36) unique not null,
    label varchar(256) not null,
    description varchar(256),
    status_cd varchar(16) not null,
    created_dttm timestamp not null,
    completed_dttm timestamp
);
create index tasks_id_index_ix01 on tasks (id);
create index tasks_label_ix02 on tasks (lower(label));
create index tasks_desc_ix03 on tasks (lower(description));

-- Test data
insert into tasks
    (id, label, description, status_cd, created_dttm, completed_dttm)
values
    ('111', 'Do laundry 1', NULL, 'notStarted', '2020-03-01T00:00:00Z', NULL),
    ('222', 'Do laundry 2', NULL, 'notStarted', '2020-03-02T00:00:00Z', NULL),
    ('333', 'Do laundry 3', NULL, 'notStarted', '2020-03-03T00:00:00Z', NULL),
    ('444', 'Do laundry 4', NULL, 'notStarted', '2020-03-04T00:00:00Z', NULL),
    ('555', 'Do laundry 5', NULL, 'notStarted', '2020-03-05T00:00:00Z', NULL)
;
