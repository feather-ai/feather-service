BEGIN;

create table if not exists gk_anon (
	ip varchar(24) primary key,
	req_count int,
	req_limit int,
	req_allocation_time timestamp not null
);

COMMIT;

---- create above / drop below ----

DROP TABLE IF EXISTS "gk_anon";