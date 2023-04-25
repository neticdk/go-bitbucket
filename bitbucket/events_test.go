package bitbucket

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseRepoPushEvent(t *testing.T) {
	ts, _ := time.Parse(time.RFC3339, "2023-01-13T22:26:25+11:00")

	var ev RepositoryPushEvent
	err := json.Unmarshal([]byte(repoPushEvent), &ev)
	assert.NoError(t, err)
	assert.Equal(t, EventKeyRepoRefsChanged, ev.EventKey)
	assert.Equal(t, ISOTime(ts), ev.Date)
	assert.Equal(t, "rep_1", ev.Repository.Slug)
	assert.Equal(t, "a00945762949b7b787ecabc388c0e20b1b85f0b4", ev.Commits[0].ID)
	assert.Equal(t, "My commit message", ev.Commits[0].Message)
	assert.Equal(t, "197a3e0d2f9a2b3ed1c4fe5923d5dd701bee9fdd", ev.Commits[0].Parents[0].ID)
	assert.Equal(t, "a00945762949b7b787ecabc388c0e20b1b85f0b4", ev.Commits[0].ID)
	assert.Equal(t, "a00945762949b7b787ecabc388c0e20b1b85f0b4", ev.ToCommit.ID)
	assert.Len(t, ev.Changes, 1)
	assert.Equal(t, RepositoryPushEventChangeTypeUpdate, ev.Changes[0].Type)
	assert.Equal(t, RepositoryPushEventRefTypeBranch, ev.Changes[0].Ref.Type)
}

func TestParsePROpenedEvent(t *testing.T) {
	var ev PullRequestEvent
	err := json.Unmarshal([]byte(prOpened), &ev)
	assert.NoError(t, err)
	assert.Equal(t, "ef8755f06ee4b28c96a847a95cb8ec8ed6ddd1ca", ev.PullRequest.Source.Latest)
	assert.Equal(t, "178864a7d521b6f5e720b386b2c2b0ef8563e0dc", ev.PullRequest.Target.Latest)
	assert.Equal(t, "admin", ev.PullRequest.Author.Author.Name)
}

func TestParsePRSourceChangedEvent(t *testing.T) {
	var ev PullRequestEvent
	err := json.Unmarshal([]byte(prSourceChange), &ev)
	assert.NoError(t, err)
	assert.Equal(t, "aab847db240ccae221f8036605b00f777eba95d2", ev.PullRequest.Source.Latest)
	assert.Equal(t, "86448735f9dee9e1fb3d3e5cd9fbc8eb9d8400f4", ev.PullRequest.Target.Latest)
	assert.Equal(t, "admin", ev.PullRequest.Author.Author.Name)
}

