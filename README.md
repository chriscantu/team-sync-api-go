gin-example
===========

Simple Go / Gin Framework example in attempt to learn programming in Go.  

## Disclaimer
I am very new to Go and would not recommend my logic as an example of how to do things in Go or Gin. (You have been warned).

## Dependencies
* Rethink DB installed locally
* github.com/gin-gonic/gin
* github.com/dancannon/gorethink

## RethinkDB Setup
* Install RethinkDB locally
* Create db 'teamsync'
* Create Table 'users'
* Create Table 'questions'

## Initial Run
* `git clone git@github.com:chriscantu/gin-example.git`
* `go get github.com/gin-gonic/gin`
* `go get github.com/dancannon/gorethink`
* `cd gin-example`
* `go run server.go`
