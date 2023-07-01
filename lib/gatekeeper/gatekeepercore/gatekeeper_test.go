package gatekeepercore

import (
	"context"
	"feather-ai/service-core/lib/gatekeeper"
	"feather-ai/service-core/lib/storage/storagemock"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCanPerformRequest_NewIP(t *testing.T) {
	core := anonGateKeeperCore{
		TimePeriodInHours: time.Duration(24) * time.Hour,
		ReqLimitPerPeriod: 10,
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	queryStatement := storagemock.NewMockPreparedStatement(controller)
	gQueryCounterByIPStmt = queryStatement

	queryStatement.EXPECT().QueryMany(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := core.CanPerformRequest(context.Background(), "127.0.0.1", "some_usage")
	assert.NoError(t, err)
}

func TestCanPerformRequest_TooManyRequests(t *testing.T) {
	// Test when a user has performed too many requests - it should return an error
	core := anonGateKeeperCore{
		TimePeriodInHours: time.Duration(24) * time.Hour,
		ReqLimitPerPeriod: 10,
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	gInsertOrUpdateAnonIPStmt = storagemock.NewMockPreparedStatement(controller)

	queryStatement := storagemock.NewMockPreparedStatement(controller)
	gQueryCounterByIPStmt = queryStatement

	queryStatement.EXPECT().QueryMany(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(ctx context.Context, args interface{}, results interface{}) {
		data := []gatekeeper.AnonRequestDTO{}
		data = append(data, gatekeeper.AnonRequestDTO{IP: "127.0.0.2", ReqCount: 11, ReqLimit: 10, ReqAllocationTime: time.Now().UTC()})

		value := reflect.ValueOf(results)
		direct := reflect.Indirect(value)
		v := reflect.ValueOf(data[0])
		direct.Set(reflect.Append(direct, v))
	}).Return(nil)

	err := core.CanPerformRequest(context.Background(), "127.0.0.2", "some_usage")
	assert.Error(t, err) // Should fail
}

func TestCanPerformRequest_AllocationExpired(t *testing.T) {
	// Test when a user has expired their allocation. We should get a new allocation
	core := anonGateKeeperCore{
		TimePeriodInHours: time.Duration(24) * time.Hour,
		ReqLimitPerPeriod: 10,
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	gInsertOrUpdateAnonIPStmt = storagemock.NewMockPreparedStatement(controller)

	queryStatement := storagemock.NewMockPreparedStatement(controller)
	gQueryCounterByIPStmt = queryStatement

	queryStatement.EXPECT().QueryMany(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(ctx context.Context, args interface{}, results interface{}) {
		data := []gatekeeper.AnonRequestDTO{}
		data = append(data, gatekeeper.AnonRequestDTO{IP: "127.0.0.2", ReqCount: 999, ReqLimit: 10, ReqAllocationTime: time.Now().UTC().Add(-time.Hour * 48)})

		value := reflect.ValueOf(results)
		direct := reflect.Indirect(value)
		v := reflect.ValueOf(data[0])
		direct.Set(reflect.Append(direct, v))
	}).Return(nil)

	writeStatement := storagemock.NewMockPreparedStatement(controller)
	gResetCountersByIPStmt = writeStatement
	writeStatement.EXPECT().Write(gomock.Any(), gomock.Any()).Times(1).Return(nil)

	err := core.CanPerformRequest(context.Background(), "127.0.0.2", "some_usage")
	assert.NoError(t, err) // Should succeed!
}
