// Write your event here

package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/energye/golcl/lcl/types"

	"github.com/energye/golcl/lcl"
)

//::private::
type TMainFormFields struct {
}

const jsonData = `{
  "id": 104690469,
  "node_id": "MDEwOlJlcG9zaXRvcnkxMDQ2OTA0Njk=",
  "name": "golcl",
  "full_name": "sxm/golcl",
  "private": false,
  "owner": {
    "login": "sxm",
    "id": 7037165,
    "node_id": "MDQ6VXNlcjcwMzcxNjU=",
    "avatar_url": "https://avatars1.githubusercontent.com/u/7037165?v=4",
    "gravatar_id": "",
    "url": "https://api.gitee.com/users/sxm",
    "html_url": "https://github.com/energye",
    "followers_url": "https://api.gitee.com/users/sxm/followers",
    "following_url": "https://api.gitee.com/users/sxm/following{/other_user}",
    "gists_url": "https://api.gitee.com/users/sxm/gists{/gist_id}",
    "starred_url": "https://api.gitee.com/users/sxm/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.gitee.com/users/sxm/subscriptions",
    "organizations_url": "https://api.gitee.com/users/sxm/orgs",
    "repos_url": "https://api.gitee.com/users/sxm/repos",
    "events_url": "https://api.gitee.com/users/sxm/events{/privacy}",
    "received_events_url": "https://api.gitee.com/users/sxm/received_events",
    "type": "User",
    "site_admin": false
  },
  "html_url": "https://github.com/energye/golcl",
  "description": " A cross-platform Golang GUI library. Use Delphi lcl and Lazarus LCL for binding. ",
  "fork": false,
  "url": "https://api.gitee.com/repos/sxm/golcl",
  "forks_url": "https://api.gitee.com/repos/sxm/golcl/forks",
  "keys_url": "https://api.gitee.com/repos/sxm/golcl/keys{/key_id}",
  "collaborators_url": "https://api.gitee.com/repos/sxm/golcl/collaborators{/collaborator}",
  "teams_url": "https://api.gitee.com/repos/sxm/golcl/teams",
  "hooks_url": "https://api.gitee.com/repos/sxm/golcl/hooks",
  "issue_events_url": "https://api.gitee.com/repos/sxm/golcl/issues/events{/number}",
  "events_url": "https://api.gitee.com/repos/sxm/golcl/events",
  "assignees_url": "https://api.gitee.com/repos/sxm/golcl/assignees{/user}",
  "branches_url": "https://api.gitee.com/repos/sxm/golcl/branches{/branch}",
  "tags_url": "https://api.gitee.com/repos/sxm/golcl/tags",
  "blobs_url": "https://api.gitee.com/repos/sxm/golcl/git/blobs{/sha}",
  "git_tags_url": "https://api.gitee.com/repos/sxm/golcl/git/tags{/sha}",
  "git_refs_url": "https://api.gitee.com/repos/sxm/golcl/git/refs{/sha}",
  "trees_url": "https://api.gitee.com/repos/sxm/golcl/git/trees{/sha}",
  "statuses_url": "https://api.gitee.com/repos/sxm/golcl/statuses/{sha}",
  "languages_url": "https://api.gitee.com/repos/sxm/golcl/languages",
  "stargazers_url": "https://api.gitee.com/repos/sxm/golcl/stargazers",
  "contributors_url": "https://api.gitee.com/repos/sxm/golcl/contributors",
  "subscribers_url": "https://api.gitee.com/repos/sxm/golcl/subscribers",
  "subscription_url": "https://api.gitee.com/repos/sxm/golcl/subscription",
  "commits_url": "https://api.gitee.com/repos/sxm/golcl/commits{/sha}",
  "git_commits_url": "https://api.gitee.com/repos/sxm/golcl/git/commits{/sha}",
  "comments_url": "https://api.gitee.com/repos/sxm/golcl/comments{/number}",
  "issue_comment_url": "https://api.gitee.com/repos/sxm/golcl/issues/comments{/number}",
  "contents_url": "https://api.gitee.com/repos/sxm/golcl/contents/{+path}",
  "compare_url": "https://api.gitee.com/repos/sxm/golcl/compare/{base}...{head}",
  "merges_url": "https://api.gitee.com/repos/sxm/golcl/merges",
  "archive_url": "https://api.gitee.com/repos/sxm/golcl/{archive_format}{/ref}",
  "downloads_url": "https://api.gitee.com/repos/sxm/golcl/downloads",
  "issues_url": "https://api.gitee.com/repos/sxm/golcl/issues{/number}",
  "pulls_url": "https://api.gitee.com/repos/sxm/golcl/pulls{/number}",
  "milestones_url": "https://api.gitee.com/repos/sxm/golcl/milestones{/number}",
  "notifications_url": "https://api.gitee.com/repos/sxm/golcl/notifications{?since,all,participating}",
  "labels_url": "https://api.gitee.com/repos/sxm/golcl/labels{/name}",
  "releases_url": "https://api.gitee.com/repos/sxm/golcl/releases{/id}",
  "deployments_url": "https://api.gitee.com/repos/sxm/golcl/deployments",
  "created_at": "2017-09-25T01:38:08Z",
  "updated_at": "2018-11-26T01:29:27Z",
  "pushed_at": "2018-11-30T06:15:04Z",
  "git_url": "git://github.com/energye/golcl.git",
  "ssh_url": "git@gitee.com:sxm/golcl.git",
  "clone_url": "https://github.com/energye/golcl.git",
  "svn_url": "https://github.com/energye/golcl",
  "homepage": "",
  "size": 9103,
  "stargazers_count": 132,
  "watchers_count": 132,
  "language": "Go",
  "has_issues": true,
  "has_projects": true,
  "has_downloads": true,
  "has_wiki": true,
  "has_pages": false,
  "forks_count": 15,
  "mirror_url": null,
  "archived": false,
  "open_issues_count": 0,
  "license": {
    "key": "apache-2.0",
    "name": "Apache License 2.0",
    "spdx_id": "Apache-2.0",
    "url": "https://api.gitee.com/licenses/apache-2.0",
    "node_id": "MDc6TGljZW5zZTI="
  },
  "forks": 15,
  "open_issues": 0,
  "watchers": 132,
  "default_branch": "master",
  "network_count": 15,
  "subscribers_count": 20
}`

