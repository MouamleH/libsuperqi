# LibSuperQi
A Golang client library to integrate with SuperQi backend easily


## ⚙️ Installation

```shell
go get -u github.com/mouamleh/libsuperqi
```

## ⚡️ Usage

```go
package main

import "github.com/mouamleh/libsuperqi/superqi"

func main() {
	err := superqi.InitSuperQiClient(superqi.Config{
		ClientID:               "YOUR_CLIENT_ID",
		GatewayURL:             "", // based on ENV, check docs https://superqi-dev-docs.pages.dev/
		MerchantPrivateKeyPath: "res/private_key.pem",
		IsDebug:                false,
		Timeout:                time.Second * 25,
	})

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	_, _ := client.ApplyToken("USER_AUTH_TOKEN")
	_, _ := client.InquiryUserInfo("USER_ACCESS_TOKEN")
	...
}
```
