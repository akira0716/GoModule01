package greetings

import (
	"errors"
	"fmt"
)

// Hello は、指定された人物に対する挨拶を返します。
func Hello(name string) (string, error) {
	// 名前が空の場合は、メッセージ付きのエラーを返します。
	if name == "" {
		return "", errors.New("名前が未入力ですよ。")
	}

	// 名前を受け取った場合は、挨拶メッセージに名前を埋め込んだ値を返します。
	message := fmt.Sprintf("やぁ, %v. ようこそ！", name)
	return message, nil
}
