package upload_test

import (
	"feather-ai/service-core/lib/upload/uploadmock"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestUpload(t *testing.T) {

	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	_ = uploadmock.NewMockUploadManager(ctrl)
}
