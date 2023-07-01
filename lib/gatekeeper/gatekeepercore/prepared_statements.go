package gatekeepercore

import (
	"feather-ai/service-core/lib/storage"
)

// Insert or update the counter for an anonymous IP
var gInsertOrUpdateAnonIPStmt storage.PreparedStatement

const InsertOrUpdateAnonIP = `INSERT INTO gk_anon
(ip, req_count, req_limit, req_allocation_time)
VALUES(:ip, :req_count, :req_limit, :req_allocation_time) on conflict (ip) do update set req_count = gk_anon.req_count + 1`

// Query counts by IP
var gQueryCounterByIPStmt storage.PreparedStatement

const QueryCounterByIP = `SELECT ip, req_count, req_limit, req_allocation_time FROM gk_anon
WHERE ip=:ip`

// Reset counters
var gResetCountersByIPStmt storage.PreparedStatement

const ResetCountersByIPQuery = `UPDATE gk_anon set req_count = 0 where ip=:ip`

/*
* Prepare all DB queries we will need to run here.
 */
func prepareDBQueries(storage storage.DB) {
	gInsertOrUpdateAnonIPStmt = storage.PrepareQuery(InsertOrUpdateAnonIP)
	gQueryCounterByIPStmt = storage.PrepareQuery(QueryCounterByIP)
	gResetCountersByIPStmt = storage.PrepareQuery(ResetCountersByIPQuery)
}
