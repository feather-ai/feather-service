package aisystemscore

import (
	"feather-ai/service-core/lib/storage"
)

// Insert a new System
var gInsertSystemStmt storage.PreparedStatement

const InsertSystemRequest = `INSERT INTO systems
(system_id, user_id, name, slug, description, keywords, created)
VALUES(:system_id, :user_id, :name, :slug, :description, :keywords, :created)`

// Query Systems by  UserID
var gQuerySystemsByUserStmt storage.PreparedStatement

const GetSystemsByUserQuery = `SELECT system_id, user_id, name, slug, description, keywords, created FROM systems
WHERE user_id=:user_id`

// Query Systems by  UserID
var gQueryAllSystemsStmt storage.PreparedStatement

const GetAllSystemsQuery = `SELECT system_id, user_id, name, slug,  description, keywords, created FROM systems`

// Query Systems by  UserID
var gQuerySystemsByIDStmt storage.PreparedStatement

const GetSystemsByIDQuery = `SELECT system_id, user_id, name, slug, description, keywords, created FROM systems
WHERE system_id=:system_id`

// Query Systems by  UserID and Name
var gQuerySystemsByUserAndSlugStmt storage.PreparedStatement

const GetSystemsByUserAndSlugQuery = `SELECT system_id, user_id, name, slug, description, keywords, created FROM systems
WHERE user_id=:user_id and slug=:slug`

// Get the latest version of a system
var gQueryLatestSystemVersionBySystemIdStmt storage.PreparedStatement

const GetLatestSystemVersionBySystemIDQuery = `SELECT version_id, system_id, tag, created, schema FROM system_versions where 
system_id = :system_id ORDER BY version_id DESC LIMIT 1`

// Get the files for a specific system version
var gQuerySystemVersionFilesStmt storage.PreparedStatement

const GetSystemVersionFilesQuery = `SELECT version_id, file_id, file_name, file_type, file_size, url, created FROM files where 
version_id = :version_id`

// Insert a new System Version
var gInsertSystemVersionStmt storage.PreparedStatement

const InsertSystemVersionRequest = `INSERT INTO system_versions
(system_id, tag, created, schema)
VALUES(:system_id, :tag, :created, :schema) RETURNING version_id`

// Query System Config by ID
var gQuerySystemConfigByIDStmt storage.PreparedStatement

const GetSystemConfigByIDQuery = `SELECT system_id, lambda_dispatch FROM model_config
WHERE system_id=:system_id`

/*
* Prepare all DB queries we will need to run here.
 */
func prepareDBQueries(storage storage.DB) {
	gInsertSystemStmt = storage.PrepareQuery(InsertSystemRequest)
	gQueryAllSystemsStmt = storage.PrepareQuery(GetAllSystemsQuery)
	gQuerySystemsByUserStmt = storage.PrepareQuery(GetSystemsByUserQuery)
	gQuerySystemsByIDStmt = storage.PrepareQuery(GetSystemsByIDQuery)
	gQuerySystemsByUserAndSlugStmt = storage.PrepareQuery(GetSystemsByUserAndSlugQuery)
	gQueryLatestSystemVersionBySystemIdStmt = storage.PrepareQuery(GetLatestSystemVersionBySystemIDQuery)
	gQuerySystemVersionFilesStmt = storage.PrepareQuery(GetSystemVersionFilesQuery)
	gInsertSystemVersionStmt = storage.PrepareQuery(InsertSystemVersionRequest)
	gQuerySystemConfigByIDStmt = storage.PrepareQuery(GetSystemConfigByIDQuery)
}
