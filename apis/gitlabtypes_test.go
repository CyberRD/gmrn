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

func TestDeJsonProject(t *testing.T) {
	assert := assert.New(t)
	jsonstr := `
{
  "id": 3,
  "description": null,
  "default_branch": "master",
  "public": false,
  "visibility_level": 0,
  "ssh_url_to_repo": "git@example.com:diaspora/diaspora-project-site.git",
  "http_url_to_repo": "http://example.com/diaspora/diaspora-project-site.git",
  "web_url": "http://example.com/diaspora/diaspora-project-site",
  "tag_list": [
    "example",
    "disapora project"
  ],
  "owner": {
    "id": 3,
    "name": "Diaspora",
    "created_at": "2013-09-30T13: 46: 02Z"
  },
  "name": "Diaspora Project Site",
  "name_with_namespace": "Diaspora / Diaspora Project Site",
  "path": "diaspora-project-site",
  "path_with_namespace": "diaspora/diaspora-project-site",
  "issues_enabled": true,
  "merge_requests_enabled": true,
  "wiki_enabled": true,
  "snippets_enabled": false,
  "created_at": "2013-09-30T13: 46: 02Z",
  "last_activity_at": "2013-09-30T13: 46: 02Z",
  "creator_id": 3,
  "namespace": {
    "created_at": "2013-09-30T13: 46: 02Z",
    "description": "",
    "id": 3,
    "name": "Diaspora",
    "owner_id": 1,
    "path": "diaspora",
    "updated_at": "2013-09-30T13: 46: 02Z"
  },
  "permissions": {
    "project_access": {
      "access_level": 10,
      "notification_level": 3
    },
    "group_access": {
      "access_level": 50,
      "notification_level": 3
    }
  },
  "archived": false,
  "avatar_url": "http://example.com/uploads/project/avatar/3/uploads/avatar.png"
}
	`
	project := Poroject{}
	json.Unmarshal([]byte(jsonstr), &project)
	assert.Equal(project.Id, 3.0)
	assert.Equal(project.PathWithNamespace, "diaspora/diaspora-project-site")

}
