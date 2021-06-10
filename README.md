# FindME

App for tracking the position of something, based on its sent GPS coordinates.

## Building this app

You will need GNU Make, SQLite3 and Go v1.16.

1. Clone or download this repository. If you download, unzip it to a folder you want.
2. Open your terminal and navigate to the folder where you've cloned or unzipped it.
3. Run `make`
4. Your binary will be available on `$FINDME_SRC/bin/findme`, where `$FINDME_SRC` is the location where you have unzipped or cloned the source code.
5. Your app will be ready at tcp port `8080`. Access it on [http://localhost:8080](http://localhost:8080/)

## Using this app

### Registering locations

To register a location, send a HTTP POST request to endpoint `/api/tracking` with body containing longitude and latitude in decimal degree format.

Example:
```shell
$ curl -X 'POST' -d 'longitude=-42.869780&latitude=-20.760753' http://localhost:8080/api/tracking
```
