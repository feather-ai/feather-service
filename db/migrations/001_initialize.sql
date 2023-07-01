BEGIN;

create table if not exists auth0(
	auth0_id varchar(32) primary key,
	feather_id uuid not null
);

create table if not exists users(
	user_id uuid primary key,
	name varchar(64) unique not null
);
create index users_by_name on users(name);

create table if not exists api_keys (
	api_key uuid primary key,
	user_id uuid not null,
	key_name varchar(32),
	created timestamp not null,
	revoked bool,
	
	foreign key(user_id) references users(user_id)
);
create index api_keys_user_id_index on api_keys(user_id);

-- A system is owned by a user
create table if not exists systems (
	system_id 		uuid primary key,
	user_id 		uuid not null,
	name 			varchar(64) not null,
	slug			varchar(64) not null,
	created 		timestamp not null,
	description 	text,
	keywords 		text,
	
	foreign key(user_id) references users(user_id)
);
create index system_user_id_index on systems(user_id);
create index system_user_and_slug_index on systems(user_id, slug);

-- Table holding, for each system, a list of published versions for that system.
-- Each version is made up of files amongst other things
create table if not exists system_versions (
	version_id        serial primary key,
	system_id         uuid not null,
	tag               varchar(64) not null,
	created           timestamp not null,
  	schema            text not null,
	
	foreign key(system_id) references systems(system_id)
);
create index system_version_id_index on system_versions(system_id);

-- Master table holding all the raw files we have in S3.
-- Files are grouped into system versions. All files for a version are used for execution
create table if not exists files (
	file_id 			serial primary key,
	version_id 			int not null,
	file_name 			varchar(128) not null,
	file_type 			varchar(16) not null,
	file_size 			int,
	url 				varchar(256),
  	created 			timestamp not null,
	
	foreign key(version_id) references system_versions(version_id)
);
create index files_version_id_index on files (version_id);

-- Table holding all the in progress uploads for a user and system
CREATE TABLE IF NOT EXISTS upload_requests (
  id                         UUID PRIMARY KEY,
  user_id                    UUID NOT NULL,
  system_id                  UUID NOT NULL,
  version_tag                VARCHAR(64) NOT NULL,
  create_time                TIMESTAMP NOT NULL,
  expire_time                TIMESTAMP NOT NULL,
  code_files                 VARCHAR NOT NULL,
  model_files                VARCHAR NOT NULL,
  code_files_signed_url      VARCHAR NOT NULL,
  model_files_signed_url     VARCHAR NOT NULL,
  schema					 text not  null,

  foreign key(user_id) references users(user_id),
  foreign key(system_id) references systems(system_id)
);
CREATE INDEX user_upload_requests_idx ON "upload_requests" (user_id);


COMMIT;

---- create above / drop below ----

DROP INDEX user_upload_requests_idx;
DROP INDEX files_version_id_index;
DROP INDEX system_user_id_index;
DROP INDEX system_version_id_index;

DROP TABLE IF EXISTS "upload_requests";
DROP TABLE IF EXISTS "files";
DROP TABLE IF EXISTS "system_versions";
DROP TABLE IF EXISTS "systems";
DROP TABLE IF EXISTS "api_keys";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "auth0";

