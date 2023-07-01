package gatekeepercore

import (
	"context"
	"encoding/json"
	"feather-ai/service-core/lib/config"
	"feather-ai/service-core/lib/gatekeeper"
	"feather-ai/service-core/lib/storage"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type anonGateKeeperCore struct {
	TimePeriodInHours time.Duration
	ReqLimitPerPeriod int
}

func New(manager storage.StorageManager, opts config.Options) gatekeeper.AnonGateKeeper {
	prepareDBQueries(manager.GetDB())
	return &anonGateKeeperCore{
		TimePeriodInHours: time.Duration(1) * time.Hour,
		ReqLimitPerPeriod: 10000,
	}
}

func (m *anonGateKeeperCore) CanPerformRequest(ctx context.Context, ip string, usage string) error {
	// Check if an IP address is within limit. If not, return an error.
	// Also resets the counters for the IP if a new time block has begin
	result := []gatekeeper.AnonRequestDTO{}

	args := map[string]interface{}{
		"ip": ip,
	}
	err := gQueryCounterByIPStmt.QueryMany(ctx, args, &result)
	if err != nil {
		logrus.WithContext(ctx).Errorf("GateKeeper.Anon.CanPerformRequest: %s - %v - Allowing\n", ip, err)
		return nil
	}

	// If no records, allow it
	if len(result) == 0 {
		return nil
	}

	// If allocation expired, allocate a new block and allow
	if result[0].ReqAllocationTime.Before(time.Now().UTC().Add(-m.TimePeriodInHours)) {

		dto := gatekeeper.AnonRequestDTO{
			IP:                ip,
			ReqCount:          0,
			ReqLimit:          m.ReqLimitPerPeriod,
			ReqAllocationTime: time.Now().UTC(),
		}
		err := gResetCountersByIPStmt.Write(ctx, &dto)

		// Error resetting! for now allow
		if err != nil {
			logrus.WithContext(ctx).Errorf("GateKeeper.Anon.CanPerformRequest: Error resetting counters for IP=%s - err=%v - Allowing\n", ip, err)
			return nil
		}

		// And allow
		return nil
	} else if result[0].ReqCount < result[0].ReqLimit {
		// Below limit, allow
		return nil
	}

	// Exceeded limit - block it
	// :TODO: Maybe clear out IP here
	data, _ := json.Marshal(&result[0])
	return errors.Errorf(string(data))
}

func (m *anonGateKeeperCore) RecordRequest(ctx context.Context, ip string, usage string) error {

	// Set the req count to 1 (this request). If this IP already exists, it will
	// not insert this record, instead it will increment the count by 1 of the existing request.
	// If there is no record for this IP, then a new row is added, with a count set to 1.
	data := gatekeeper.AnonRequestDTO{
		IP:                ip,
		ReqLimit:          m.ReqLimitPerPeriod,
		ReqCount:          1,
		ReqAllocationTime: time.Now().UTC(),
	}

	err := gInsertOrUpdateAnonIPStmt.Write(ctx, &data)
	if err != nil {
		logrus.WithContext(ctx).Errorf("GateKeeper.RecordRequest: %v", err)
		return err
	}
	return nil
}
