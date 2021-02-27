# TUM-Live

**Work in progress** of TUMs lecture streaming service

## Disclaimer

~ 95% of modern browsers are supported. IE is not and probably won't be (officially).
Reason being is that I want to avoid using jQuery because it's way too bloated. 
This is about what you can expect:

Browser | Chrome | Edge | Firefox | IE | Opera | Safari | Android Webview | Chrome Android | Firefox Android | Safari iOS | Samsung Internet
--- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | ---
version | 42+ | 14+ | 39+ | no | 29+ | 10.1+ | 42+ | 42+ | 39+ | 29+ | 10.3+ | 4.0+


## Credit & Licenses

- [ffmpeg](https://ffmpeg.org/) Licenced under [lgpl-2.1](http://www.gnu.org/licenses/old-licenses/lgpl-2.1.html)
- [NGINX](https://www.nginx.com/) Licenced under [FreeBSD](http://nginx.org/LICENSE)
- [go](https://golang.org/) Licence [here](https://golang.org/LICENSE)
- [mariadb](https://mariadb.com/) Licenced under [GPL](https://mariadb.com/kb/en/mariadb-license/)
- todo

## Setup

todo

## Development

To avoid rebuilding the entire system everytime you test something you can start by 
running the stack using `docker-compose build && docker-compose up`

Now log in to the database on localhost:3306 and create a database called rbglive.
If you now stop the stack and run `docker-compose up` again your system should be up
and you can start a development server with `npm i --no-dev && go run app/server/main.go`

Now head over to localhost:8080

A more solid dev environment will follow soon.
