package database

import (
	"fmt"
	"golang_kafka/config"
	"golang_kafka/config/globals"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

var (
	// DBEndpoint DB Endpoint
	DBEndpoint = "http://localhost:8529/"
	// DBName database name
	DBName = config.DB
	// DBUsername database username
	DBUsername = config.DBUsername
	// DBPassword database password
	DBPassword = config.DBPassword
	// DB is the underlying database connection
	DB driver.Database
	// Client is the underlying client connection
	Client driver.Client
)

// Connect connect to database
func Connect() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBEndpoint},
	})

	if err != nil {
		fmt.Printf("ERROR : %v", err)
		panic(err)
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(DBUsername, DBPassword),
	})

	if err != nil {
		fmt.Println("[CLIENT]::UNAUTHENTICATED")
		panic(err)
	}

	Client = c

	db, err := c.Database(globals.Ctx, DBName)
	if err != nil {
		fmt.Println("[DATABASE]::NAME_ERROR")
		panic(err)
	}

	DB = db

	fmt.Println("[DATABASE]::CONNECTED")
}

// CheckCollection initialize undefined Collection on DB
func CheckCollection() {
	collectionChecking(globals.BranchColletion)
	collectionChecking(globals.AgentCollection)
	collectionChecking(globals.MerchantCollection)
	collectionChecking(globals.MerchantAdminCollection)
	collectionChecking(globals.UserCollection)
	collectionChecking(globals.RoleCollection)
	collectionChecking(globals.CustomerRegistrationCollection)
	collectionChecking(globals.ImageCollection)
}

func collectionChecking(collectionName string) {
	found, err := DB.CollectionExists(globals.Ctx, collectionName)
	if err != nil {
		fmt.Println("[COLLECTION]::CHECKING_ERROR")
		panic(err)
	}

	if !found {
		options := &driver.CreateCollectionOptions{ /* ... */ }
		_, err := DB.CreateCollection(globals.Ctx, collectionName, options)
		if err != nil {
			fmt.Println("[COLLECTION]::CREATE_ERROR")
			panic(err)
		}
		fmt.Println("[COLLECTION]::CREATED")
	}
}

// TestConnect untuk testing koneksi ke database
func TestConnect() (driver.Database, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBEndpoint},
	})

	if err != nil {
		//panic(err)
		fmt.Printf("ERROR : %v", err)
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(DBUsername, DBPassword),
	})

	if err != nil {
		//panic(err)
		fmt.Printf("ERROR : %v", err)
	}

	//ctx := context.Background()
	db, err := c.Database(globals.Ctx, DBName)
	if err != nil {
		//panic(err)
		fmt.Printf("ERROR : %v", err)
	}

	DB = db
	return db, err
}
