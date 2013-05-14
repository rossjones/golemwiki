begin;

create table golem_group (
    id SERIAL,
    name character varying(64)
    slug character varying(64)
);

create table golem_user (
    id SERIAL,
    username character varying(64)
);

create table golem_page (
    id SERIAL,
    name character varying(100)
);


commit;