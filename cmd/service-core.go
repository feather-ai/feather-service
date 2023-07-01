package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"

	handlers "feather-ai/service-core/api"
	"feather-ai/service-core/api/generated/restapi"
	"feather-ai/service-core/api/generated/restapi/operations"
	"feather-ai/service-core/lib/aisystems"
	"feather-ai/service-core/lib/aisystems/aisystemscore"
	"feather-ai/service-core/lib/config"
	"feather-ai/service-core/lib/gatekeeper"
	"feather-ai/service-core/lib/gatekeeper/gatekeepercore"
	"feather-ai/service-core/lib/storage"
	"feather-ai/service-core/lib/storage/storagecore"
	"feather-ai/service-core/lib/upload"
	"feather-ai/service-core/lib/upload/uploadcore"
	"feather-ai/service-core/lib/users"
	"feather-ai/service-core/lib/users/userscore"
)

var gUploadManager upload.UploadManager
var gStorageManager storage.StorageManager
var gUsersManager users.UserManager
var gAISystemsManager aisystems.AISystemsManager
var gAnonymousGateKeeper gatekeeper.AnonGateKeeper

/*
* Parse all provided options from the environment variables
 */
func ParseOptions() config.Options {
	options := config.Options{}

	_, err := flags.Parse(&options)
	if err != nil {
		logrus.Fatalf("Could not parse arg: %v", err)
	}

	b, _ := json.Marshal(options)
	log.Printf("Options: %v\n", string(b))
	return options
}

/*
* Create all our manager interfaces. This creates the concrete classes to inject into all sub-modules
 */
func BindInterfaces(options config.Options) {
	gStorageManager = storagecore.NewStorageManager(options)

	gUsersManager = userscore.New(gStorageManager, options)
	gAISystemsManager = aisystemscore.New(gStorageManager, options)
	gUploadManager = uploadcore.New(gStorageManager, gAISystemsManager, options)
	gAnonymousGateKeeper = gatekeepercore.New(gStorageManager, options)
}

/*
* Start and run the server - this method doesn't return until the server shuts down.
 */
func RunServer(options config.Options, api *operations.CoreAPI) {
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "FeatherAI ServiceCore API"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

/*
* Entry point
 */
func main() {

	// Load options

	options := ParseOptions()

	// Boot up all the core systems

	BindInterfaces(options)

	// Setup our handlers

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCoreAPI(swaggerSpec)

	// For each of our custom API handlers, install them here...

	handlers.InstallCoreHandlers(api,
		swaggerSpec.Spec(),
		gUploadManager,
		gUsersManager,
		gAISystemsManager,
		gAnonymousGateKeeper)
	handlers.InstallUploadHandlers(api)
	handlers.InstallAiSystemHandlers(api)
	handlers.InstallUserHandlers(api)

	// Create and run the Server

	RunServer(options, api)
}
