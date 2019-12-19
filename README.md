# go-here-map
Go service to communicate with here.com's Location Service v1.6 API

Grab your HERE Location Services REST API App ID and App Code here
https://developer.here.com/projects

This service does not work with the following request parameters: ```apiKey``` ```bbox``` ```ctr``` ```e``` ```ectr``` ```ml2``` ```poifc``` ```poilbl``` ```poithm``` ```poitxc``` ```poitxs``` ```pview``` ```ra``` ```rad,rad0,rad1,...``` ```scp``` ```token``` ```tx.xy,tx.xy0,tx.xy1,...``` ```txc``` ```txs``` ```txsc``` ```u``` 

#### HERE API Resource Parameters
https://developer.here.com/documentation/map-image/dev_guide/topics/resource-map.html


## Tests

```HERE_MAP_APP_ID=<YOUR_APP_ID> HERE_MAP_APP_CODE=<YOUR_APP_CODE> go test -v -cover```

_coverage: **93.7%** of statements_

Check out some wonderful [examples](https://github.com/steveperjesi/go-here-map/tree/master/examples)

