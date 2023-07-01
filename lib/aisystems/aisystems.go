package aisystems

import (
	"context"
	"encoding/json"
	"feather-ai/service-core/lib/feather"
	"time"

	uuid "github.com/satori/go.uuid"
)

type SystemDTO struct {
	SystemID    uuid.UUID      `db:"system_id"`
	UserID      uuid.UUID      `db:"user_id"`
	Name        string         `db:"name"`
	Slug        string         `db:"slug"`
	Description feather.String `db:"description"`
	Keywords    feather.String `db:"keywords"`
	Created     time.Time      `db:"created"`
}

type SystemVersionDTO struct {
	VersionID int64          `db:"version_id"`
	SystemID  uuid.UUID      `db:"system_id"`
	Tag       feather.String `db:"tag"`
	Schema    string         `db:"schema"`
	Created   time.Time      `db:"created"`
}

type FileDTO struct {
	FileID          int       `db:"file_id"`
	SystemVersionID int64     `db:"version_id"`
	FileName        string    `db:"file_name"`
	FileType        string    `db:"file_type"`
	FileSize        int       `db:"file_size"`
	URL             string    `db:"url"`
	Created         time.Time `db:"created"`
}

type SystemConfigDTO struct {
	SystemID       uuid.UUID      `db:"system_id"`
	LambdaDispatch feather.String `db:"lambda_dispatch"`
}

type AISystemsManager interface {
	// Get all systems - WARNING - no pagination - only to be used for MVP until marketplace is in place
	GetSystems(ctx context.Context) ([]SystemDTO, error)

	// Get all the systems belonging to a user
	GetSystemsByUser(ctx context.Context, userId uuid.UUID) ([]SystemDTO, error)

	// Get a system by ID
	GetSystemByID(ctx context.Context, systemId uuid.UUID) (*SystemDTO, error)

	// Get a system by user and name. Will return 1 or 0 systems
	GetSystemByUserAndSlug(ctx context.Context, userId uuid.UUID, slug string) *SystemDTO

	// Get full details of a specific system
	GetSystemDetails(ctx context.Context, systemId uuid.UUID) (*SystemDTO, *SystemVersionDTO, []FileDTO, error)

	// Get or create a system.
	GetOrCreateNewSystem(ctx context.Context, userId uuid.UUID, name string, slug string, description string) (uuid.UUID, error)

	// Create a new version of a system
	CreateNewVersion(ctx context.Context, systemId uuid.UUID, tag string, schema string) (int64, error)

	// Runs a system
	RunSystem(ctx context.Context, systemId uuid.UUID, stepIndex int, inputDataJson string) ([]json.RawMessage, string, string, error)

	// Get the execute request for the specified system
	DebugGetRunSystemSchema(ctx context.Context, systemId uuid.UUID, stepIndex int) (interface{}, error)
}
