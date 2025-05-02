# tf - testing framework

This package is a set of [one-liners](https://en.wikipedia.org/wiki/One-liner_program) built up on [github.com/google/go-cmp](https://pkg.go.dev/github.com/google/go-cmp/cmp) package.

## Why?

- I like [one-liners](https://en.wikipedia.org/wiki/One-liner_program), they make code easier to understand.

- I just really ~~hate [testify](https://github.com/stretchr/testify)~~ like [go-cmp](https://pkg.go.dev/github.com/google/go-cmp/cmp) and it's [options](https://pkg.go.dev/github.com/google/go-cmp/cmp#Option).

<details>
<summary>Some of my thoughts about testify</summary>
<br>

- The whole world uses `assert` to fail test immediately, and `expect` to mark test as failig but continue execution, but [testify](https://github.com/stretchr/testify) uses `require`/`assert` respectively, why?

For example, [GoogleTest Primer](https://google.github.io/googletest/primer.html) says:

> The assertions come in pairs that test the same thing but have different effects on the current function. `ASSERT_*` versions generate fatal failures when they fail, and abort the current function. `EXPECT_*` versions generate nonfatal failures, which don’t abort the current function.

- Sometimes [testify](https://github.com/stretchr/testify) gives you enormous/weird outputs.

For example, try to compare structs with the same time in different timezones:

```go
package tf

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type User struct {
	CreatedAt time.Time
}

func Test_assert(t *testing.T) {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic(err)
	}

	now := time.Now()
	u1 := User{
		CreatedAt: now.In(berlin),
	}
	u2 := User{
		CreatedAt: now.In(time.UTC),
	}
	assert.Equal(t, u1, u2)
}
```

You will get this output:

```
--- FAIL: Test_assert (0.00s)
    tf_test.go:27:
                Error Trace:    /tf/tf_test.go:27
                Error:          Not equal:
                                expected: tf.User{CreatedAt:time.Date(2025, time.May, 2, 21, 38, 18, 842555000, time.Location("Europe/Berlin"))}
                                actual  : tf.User{CreatedAt:time.Date(2025, time.May, 2, 19, 38, 18, 842555000, time.UTC)}

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -4,920 +4,3 @@
                                   ext: (int64) 63881811498,
                                -  loc: (*time.Location)({
                                -   name: (string) (len=13) "Europe/Berlin",
                                -   zone: ([]time.zone) (len=9) {
                                -    (time.zone) {
                                -     name: (string) (len=3) "LMT",
                                -     offset: (int) 3208,
                                -     isDST: (bool) false
                                -    },
                                -    (time.zone) {
                                -     name: (string) (len=4) "CEST",
                                -     offset: (int) 7200,
                                -     isDST: (bool) true
                                -    },
                                -    (time.zone) {
                                -     name: (string) (len=3) "CET",
                                -     offset: (int) 3600,
                                -     isDST: (bool) false
                                -    },
                                -    (time.zone) {
                                -     name: (string) (len=4) "CEST",
                                -     offset: (int) 7200,
                                -     isDST: (bool) true
                                -    },
                                -    (time.zone) {
                                -     name: (string) (len=3) "CET",
                                -     offset: (int) 3600,
                                -     isDST: (bool) false
                                -    },
                                -    (time.zone) {
                                -     name: (string) (len=4) "CEMT",
                                -     offset: (int) 10800,
                                -     isDST: (bool) true
                                -    },
                                -    (time.zone) {
                                -     name: (string) (len=4) "CEMT",
                                -     offset: (int) 10800,
                                -     isDST: (bool) true
                                -    },
                                -    (time.zone) {
                                -     name: (string) (len=4) "CEST",
                                -     offset: (int) 7200,
                                -     isDST: (bool) true
                                -    },
                                -    (time.zone) {
                                -     name: (string) (len=3) "CET",
                                -     offset: (int) 3600,
                                -     isDST: (bool) false
                                -    }
                                -   },
                                -   tx: ([]time.zoneTrans) (len=143) {
                                -    (time.zoneTrans) {
                                -     when: (int64) -2422054408,
                                -     index: (uint8) 2,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -1693706400,
                                -     index: (uint8) 1,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -1680483600,
                                -     index: (uint8) 2,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -1663455600,
                                -     index: (uint8) 3,
                                -     isstd: (bool) true,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -1650150000,
                                -     index: (uint8) 4,
                                -     isstd: (bool) true,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -1632006000,
                                -     index: (uint8) 3,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -1618700400,
                                -     index: (uint8) 4,
                                -     isstd: (bool) true,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -938905200,
                                -     index: (uint8) 3,
                                -     isstd: (bool) true,
                                -     isutc: (bool) true
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -857257200,
                                -     index: (uint8) 4,
                                -     isstd: (bool) true,
                                -     isutc: (bool) true
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -844556400,
                                -     index: (uint8) 3,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -828226800,
                                -     index: (uint8) 4,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -812502000,
                                -     index: (uint8) 3,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -796777200,
                                -     index: (uint8) 4,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -781052400,
                                -     index: (uint8) 3,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -776563200,
                                -     index: (uint8) 5,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -765936000,
                                -     index: (uint8) 1,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -761180400,
                                -     index: (uint8) 4,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -748479600,
                                -     index: (uint8) 3,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -733273200,
                                -     index: (uint8) 4,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -717631200,
                                -     index: (uint8) 3,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -714610800,
                                -     index: (uint8) 6,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -710380800,
                                -     index: (uint8) 1,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -701910000,
                                -     index: (uint8) 4,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -684975600,
                                -     index: (uint8) 3,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -670460400,
                                -     index: (uint8) 4,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -654130800,
                                -     index: (uint8) 3,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) -639010800,
                                -     index: (uint8) 4,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 323830800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 338950800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 354675600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 370400400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 386125200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 401850000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 417574800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 433299600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 449024400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 465354000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 481078800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 496803600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 512528400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 528253200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 543978000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 559702800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 575427600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 591152400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 606877200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 622602000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 638326800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 654656400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 670381200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 686106000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 701830800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 717555600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 733280400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 749005200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 764730000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 780454800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 796179600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 811904400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 828234000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 846378000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 859683600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 877827600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 891133200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 909277200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 922582800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 941331600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 954032400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 972781200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 985482000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1004230800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1017536400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1035680400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1048986000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1067130000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1080435600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1099184400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1111885200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1130634000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1143334800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1162083600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1174784400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1193533200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1206838800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1224982800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1238288400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1256432400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1269738000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1288486800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1301187600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1319936400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1332637200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1351386000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1364691600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1382835600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1396141200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1414285200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1427590800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1445734800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1459040400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1477789200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1490490000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1509238800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1521939600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1540688400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1553994000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1572138000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1585443600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1603587600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1616893200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1635642000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1648342800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1667091600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1679792400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1698541200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1711846800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1729990800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1743296400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1761440400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1774746000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1792890000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1806195600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1824944400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1837645200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1856394000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1869094800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1887843600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1901149200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1919293200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1932598800,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1950742800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1964048400,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1982797200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 1995498000,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 2014246800,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 2026947600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 2045696400,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 2058397200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 2077146000,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 2090451600,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 2108595600,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 2121901200,
                                -     index: (uint8) 7,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    },
                                -    (time.zoneTrans) {
                                -     when: (int64) 2140045200,
                                -     index: (uint8) 8,
                                -     isstd: (bool) false,
                                -     isutc: (bool) false
                                -    }
                                -   },
                                -   extend: (string) (len=26) "CET-1CEST,M3.5.0,M10.5.0/3",
                                -   cacheStart: (int64) 1743296400,
                                -   cacheEnd: (int64) 1761440400,
                                -   cacheZone: (*time.zone)({
                                -    name: (string) (len=4) "CEST",
                                -    offset: (int) 7200,
                                -    isDST: (bool) true
                                -   })
                                -  })
                                +  loc: (*time.Location)(<nil>)
                                  }
                Test:           Test_assert
FAIL
FAIL    github.com/Snikimonkd/tf        0.180s
FAIL
```

</details>

## Usage

There are 2 packages in this lib: `assert` and `expect`. Both of them provide same functions, the only difference is that `assert` fails test immediately, while `expect` marks test as failed and continues execution.

```go
package tf

import (
	"testing"
	"time"

	"github.com/Snikimonkd/tf/assert"
)

type User struct {
	CreatedAt time.Time
}

func Test_User(t *testing.T) {
	berlin, err := time.LoadLocation("Europe/Berlin")
	assert.Nil(t, err)

	now := time.Now()
	u1 := User{
		CreatedAt: now.In(berlin),
	}
	u2 := User{
		CreatedAt: now.In(time.UTC),
	}
	assert.Equal(t, u1, u2)
}
```

You can use go-cmp [options](https://pkg.go.dev/github.com/google/go-cmp/cmp#Option):

```go
package tf

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/Snikimonkd/tf/assert"
)

func Test_Slices(t *testing.T) {
	want := []int{1, 2, 3}
	got := []int{3, 2, 1}

	assert.Equal(
		t,
		want,
		got,
		cmpopts.SortSlices(func(x, y int) bool {
			return x < y
		}),
	)
}
```

There are some predeclared go-cmp [options](https://pkg.go.dev/github.com/google/go-cmp/cmp#Option) located in package `github.com/Snikimonkd/tf/options`. For example previous example can be rewritten shorter:

```go
package tf

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/Snikimonkd/tf/assert"
	opts "github.com/Snikimonkd/tf/options"
)

func Test_Slices(t *testing.T) {
	want := []int{1, 2, 3}
	got := []int{3, 2, 1}

	assert.Equal(
		t,
		want,
		got,
		opts.SortSlicesOpt[int](),
	)
}
```
