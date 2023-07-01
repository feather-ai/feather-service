BEGIN;

-- Table that holds configuration information about a model
-- lambda_dispatch tells the system what Lambda to use to execute this model. If it's null or missing, then we use the generic runner
create table if not exists model_config (
	system_id         uuid primary key,
	lambda_dispatch	  varchar(64),

	foreign key(system_id) references systems(system_id)
);

COMMIT;

---- create above / drop below ----

DROP TABLE IF EXISTS "model_config";