# golang_encode_and_decode_strings

Design an algorithm to encode a list of strings to a string. The encoded string is then sent over the network and is decoded back to the original list of strings.

Please implement `encode` and `decode`

*Contact me on wechat to get **Amazon、Google** requent Interview questions . (wechat id : **jiuzhang0607**)*

## Examples

**Example1**

```
Input: ["lint","code","love","you"]
Output: ["lint","code","love","you"]
Explanation:
One possible encode method is: "lint:;code:;love:;you"

```

**Example2**

```
Input: ["we", "say", ":", "yes"]
Output: ["we", "say", ":", "yes"]
Explanation:
One possible encode method is: "we:;say:;:::;yes"

```

## 解析

要求實作一個演算法用來把一個字串陣列編碼成一個字串

情境是先把字串陣列轉編碼成一個字字串 str

再把字串 str 解碼為原本的字傳陣列

所以需要實作兩個 method: encode 跟 decode

要把字串陣列編碼成一個字串，並且能夠解碼回來元字串

最容易想到的方法是透過一個分隔符合來把每個字傳做連結 當作編碼方式

而解碼方式就是把編碼後的字串依據這個分隔符號做拆分

然而如果用單一個字元做分隔符怕會出現分隔符剛好是某個字串的字

![](https://i.imgur.com/DwCrVTR.png)

這樣就會導致錯誤切分字串

要避免切割錯誤的字元數目其實只要代切割字元之前加入字串長度的訊息

如下圖：
![](https://i.imgur.com/CcSMYy6.png)

而 Decode 時同樣先從 第一個 ‘#’ 前的數字讀取出字串長度，然後再從第一個 ‘#’ 的下一個字元開始蒐集字串

如下圖:
![](https://i.imgur.com/RNo8yf3.png)

具體作法如下：

![](https://i.imgur.com/rvS3Pmg.png)

這樣每次encode 跟 decode 都需要走訪整個串列 所以時間複雜度是 O(n)

## 程式碼
```go
package sol

import (
	"fmt"
	"strconv"
)

type ListHandle struct {
	sep string
}

func Constructor(sep string) ListHandle {
	return ListHandle{
		sep: sep,
	}
}
func (h *ListHandle) encode(list []string) string {
	encodeString := ""
	for _, val := range list {
		encodeString += fmt.Sprintf("%d%s%s", len(val), h.sep, val)
	}
	return encodeString
}

func (h *ListHandle) decode(str string) []string {
	decodeResult := []string{}
	pos := 0
	strLen := len(str)
	tempLen := ""
	for pos < strLen {
		if string(str[pos]) != h.sep {
			tempLen += string(str[pos])
		} else { // first meet #
			num, _ := strconv.Atoi(tempLen)
			decodeResult = append(decodeResult, str[pos+1:pos+num+1])
			tempLen = ""
			pos += num + 1
			continue
		}
		pos++
	}
	return decodeResult
}

```
## 困難點

1. 需要想出一個好的方法來對字串陣列來做編碼與解碼
2. 需要注意到單一字元可能會成為字串的一部份

## Solve Point
- [x]  實作 encode : 把輸入字串陣列以 字串長度+分隔符號+字串來做編碼
- [x]  實作 decode: 把輸入字串先找出分隔符號以其之前的字串作為字串長度做分割