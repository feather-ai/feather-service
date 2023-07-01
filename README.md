# Overview

Service-Core, for lack of a better name, represents the core service of feather AI, coordinating logic between the various sub-components.

## Requirements

- go 1.16
- Install [swagger](https://goswagger.io/install.html)
    - For Macs
        ```
        brew tap go-swagger/go-swagger
        brew install go-swagger
        ```
- Install [migrate](https://github.com/jackc/tern)
    - go get -u github.com/jackc/tern
- Install [Go Mock](https://github.com/golang/mock)
- Install [Postgres DB](https://www.postgresql.org/download/)
    - Once installed, create a DB called **feather-ai**
        - From the command line:
        ```
        psql
        CREATE DATABSE "feather-ai";
        ```
- Install [DBeaver](https://dbeaver.io/) - a tool for interacting with DBs
    - Once installed, setup a connection to the DB you just created.
        ```
        New Connection -> Postgres ->
        Host        = localhost
        Port        = 5432 (default)
        Database    = feather-ai
        Username    = postgres
        Password    = <blank>
        ```
- Install Docker

## Database

DB migrations use [tern](https://github.com/jackc/tern)

## Working with this code

To run the service locally, use **make run**. This will start the service listening on port 8080, and it will connect to the AWS account, and your Local Postgres DB. This allows you to access S3 and Lambdas, but safely via your own database.
For AWS access, ensure you have a profile named **featherai** in your ~/.aws/credentials file. Eg:

    [featherai]
    aws_access_key_id = <your key id>
    aws_secret_access_key = <your key>

Once the service is running, you can use Curl or Postman to send requests to the service. Requests still require authentication, but when running locally, *debug auth* is enabled. Set the X-AUTH0-TOKEN to *debug <some fake@email>* to the Login endpoint, which will either create or login a user matching the fake email.

## Deployments

Deployments are not hooked up to any CD system and are done manually for now. Also, the EC2 instance we're using is too small to allow >1 container and there's no ALB (so no dynamic port mappings). 
All commands that follow require a .local/cloud.env file (not checked in) which contains your secrets. So deployments are as follow:

Build and public the docker container to the ECR registry. This will publish **ftr-service-core:latest** image, overwriting the latest image in the registry.
    
    make deploy

If you've made any DB changes, you need to run Migrations:

    make migrate_cloud

At this point, the DB is ready and the image is published. Log into AWS, goto ECS and select our cluster. You need to create a new deployment, after deleting the current deployment.
Of course there will be down-time since we only have 1 instance for now.
