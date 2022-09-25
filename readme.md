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
- hit endpoint __Create Player__  <br />
- hit endpoint __Create Team__ <br />
- hit endpoint __Add Player to team__ <br />
- hit endpoint __Find players__ or hit endpoint __Find player__   <br />
    *Note it can be free players or team's player, depends on your input __isFreePlayer__

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