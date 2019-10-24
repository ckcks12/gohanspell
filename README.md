# gohanspell

![](https://github.com/ckcks12/gohanspell/workflows/Go/badge.svg)

`gohanspell` 은 `hanspell` 을 golang 으로 포팅한 라이브러리입니다.


### 사용법

```go
txt, err := PostPusanUniv(`외않되?`)
if err != nil {
    panic(err)
}
log.Println(txt) // 왜 안 돼
```


### 주의

특정 단어가 여러 개의 단어로 교정될 수 있으면 그중 랜덤으로 하나를 고릅니다.
