# url-shortener-go
This is a simple URL shortener written in Go. It is one of my first projects written in Golang.

I included an example redirects file called `redirects.json.example`. Just rename that to `redirects.json` and change the redirects however you want. I also included code comments for readers to know what the code does.

To build the project simply run `go build shortener.go`

- Default port is 3000. To change the port pass a `PORT` environment variable.
- The "/" redirect is the fallback and will be used in the event of any unknown URL. For example, if I go to `/lalalejfiodkjcjduisk` and it is not in the `redirects.json`, I will be redirected to whatever the "/" is set to.

[Star this repo](https://github.com/jellz/url-shortener-go)
