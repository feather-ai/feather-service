// +build integration

package uploadcore_test

import (
	"context"
	"feather-ai/service-core/lib/config"
	"feather-ai/service-core/lib/storage/storagemock"
	"feather-ai/service-core/lib/upload/uploadcore"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := config.CreateRequestContext(context.Background(), &config.RequestContext{UserID: uuid.NewV4()})

	options := config.Options{}
	storagemanager := storagemock.NewMockStorageManager(ctrl)
	mockdb := storagemock.NewMockDB(ctrl)
	mocks3 := storagemock.NewMockBlobStorage(ctrl)

	// Setup the expected function calls
	storagemanager.EXPECT().GetDB().Return(mockdb)
	storagemanager.EXPECT().GetBlobStorage(gomock.Any()).Return(mocks3)
	mocks3.EXPECT().GenerateUploadURL("somename", time.Minute*15, ctx).Return("some_url", nil)
	mockdb.EXPECT().PrepareQuery(gomock.Any()).Return(storagemock.NewMockPreparedStatement(ctrl))

	uploader := uploadcore.New(storagemanager, options)

	req, err := uploader.CreateUploadRequest("somename", ctx)
	assert.NoError(t, err, "Upload request failed")
	assert.Equal(t, "somename", req.ModelName, "Resultant model name does not match")
}
