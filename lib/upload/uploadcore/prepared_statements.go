package uploadcore

import (
	"feather-ai/service-core/lib/storage"
)

// Insert a new Upload Request into the upload_requests table. On conflic it fails!
var gInsertUploadRequestStmt storage.PreparedStatement

const InsertUploadRequest = `INSERT INTO upload_requests
(id, user_id, create_time, expire_time, system_id, version_tag, code_files, model_files, code_files_signed_url, model_files_signed_url, schema)
VALUES(:id, :user_id, :create_time, :expire_time, :system_id, :version_tag, :code_files, :model_files, :code_files_signed_url, :model_files_signed_url, :schema)`

// Return all rows for a specific user
var gQueryUserUploadRequestsStmt storage.PreparedStatement

const UserUploadRequestsQuery = `SELECT id, user_id, create_time, expire_time, system_id, version_tag, code_files, model_files, code_files_signed_url, model_files_signed_url, schema FROM upload_requests
WHERE user_id=:userId`

// Load an upload request by id
var gQueryUploadRequestByIdStmt storage.PreparedStatement

const UploadRequestByIdQuery = `SELECT id, user_id, create_time, expire_time, system_id, version_tag, code_files, model_files, code_files_signed_url, model_files_signed_url, schema FROM upload_requests
WHERE id=:id`

//  Delete UploadRequest row
var gDeleteUploadRequestByIdStmt storage.PreparedStatement

const DeleteRequestByIdQuery = `DELETE FROM upload_requests  WHERE id=:id`

//  Insert a File record
var gInsertFileRequestStmt storage.PreparedStatement

const InsertFileRequest = `INSERT INTO files
(version_id, file_name, file_type, file_size, url, created)
VALUES(:version_id, :file_name, :file_type, :file_size, :url, :created)`

/*
* Prepare all DB queries we will need to run here.
 */
func prepareDBQueries(storage storage.DB) {
	gInsertUploadRequestStmt = storage.PrepareQuery(InsertUploadRequest)
	gQueryUserUploadRequestsStmt = storage.PrepareQuery(UserUploadRequestsQuery)
	gQueryUploadRequestByIdStmt = storage.PrepareQuery(UploadRequestByIdQuery)

	gInsertFileRequestStmt = storage.PrepareQuery(InsertFileRequest)
	gDeleteUploadRequestByIdStmt = storage.PrepareQuery(DeleteRequestByIdQuery)
}
