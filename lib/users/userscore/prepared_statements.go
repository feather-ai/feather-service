package userscore

import (
	"feather-ai/service-core/lib/storage"
)

var gInsertApiKeyStmt storage.PreparedStatement

const InsertApiKeyRequest = `INSERT INTO api_keys
(api_key, user_id, key_name, created, revoked)
VALUES(:api_key, :user_id, :key_name, :created, :revoked)`

// Insert a user
var gInsertUserStmt storage.PreparedStatement

const InsertUserRequest = `INSERT INTO users
(user_id, name)
VALUES(:user_id, :name)`

// Insert an auth0 user
var gInsertAuth0UserStmt storage.PreparedStatement

const InsertAuth0UserRequest = `INSERT INTO auth0
(auth0_id, feather_id)
VALUES(:auth0_id, :feather_id)`

// Fetch an Auth0 user
var gQueryAuth0UserIDStmt storage.PreparedStatement

const GetAuth0UserIdQuery = `SELECT feather_id FROM auth0
WHERE auth0_id=:auth0_id`

// Fetch an API Key by ID
var gQueryGetApiKeyStmt storage.PreparedStatement

const GetAPIKeyQuery = `SELECT api_key, user_id, key_name, created, revoked FROM api_keys
WHERE api_key=:api_key`

// Get all API keys for a user
var gQueryGetAllApiKeyStmt storage.PreparedStatement

const GetAllAPIKeyQuery = `SELECT api_key, user_id, key_name, created, revoked FROM api_keys
WHERE user_id=:user_id`

// Delete an API key
var gDeleteApiKeyStmt storage.PreparedStatement

const DeleteAPIKey = `DELETE FROM api_keys
WHERE api_key=:api_key;`

// Fetch a user by name
var gQueryUserByNameStmt storage.PreparedStatement

const GetUserByNameQuery = `SELECT user_id, name FROM users
WHERE name=:name`

// Fetch a user by id
var gQueryUserByIDStmt storage.PreparedStatement

const GetUserByIDQuery = `SELECT user_id, name FROM users
WHERE user_id=:user_id`

/*
* Prepare all DB queries we will need to run here.
 */
func prepareDBQueries(storage storage.DB) {
	gInsertUserStmt = storage.PrepareQuery(InsertUserRequest)
	gInsertAuth0UserStmt = storage.PrepareQuery(InsertAuth0UserRequest)
	gQueryAuth0UserIDStmt = storage.PrepareQuery(GetAuth0UserIdQuery)

	gQueryUserByNameStmt = storage.PrepareQuery(GetUserByNameQuery)
	gQueryUserByIDStmt = storage.PrepareQuery(GetUserByIDQuery)

	gInsertApiKeyStmt = storage.PrepareQuery(InsertApiKeyRequest)
	gQueryGetApiKeyStmt = storage.PrepareQuery(GetAPIKeyQuery)
	gQueryGetAllApiKeyStmt = storage.PrepareQuery(GetAllAPIKeyQuery)
	gDeleteApiKeyStmt = storage.PrepareQuery(DeleteAPIKey)
}