const repoPushEvent = `{
	"eventKey": "repo:refs_changed",
	"date": "2023-01-13T22:26:25+1100",
	"actor": {
	  "name": "admin",
	  "emailAddress": "admin@example.com",
	  "active": true,
	  "displayName": "Administrator",
	  "id": 2,
	  "slug": "admin",
	  "type": "NORMAL"
	},
	"repository": {
		"slug": "rep_1",
		"id": 1,
		"name": "rep_1",
		"hierarchyId": "af05451fc6eb4bf4e0bd",
		"scmId": "git",
		"state": "AVAILABLE",
		"statusMessage": "Available",
		"forkable": true,
		"project": {
			"key": "PROJECT_1",
			"id": 1,
			"name": "Project 1",
			"description": "PROJECT_1",
			"public": false,
			"type": "NORMAL"
		},
		"public": false,
		"archived": false
	},
	"changes": [
		{
		"ref": {
			"id": "refs/heads/master",
			"displayId": "master",
			"type": "BRANCH"
		},
		"refId": "refs/heads/master",
		"fromHash": "197a3e0d2f9a2b3ed1c4fe5923d5dd701bee9fdd",
		"toHash": "a00945762949b7b787ecabc388c0e20b1b85f0b4",
		"type": "UPDATE"
		}
	],
	"commits": [
		{
		"id": "a00945762949b7b787ecabc388c0e20b1b85f0b4",
		"displayId": "a0094576294",
		"author": {
			"name": "Administrator",
			"emailAddress": "admin@example.com"
		},
		"authorTimestamp": 1673403328000,
		"committer": {
			"name": "Administrator",
			"emailAddress": "admin@example.com"
		},
		"committerTimestamp": 1673403328000,
		"message": "My commit message",
		"parents": [
			{
				"id": "197a3e0d2f9a2b3ed1c4fe5923d5dd701bee9fdd",
				"displayId": "197a3e0d2f9"
			}
		]
		}
	],
	"toCommit": {
		"id": "a00945762949b7b787ecabc388c0e20b1b85f0b4",
		"displayId": "a0094576294",
		"author": {
			"name": "Administrator",
			"emailAddress": "admin@example.com"
		},
		"authorTimestamp": 1673403328000,
		"committer": {
			"name": "Administrator",
			"emailAddress": "admin@example.com"
		},
		"committerTimestamp": 1673403328000,
		"message": "My commit message",
		"parents": [
			{
				"id": "197a3e0d2f9a2b3ed1c4fe5923d5dd701bee9fdd",
				"displayId": "197a3e0d2f9",
				"author": {
					"name": "Administrator",
					"emailAddress": "admin@example.com"
				},
				"authorTimestamp": 1673403292000,
				"committer": {
					"name": "Administrator",
					"emailAddress": "admin@example.com"
				},
				"committerTimestamp": 1673403292000,
				"message": "My commit message",
				"parents": [
					{
					"id": "f870ce6bf6fe633e1a2bbe655970bde25535669f",
					"displayId": "f870ce6bf6f"
					}
				]
			}
		]
	}
}`

const prOpened = `{  
	"eventKey":"pr:opened",
	"date":"2017-09-19T09:58:11+1000",
	"actor":{  
	  "name":"admin",
	  "emailAddress":"admin@example.com",
	  "id":1,
	  "displayName":"Administrator",
	  "active":true,
	  "slug":"admin",
	  "type":"NORMAL"
	},
	"pullRequest":{  
	  "id":1,
	  "version":0,
	  "title":"a new file added",
	  "state":"OPEN",
	  "open":true,
	  "closed":false,
	  "createdDate":1505779091796,
	  "updatedDate":1505779091796,
	  "fromRef":{  
		"id":"refs/heads/a-branch",
		"displayId":"a-branch",
		"latestCommit":"ef8755f06ee4b28c96a847a95cb8ec8ed6ddd1ca",
		"repository":{  
		  "slug":"repository",
		  "id":84,
		  "name":"repository",
		  "scmId":"git",
		  "state":"AVAILABLE",
		  "statusMessage":"Available",
		  "forkable":true,
		  "project":{  
			"key":"PROJ",
			"id":84,
			"name":"project",
			"public":false,
			"type":"NORMAL"
		  },
		  "public":false
		}
	  },
	  "toRef":{  
		"id":"refs/heads/master",
		"displayId":"master",
		"latestCommit":"178864a7d521b6f5e720b386b2c2b0ef8563e0dc",
		"repository":{  
		  "slug":"repository",
		  "id":84,
		  "name":"repository",
		  "scmId":"git",
		  "state":"AVAILABLE",
		  "statusMessage":"Available",
		  "forkable":true,
		  "project":{  
			"key":"PROJ",
			"id":84,
			"name":"project",
			"public":false,
			"type":"NORMAL"
		  },
		  "public":false
		}
	  },
	  "locked":false,
	  "author":{  
		"user":{  
		  "name":"admin",
		  "emailAddress":"admin@example.com",
		  "id":1,
		  "displayName":"Administrator",
		  "active":true,
		  "slug":"admin",
		  "type":"NORMAL"
		},
		"role":"AUTHOR",
		"approved":false,
		"status":"UNAPPROVED"
	  },
	  "reviewers":[  
  
	  ],
	  "participants":[  
  
	  ],
	  "links":{  
		"self":[  
		  null
		]
	  }
	}
  }`

