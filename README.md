# LibSuperQi
A Golang client library to integrate with SuperQi backend easily


## ⚙️ Installation

```shell
go get -u github.com/qi-mobile/libsuperqi@v1.1.0
```

## ⚡️ Usage

```go
package main

import (
	"log"
	"time"

	"github.com/qi-mobile/libsuperqi/superqi"
)

func main() {
	client, err := superqi.InitSuperQiClient(superqi.Config{
		ClientID:               "YOUR_CLIENT_ID",
		GatewayURL:             "", // based on ENV, check docs https://superqi-docs.qi-mobile.tech
		MerchantPrivateKeyPath: "res/private_key.pem",
		IsDebug:                false,
		Timeout:                time.Second * 25,
	})

	if err != nil {
		log.Fatalln(err)
	}

	_, _ = client.ApplyToken("USER_AUTH_TOKEN")
	_, _ = client.InquiryUserInfo("USER_ACCESS_TOKEN")
	// ...
}

```

## 📖 Methods

### `InitSuperQiClient(config Config) (*Client, error)`
Loads the merchant RSA private key from `config.MerchantPrivateKeyPath` (PKCS#1, falls back to PKCS#8) and returns a configured client. The HTTP timeout comes from `config.Timeout`; set `IsDebug` to dump signed requests and responses to the standard logger.

### `ApplyToken(authCode string) (ApplyTokenResponse, error)`
Exchanges a one-time `authCode` from the user's auth flow for an `accessToken`, `refreshToken`, and `customerId`. Hold on to the access token — every other user-scoped call needs it.

### `InquiryUserInfo(accessToken string) (InquiryUserInfoResponse, error)`
Returns the user's profile: name (English and Arabic), avatar, gender, birth date, nationality, and contact info.

### `InquiryUserCardList(accessToken string) (InquiryUserCardListResponse, error)`
Lists the masked cards linked to the user.

### `InquiryUserAccountList(accessToken string) (InquiryUserAccountListResponse, error)`
Lists the bank accounts linked to the user.

### `Pay(amount int, requestId, accessToken, customerId, orderDesc, notifyUrl, productCode string) (PayResponse, error)`
Submits a payment. `amount` is whole IQD — the library multiplies by 1000 before sending, so pass `5` for 5 IQD, not `5000`. `requestId` is your idempotency key. `productCode` is one of:

```go
superqi.OnlinePurchase            // 51051000101000000011
superqi.AgreementPayment          // 51051000101000100031
superqi.OnlinePurchaseAuthCapture // 51051000101000000012
```

The response carries the gateway `paymentId` and a `redirectActionForm` you may need to render to finish 3DS / authorization.

### `InquiryPayment(paymentId, paymentRequestId string) (InquiryPaymentResponse, error)`
Looks up a payment by either ID — pass `""` for whichever you don't have. The `paymentStatus` field comes back as one of `PROCESSING`, `AUTH_SUCCESS`, `SUCCESS`, or `FAIL` (see the `PaymentStatus*` constants).

