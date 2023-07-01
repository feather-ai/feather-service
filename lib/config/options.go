package config

type Options struct {
	DBurl     string `long:"database_url" description:"Connection string for DB" default:"host=localhost port=5432 sslmode=disable user=postgres dbname=feather-ai" env:"DATABASE_URL"`
	DebugUser bool   `long:"debug_user" description:"Enable debug user login" env:"DEBUG_USER"`

	// AWS credentials
	AwsRegion           string `long:"aws_region" description:"Region to use" default:"us-east-2" env:"AWS_REGION"`
	AwsKeyID            string `long:"aws_access_key_id" description:"The Key ID" default:"" env:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey  string `long:"aws_secret_access_key" description:"The Secret Key" default:"" env:"AWS_SECRET_ACCESS_KEY"`
	AwsUploadBucketName string `long:"s3_upload_bucket" description:"Name of bucket to use to upload models" default:"feather-ai-dev-front-storage" env:"S3_UPLOAD_BUCKET"`

	Auth0IdentityURL string `long:"auth0_identity_url" description:"URL for the auth0 /identity api" default:"https://feather-ai.us.auth0.com/userinfo" env:"AUTH0_IDENTITY_URL"`
}