const prSourceChange = `{
	"eventKey": "pr:from_ref_updated",
	"date": "2020-02-20T14:49:41+1100",
	"actor": {
	  "name": "admin",
	  "emailAddress": "admin@example.com",
	  "id": 1,
	  "displayName": "Administrator",
	  "active": true,
	  "slug": "admin",
	  "type": "NORMAL",
	  "links": {
		"self": [
		  {
			"href": "http://localhost:7990/bitbucket/users/admin"
		  }
		]
	  }
	},
	"pullRequest": {
	  "id": 2,
	  "version": 16,
	  "title": "Webhook",
	  "state": "OPEN",
	  "open": true,
	  "closed": false,
	  "createdDate": 1582065825700,
	  "updatedDate": 1582170581372,
	  "fromRef": {
		"id": "refs/heads/pr-webhook",
		"displayId": "pr-webhook",
		"latestCommit": "aab847db240ccae221f8036605b00f777eba95d2",
		"repository": {
		  "slug": "dvcs",
		  "id": 33,
		  "name": "dvcs",
		  "hierarchyId": "09992c6ad9e001f01120",
		  "scmId": "git",
		  "state": "AVAILABLE",
		  "statusMessage": "Available",
		  "forkable": true,
		  "project": {
			"key": "GIT",
			"id": 62,
			"name": "Bitbucket",
			"public": false,
			"type": "NORMAL",
			"links": {
			  "self": [
				{
				  "href": "http://localhost:7990/bitbucket/projects/GIT"
				}
			  ]
			}
		  },
		  "public": false,
		  "links": {
			"clone": [
			  {
				"href": "ssh://git@localhost:7999/git/dvcs.git",
				"name": "ssh"
			  },
			  {
				"href": "http://localhost:7990/bitbucket/scm/git/dvcs.git",
				"name": "http"
			  }
			],
			"self": [
			  {
				"href": "http://localhost:7990/bitbucket/projects/GIT/repos/dvcs/browse"
			  }
			]
		  }
		}
	  },
	  "toRef": {
		"id": "refs/heads/master",
		"displayId": "master",
		"latestCommit": "86448735f9dee9e1fb3d3e5cd9fbc8eb9d8400f4",
		"repository": {
		  "slug": "dvcs",
		  "id": 33,
		  "name": "dvcs",
		  "hierarchyId": "09992c6ad9e001f01120",
		  "scmId": "git",
		  "state": "AVAILABLE",
		  "statusMessage": "Available",
		  "forkable": true,
		  "project": {
			"key": "GIT",
			"id": 62,
			"name": "Bitbucket",
			"public": false,
			"type": "NORMAL",
			"links": {
			  "self": [
				{
				  "href": "http://localhost:7990/bitbucket/projects/GIT"
				}
			  ]
			}
		  },
		  "public": false,
		  "links": {
			"clone": [
			  {
				"href": "ssh://git@localhost:7999/git/dvcs.git",
				"name": "ssh"
			  },
			  {
				"href": "http://localhost:7990/bitbucket/scm/git/dvcs.git",
				"name": "http"
			  }
			],
			"self": [
			  {
				"href": "http://localhost:7990/bitbucket/projects/GIT/repos/dvcs/browse"
			  }
			]
		  }
		}
	  },
	  "locked": false,
	  "author": {
		"user": {
		  "name": "admin",
		  "emailAddress": "admin@example.com",
		  "id": 1,
		  "displayName": "Administrator",
		  "active": true,
		  "slug": "admin",
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "http://localhost:7990/bitbucket/users/admin"
			  }
			]
		  }
		},
		"role": "AUTHOR",
		"approved": false,
		"status": "UNAPPROVED"
	  },
	  "reviewers": [],
	  "participants": [],
	  "links": {
		"self": [
		  {
			"href": "http://localhost:7990/bitbucket/projects/GIT/repos/dvcs/pull-requests/2"
		  }
		]
	  }
	},
	"previousFromHash": "99f3ea32043ba3ecaa28de6046b420de70257d80"
  }`