const jsonData2 = `{
   "key":1,
   "key2":[
   {"aaa":"aaaa", "a122":"a1"}, 
   {"bbb": "bbb"}
  ],
  "key3": [
     1, 2,3
   ],
     "key4": [
     "str1", "str2"
   ]
}
`

// https://api.gitee.com/repos/sxm/golcl

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {

}

func (f *TMainForm) jsonTree(str string) {
	var data interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		lcl.ShowMessage(err.Error())
		return
	}
	f.TreeView1.Items().BeginUpdate()
	defer f.TreeView1.Items().EndUpdate()
	root := f.TreeView1.Items().AddChild(nil, "JSON")
	f.buildTree(root, data, "")
	f.TreeView1.FullExpand()
}

func (f *TMainForm) buildTree(node *lcl.TTreeNode, data interface{}, keyName string) {

	switch data.(type) {
	case map[string]interface{}:

		var child *lcl.TTreeNode
		if keyName != "" {
			child = f.TreeView1.Items().AddChild(node, keyName)
			child = f.TreeView1.Items().AddChild(child, "Object")
		} else {
			child = f.TreeView1.Items().AddChild(node, "Object")
		}

		// ???????????????????????????????????????????????????????????????????????????
		keys := make([]string, 0)
		for key := range data.(map[string]interface{}) {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			if val, ok := data.(map[string]interface{})[key]; ok {
				f.buildTree(child, val, key)
			}
		}

	case []interface{}:
		var child *lcl.TTreeNode
		if keyName != "" {
			child = f.TreeView1.Items().AddChild(node, keyName)
			child = f.TreeView1.Items().AddChild(child, "Array")
		} else {
			child = f.TreeView1.Items().AddChild(node, "Array")
		}
		for _, val := range data.([]interface{}) {
			f.buildTree(child, val, "")
		}

	default:

		if node != nil && node.IsValid() {
			if keyName == "" {
				f.TreeView1.Items().AddChild(node, fmt.Sprintf("%v", data))
			} else {
				f.TreeView1.Items().AddChild(node, fmt.Sprintf("%s:%v", keyName, data))
			}
		}
	}
}

func (f *TMainForm) OnBtnPasteClick(sender lcl.IObject) {
	f.TreeView1.Items().Clear()
	if lcl.Clipboard.HasFormat(types.CF_TEXT) {
		f.jsonTree(lcl.Clipboard.AsText())
	}
}
