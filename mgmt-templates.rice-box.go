package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file4 := &embedded.EmbeddedFile{
		Filename:    "body.tmpl",
		FileModTime: time.Unix(1521654310, 0),
		Content:     string("<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <title>CUE Vision LFS Server Management</title>\n    <style type=\"text/css\">\n      @import \"/mgmt/css/primer.css\";\n      .masthead{\n        padding-top:1rem;\n        padding-bottom:1rem;\n        margin-bottom:1.5rem;\n        background-color:#4183c4;\n        color: white;\n      }\n      td {\n        padding-right: 1rem;\n        padding-bottom: 1rem;\n      }\n    </style>\n  </head>\n  <body>\n    <header class=\"masthead\">\n    <div class=\"container\">\n      <h1>CUE Vision LFS Server</h1>\n    </div>\n    </header>\n\n    <div class=\"container\">\n      <div class=\"columns\">\n        <div class=\"one-fourth column\">\n          <nav class=\"menu\">\n            <a class=\"menu-item {{if eq .Name \"index\"}}selected{{end}}\" href=\"/mgmt\">LFS Server</a>\n            <a class=\"menu-item {{if eq .Name \"users\"}}selected{{end}}\" href=\"/mgmt/users\">Users</a>\n            <a class=\"menu-item {{if eq .Name \"objects\"}}selected{{end}}\" href=\"/mgmt/objects\">Objects</a>\n            <a class=\"menu-item {{if eq .Name \"locks\"}}selected{{end}}\" href=\"/mgmt/locks\">Locks</a>\n          </nav>\n        </div>\n        <div class=\"three-fourths column\">\n          {{template \"content\" .}}\n      </div>\n    </div>\n  </body>\n</html>\n"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "config.tmpl",
		FileModTime: time.Unix(1521653224, 0),
		Content:     string("<div class=\"container\">\n  <p><strong>URL:</strong> {{.Config.Scheme}}://{{.Config.Host}}</p>\n  <p><strong>Listen Address:</strong> {{.Config.Listen}}</p>\n  <p><strong>Database:</strong> {{.Config.MetaDB}}</p>\n  <p><strong>Content:</strong> {{.Config.ContentPath}}</p>\n</div>\n<div class=\"container\">\n  <p>To configure a repository to use this LFS server, add the following to the repository's <code>.gitconfig</code> file:</p>\n  <pre>\n<code>[lfs]\n    url = \"{{.Config.Scheme}}://{{.Config.Host}}/\"\n</code>\n</pre>\n\n{{if eq .Config.Scheme \"https\"}}\n<p>Your server is configured to use https. If you're using self signed certificates, or are getting SSL errors, you can add the following to your <code>.gitconfig</code> file:</p>\n<pre>\n<code>[http]\n    sslverify = false\n</code>\n</pre>\n{{end}}\n</div>\n"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "locks.tmpl",
		FileModTime: time.Unix(1521653224, 0),
		Content:     string("<div class=\"container\">\n  <table>\n    <tr>\n      <th>ID</th>\n      <th>Path</th>\n      <th>Owner</th>\n      <th>LockedAt</th>\n    </tr>\n    {{range .Locks}}\n      <tr>\n        <td>{{.Id}}</td>\n        <td>{{.Path}}</td>\n        <td>{{.Owner.Name}}</td>\n        <td>{{.LockedAt.Format \"2006-01-02 15:04:05\"}}</td>\n      </tr>\n    {{end}}\n  </table>\n</div>\n"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "objects.tmpl",
		FileModTime: time.Unix(1521653224, 0),
		Content:     string("<div class=\"container\">\n  <table>\n    <tr>\n      <th>OID</th>\n      <th>Size</th>\n    </tr>\n    {{range .Objects}}\n      <tr>\n        <td><a target=\"_blank\" href=\"/mgmt/raw/{{.Oid}}\">{{.Oid}}</a></td>\n        <td>{{.Size}}</td>\n      </tr>\n    {{end}}\n  </table>\n</div>\n"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "users.tmpl",
		FileModTime: time.Unix(1521653224, 0),
		Content:     string("<div class=\"container\">\n  <table>\n    {{range .Users}}\n      <tr>\n        <td>{{.Name}}</td>\n        <td><form method=\"POST\" action=\"/mgmt/del\"><input type=\"hidden\" name=\"name\" value=\"{{.Name}}\"/><button type=\"submit\" class=\"btn btn-sm btn-danger\">Remove</button></form></td>\n      </tr>\n    {{end}}\n  </table>\n</div>\n<div class=\"container\">\n  <form method=\"POST\" action=\"/mgmt/add\">\n    <input type=\"text\" name=\"name\" placeholder=\"Username\">\n    <input type=\"password\" name=\"password\" placeholder=\"Password\">\n    <button type=\"submit\" class=\"btn\">Add User</button>\n  </form>\n</div>\n"),
	}

	// define dirs
	dir3 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1497862801, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file4, // body.tmpl
			file5, // config.tmpl
			file6, // locks.tmpl
			file7, // objects.tmpl
			file8, // users.tmpl

		},
	}

	// link ChildDirs
	dir3.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`mgmt/templates`, &embedded.EmbeddedBox{
		Name: `mgmt/templates`,
		Time: time.Unix(1497862811, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir3,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"body.tmpl":    file4,
			"config.tmpl":  file5,
			"locks.tmpl":   file6,
			"objects.tmpl": file7,
			"users.tmpl":   file8,
		},
	})
}
