package modeltests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/smhtkn/testpostgre/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllPosts(t *testing.T) {

	err := refreshClientAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing client and post table %v\n", err)
	}
	_, _, err = seedClientsAndPosts()
	if err != nil {
		log.Fatalf("Error seeding client and post  table %v\n", err)
	}
	posts, err := postInstance.FindAllPosts(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the posts: %v\n", err)
		return
	}
	assert.Equal(t, len(*posts), 2)
}

func TestSavePost(t *testing.T) {

	err := refreshClientAndPostTable()
	if err != nil {
		log.Fatalf("Error client and post refreshing table %v\n", err)
	}

	client, err := seedOneClient()
	if err != nil {
		log.Fatalf("Cannot seed client %v\n", err)
	}

	newPost := models.Post{
		ID:       1,
		Title:    "This is the title",
		Content:  "This is the content",
		AuthorID: client.ID,
	}
	savedPost, err := newPost.SavePost(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the post: %v\n", err)
		return
	}
	assert.Equal(t, newPost.ID, savedPost.ID)
	assert.Equal(t, newPost.Title, savedPost.Title)
	assert.Equal(t, newPost.Content, savedPost.Content)
	assert.Equal(t, newPost.AuthorID, savedPost.AuthorID)

}

func TestGetPostByID(t *testing.T) {

	err := refreshClientAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing client and post table: %v\n", err)
	}
	post, err := seedOneClientAndOnePost()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundPost, err := postInstance.FindPostByID(server.DB, post.ID)
	if err != nil {
		t.Errorf("this is the error getting one client: %v\n", err)
		return
	}
	assert.Equal(t, foundPost.ID, post.ID)
	assert.Equal(t, foundPost.Title, post.Title)
	assert.Equal(t, foundPost.Content, post.Content)
}

func TestUpdateAPost(t *testing.T) {

	err := refreshClientAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing client and post table: %v\n", err)
	}
	post, err := seedOneClientAndOnePost()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	postUpdate := models.Post{
		ID:       1,
		Title:    "modiUpdate",
		Content:  "modiupdate@gmail.com",
		AuthorID: post.AuthorID,
	}
	updatedPost, err := postUpdate.UpdateAPost(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the client: %v\n", err)
		return
	}
	assert.Equal(t, updatedPost.ID, postUpdate.ID)
	assert.Equal(t, updatedPost.Title, postUpdate.Title)
	assert.Equal(t, updatedPost.Content, postUpdate.Content)
	assert.Equal(t, updatedPost.AuthorID, postUpdate.AuthorID)
}

func TestDeleteAPost(t *testing.T) {

	err := refreshClientAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing client and post table: %v\n", err)
	}
	post, err := seedOneClientAndOnePost()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}
	isDeleted, err := postInstance.DeleteAPost(server.DB, post.ID, post.AuthorID)
	if err != nil {
		t.Errorf("this is the error updating the client: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}
