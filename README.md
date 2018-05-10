# Kakao REST API wrapper for Golang

This is a wrapper library for [Kakao API](https://developers.kakao.com/docs) written in Golang.

## How to get

```bash
$ go get github.com/meinside/kakao-api-go
```

## How to use

```go
// sample
package main

import (
	"github.com/meinside/kakao-api-go"
)

const (
	apiKey = "0123456789abcdefghijklmnopqrstuvwxyz" // XXX - change this to yours
)

func main() {
	client := kakaoapi.NewClient(apiKey)
	//client.Verbose = true

	// TODO - do something with `client`
	// ...
}
```

See the [samples here](https://github.com/meinside/kakao-api-go/tree/master/samples).

## API coverages

- [ ] [UserManagement](https://developers.kakao.com/docs/restapi/user-management)
- [ ] [KakaoTalk](https://developers.kakao.com/docs/restapi/kakaotalk-api)
- [ ] [KakaoPay](https://developers.kakao.com/docs/restapi/kakaopay-api)
- [ ] [KakaoStory](https://developers.kakao.com/docs/restapi/kakaostory-api)
- [ ] [PushNotification](https://developers.kakao.com/docs/restapi/push-notification)
- [ ] [Search](https://developers.kakao.com/docs/restapi/search)
- [ ] [Local](https://developers.kakao.com/docs/restapi/local)
- [X] [Vision](https://developers.kakao.com/docs/restapi/vision), last update: 2018.05.

## License

MIT

