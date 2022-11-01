[![Run on Repl.it](https://replit.com/badge/github/ej-agas/go-exercise-4.12)](https://replit.com/github/ej-agas/go-exercise-4.12)

# Go programming language book excercise 4.12 solution

Solution to exercise 4.12 in [The Go Programming Language](http://www.gopl.io) book.

The popular web comic "xkcd" has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd that, using this index, prints the URL, date, and title of each comic whose *title* and *transcript* matches a *list* of *search terms* provided on the command line.

![Can't sleep](https://imgs.xkcd.com/comics/cant_sleep.png "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.")

`load.go` reads the xkcd.com JSON endpoint and writes it to a file named after what you have passed as an argument until it encounters 2 HTTP 404 responses (2 because there is no comic #404, xkcd.com returns HTTP 404 on comic #404 as an easter egg).

```shell
$ go run ./loader xkcd.json
read 2691 comics
```

`read.go` takes some words to find from the command line (after the JSON file) and finds comics whose title or transcript matches all the words

```shell
$ go run ./reader xkcd.json "can't sleep"
read 2691 comics
https://xkcd.com/571/ 4/20/2009  "Can't Sleep"
found 1 comic(s)