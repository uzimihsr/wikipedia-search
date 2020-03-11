# wikipedia-search

```bash
$ go build -o $GOPATH/bin/wikipedia .
$ wikipedia -srlimit=5 'イチロー'
---------------------------------------------------
イチロー
https://ja.wikipedia.org/?curid=1432262
---------------------------------------------------
首位打者 (日本プロ野球)
https://ja.wikipedia.org/?curid=38085
---------------------------------------------------
国道262号
https://ja.wikipedia.org/?curid=126147
---------------------------------------------------
河上イチロー
https://ja.wikipedia.org/?curid=3682529
---------------------------------------------------
新井宏昌
https://ja.wikipedia.org/?curid=688515
---------------------------------------------------

$ wikipedia -lang=en -srlimit=5 'ichiro'
---------------------------------------------------
Ichiro Suzuki
https://en.wikipedia.org/?curid=66417
---------------------------------------------------
Ichirō
https://en.wikipedia.org/?curid=1067866
---------------------------------------------------
Shimada Ichirō
https://en.wikipedia.org/?curid=186544
---------------------------------------------------
Ichiro Abe
https://en.wikipedia.org/?curid=33964428
---------------------------------------------------
Ichiro Ito
https://en.wikipedia.org/?curid=4424576
---------------------------------------------------
```
