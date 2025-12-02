-- 20251202161853_init_table_todos.sql

-- +goose Up
create extension if not exists "pgcrypto";

create table if not exists todos (
	id uuid primary key default gen_random_uuid(),
	text text not null,
	is_done boolean not null default false
);

-- +goose Down
drop table if exists todos;