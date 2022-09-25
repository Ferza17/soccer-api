<h1>How To run</h1>

<h2>Soccer-API</h2>
- make sure you have installed go on your machine
-  open terminal, route to this project
- type : <br />

```
go run main.go soccer
```

<h4>Soccer-API Documentation via gqlgen playground</h4>
- open <a href="http://localhost:4000">http://localhost:4000</a> on your browser

<h4>Soccer-API Documentation via <a href="https://www.postman.com/downloads/">postman</a>  </h4>
- Make sure you have installed postman on your machine
- open postman
- import file  <br />

```
soccer-api.postman_collection.json
```

<h3>TEST Soccer-API</h3>
Please Follow The step : <br />
- hit endpoint <b>Create Player</b>
- hit endpoint <b>Create Team</b>
- hit endpoint <b>Add Player to team </b>
- hit endpoint <b>Find players</b> or hit endpoint <b>Find player</b>   <br />
    *Note it can be free players or team's player, depends on your input <b>isFreePlayer</b>

<h3>run unit test</h3>
- open terminal
- type : 

```
go test -v -race -short -failfast -cover ./...
```

- if you want grep coverage then type

```
go test -v -race -short -failfast -cover ./... && grep coverage
```