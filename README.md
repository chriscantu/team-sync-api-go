Team Sync
===========

Simple Go / Gin Framework API application designed around keeping teams up-to-date.

## Disclaimer
I am very new to Go and would not recommend my logic as an example of how to do things in Go or Gin. (You have been warned).

## Dependencies
* Rethink DB installed locally
* github.com/gin-gonic/gin
* github.com/dancannon/gorethink

## RethinkDB Setup
* Install RethinkDB locally
* Create db 'teamsync'
* Create Tables 'users', 'questions', 'reports', 'teams'

## Initial Run
* `git clone git@github.com:chriscantu/team-sync.git`
* `go get github.com/gin-gonic/gin`
* `go get github.com/dancannon/gorethink`
* `cd team-sync`
* `go run server.go`
