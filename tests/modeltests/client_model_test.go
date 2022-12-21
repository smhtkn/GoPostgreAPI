package modeltests

import (
	"log"
	"testing"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"github.com/smhtkn/testpostgre/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllClients(t *testing.T) {

	err := refreshClientTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedClients()
	if err != nil {
		log.Fatal(err)
	}

	clients, err := clientInstance.FindAllClients(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the clients: %v\n", err)
		return
	}
	assert.Equal(t, len(*clients), 2)
}

func TestSaveClient(t *testing.T) {

	err := refreshClientTable()
	if err != nil {
		log.Fatal(err)
	}
	newClient := models.Client{
		ID:       1,
		Email:    "test@gmail.com",
		Nickname: "test",
		Password: "password",
	}
	savedClient, err := newClient.SaveClient(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the clients: %v\n", err)
		return
	}
	assert.Equal(t, newClient.ID, savedClient.ID)
	assert.Equal(t, newClient.Email, savedClient.Email)
	assert.Equal(t, newClient.Nickname, savedClient.Nickname)
}

func TestGetClientByID(t *testing.T) {

	err := refreshClientTable()
	if err != nil {
		log.Fatal(err)
	}

	client, err := seedOneClient()
	if err != nil {
		log.Fatalf("cannot seed clients table: %v", err)
	}
	foundClient, err := clientInstance.FindClientByID(server.DB, client.ID)
	if err != nil {
		t.Errorf("this is the error getting one client: %v\n", err)
		return
	}
	assert.Equal(t, foundClient.ID, client.ID)
	assert.Equal(t, foundClient.Email, client.Email)
	assert.Equal(t, foundClient.Nickname, client.Nickname)
}

func TestUpdateAClient(t *testing.T) {

	err := refreshClientTable()
	if err != nil {
		log.Fatal(err)
	}

	client, err := seedOneClient()
	if err != nil {
		log.Fatalf("Cannot seed client: %v\n", err)
	}

	clientUpdate := models.Client{
		ID:       1,
		Nickname: "modiUpdate",
		Email:    "modiupdate@gmail.com",
		Password: "password",
	}
	updatedClient, err := clientUpdate.UpdateAClient(server.DB, client.ID)
	if err != nil {
		t.Errorf("this is the error updating the client: %v\n", err)
		return
	}
	assert.Equal(t, updatedClient.ID, clientUpdate.ID)
	assert.Equal(t, updatedClient.Email, clientUpdate.Email)
	assert.Equal(t, updatedClient.Nickname, clientUpdate.Nickname)
}

func TestDeleteAClient(t *testing.T) {

	err := refreshClientTable()
	if err != nil {
		log.Fatal(err)
	}

	client, err := seedOneClient()

	if err != nil {
		log.Fatalf("Cannot seed client: %v\n", err)
	}

	isDeleted, err := clientInstance.DeleteAClient(server.DB, client.ID)
	if err != nil {
		t.Errorf("this is the error updating the client: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}
