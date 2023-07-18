# Kakao REST API wrapper for Golang

This is a wrapper library for [Kakao API](https://developers.kakao.com/docs) written in Golang.

## How to get

```bash
$ go get -u github.com/meinside/kakao-api-go
```

## How to use

```go
// sample
package main

import (
	kakaoapi "github.com/meinside/kakao-api-go"
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

- [ ] [KakaoLogin](https://developers.kakao.com/docs/latest/ko/kakaologin/rest-api)
- [ ] [KakaoSync](https://developers.kakao.com/docs/latest/ko/kakaosync/common)
- [ ] [Message](https://developers.kakao.com/docs/latest/ko/message/rest-api)
- [ ] [KakaotalkSocial](https://developers.kakao.com/docs/latest/ko/kakaotalk-social/common)
- [ ] [KakaotalkChannel](https://developers.kakao.com/docs/latest/ko/kakaotalk-channel/common)
- [ ] [KakaoStory](https://developers.kakao.com/docs/latest/ko/kakaostory/rest-api)
- [ ] [PushNotification](https://developers.kakao.com/docs/latest/ko/push/rest-api)
- [ ] [TalkCalendar](https://developers.kakao.com/docs/latest/ko/talkcalendar/common)
- [ ] [Map](https://developers.kakao.com/docs/latest/ko/kakaomap/common)
- [ ] [Local](https://developers.kakao.com/docs/latest/ko/local/dev-guide)
- [ ] [KakaoNavi](https://developers.kakao.com/docs/latest/ko/kakaonavi/common)
- [ ] [DaumSearch](https://developers.kakao.com/docs/latest/ko/daum-search/dev-guide)
- [X] [KoGPT](https://developers.kakao.com/docs/latest/ko/kogpt/common)
- [X] [Karlo](https://developers.kakao.com/docs/latest/ko/karlo/common)
- [ ] [KakaoMoment](https://developers.kakao.com/docs/latest/ko/kakaomoment/common)
- [ ] [KeywordAd](https://developers.kakao.com/docs/latest/ko/keyword-ad/common)
- [ ] [KakaoPay](https://developers.kakao.com/docs/latest/ko/kakaopay/common)

## License

MIT

