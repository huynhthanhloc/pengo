# pengo [![GoDoc](https://img.shields.io/badge/doc-reference-blue.svg)](https://godoc.org/github.com/penlook/pengo) [![Build Status](https://travis-ci.org/penlook/pengo.svg)](https://travis-ci.org/penlook/pengo) [![Build status](https://ci.appveyor.com/api/projects/status/u6m54q5v1tgl9sxh?svg=true)](https://ci.appveyor.com/project/loint/pengo) [![Coverage Status](https://coveralls.io/repos/penlook/pengo/badge.svg?branch=master)](https://coveralls.io/r/penlook/pengo?branch=master)



![Go fly with Pengo](doc/image/gofly.png)

###Taste of Pengo (technical preview)

```go
@Controller App

@Route /login/:uid/:password
@Method GET
@Pick "Login Form"
func Login(uid int, password string) {

	// Using table
	user := #User {
		Id: uid,
		Password: password,
	}

	// Object relational mapping
	me := user.First()

	// Pass my information to view
	$name  = me.Name
	$email = me.Email

	// Model of current controller
	// Pass list user to view
	$users = @ListUser()
}
```

### Pengo syntax
- You can use Go mixed with Pengo syntax to create web application faster, more meaningful
- Eliminating the complexity of the Go syntax and deep integrated with framework
- All of your code will be compiled into Go as you wrote it.

###Roadmap

- MVC Architecture (testing)
- Template engine (testing)
- Hot-code reload (not started)
- Improve performance (implementing)
- Command line for development (implementing)
- Model engine (not started)
- Unit - Integration test structure (implementing)
- Multiple languages (implementing)
- Documentation (not started)
- Flow Tracking (testing)
- Extend framework by using C (implementing)
- Annotation Parser (not started)
- Rest API (not started)
- Pengo website (not started)

####Support Database

- SQLite	  -   Table      - ORM  (implementing)
- MySQL      -   Table      - ORM  (testing)
- Cassandra  -   Table      - CQL  (not started)
- MongoDB    -   Document   - ODM  (implemeting)
- CouchDB    -   Document          (not started)
- Memcache   -   Key-value         (not started)
- Redis      -   Key-value         (implemeting)
- Cayley     -   Graph             (not started)

###Contributors
We look forward to your pull requests. If you would like to be the contributor please accept some rules.

- The pull requests will be accepted only in "development" branch
- All modifications or additions should be tested

Thank you for your understanding.
Open a pull request, we are waiting.

###License
This is open source project for community development under GNU Affero General Public License

Authors
	- Loi Nguyen <loint@penlook.com>
	- and contributors





