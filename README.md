deepl-sdk-go
===

This is a an unofficial Go SDK for using the DeepL API.

# Usage

```go
func main() {
	authnKey := os.Getenv("DEEPL_AUTHN_KEY")

	client := deepl.NewClient(authnKey)

	text := []string{
		"こんにちは",
		"これはサンプルテキストです。",
	}
	params := &types.TranslateTextParams{
		AuthKey:    client.AuthenticationKey,
		TargetLang: types.EN,
		Text:       text,
	}

	res, err := client.TranslateText(context.TODO(), params)

	if err != nil {
		fmt.Println(err)
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

This sample code is included in the `sample` directory.