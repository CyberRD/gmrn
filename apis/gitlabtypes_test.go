package apis

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeJosnAuthor(t *testing.T) {
	assert := assert.New(t)
	jsonstr := `{
		"name": "Chad Hamill",
		"username": "jarrett",
		"id": 5,
		"state": "active",
		"avatar_url": "http://www.gravatar.com/avatar/b95567800f828948baf5f4160ebb2473?s=40&d=identicon"
	}`
	author := Author{}
	json.Unmarshal([]byte(jsonstr), &author)
	assert.Equal(author.Name, "Chad Hamill")
	assert.Equal(author.Id, 5.0)

}

func TestDeJosnMergeRequest(t *testing.T) {
	assert := assert.New(t)
	jsonstr := `{
  "id": 1,
  "iid": 1,
  "target_branch": "master",
  "source_branch": "test1",
  "project_id": 3,
  "title": "test1",
  "state": "merged",
  "upvotes": 0,
  "downvotes": 0,
  "author": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "name": "Administrator",
    "state": "active",
    "created_at": "2012-04-29T08:46:00Z"
  },
  "assignee": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "name": "Administrator",
    "state": "active",
    "created_at": "2012-04-29T08:46:00Z"
  },
  "description":"fixed login page css paddings",
  "work_in_progress": false
}
	`
	mr := MergeRequest{}
	json.Unmarshal([]byte(jsonstr), &mr)
	assert.Equal(mr.Id, 1.0)
	assert.Equal(mr.Iid, 1.0)
	assert.Equal(mr.Author.Name, "Administrator")
}
