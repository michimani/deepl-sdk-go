deepl-sdk-go
===

This is a an unofficial Go SDK for using the DeepL API.

# Usage

```bash
go get github.com/michimani/deepl-sdk-go
```

# Sample

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/michimani/deepl-sdk-go"
	"github.com/michimani/deepl-sdk-go/types"
)

func main() {
	authnKey := os.Getenv("DEEPL_AUTHN_KEY")
  isPro := true

	client := deepl.NewClient(authnKey, isPro)

	text := []string{
		"こんにちは",
		"これはサンプルテキストです。",
	}
	params := &types.TranslateTextParams{
		TargetLang: types.TargetLangEN,
		Text:       text,
	}

	res, errRes, err := c.TranslateText(context.TODO(), params)

	if err != nil {
		fmt.Println(err)
	}

	if errRes != nil {
		fmt.Println("ErrorResponse", errRes.Message)
	}

	for i := range res.Translations {
		fmt.Printf("%s -> %s\n", text[i], res.Translations[i].Text)
	}
}
```

```bash
$ DEEPL_AUTHN_KEY="your-authn-key" go run main.go

こんにちは -> hello
これはサンプルテキストです。 -> This is a sample text.
```
