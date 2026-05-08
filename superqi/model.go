package superqi

import "time"

type PaymentStatus string

//goland:noinspection GoUnusedConst
const (
	PaymentStatusProcessing  PaymentStatus = "PROCESSING"
	PaymentStatusAuthSuccess PaymentStatus = "AUTH_SUCCESS"
	PaymentStatusSuccess     PaymentStatus = "SUCCESS"
	PaymentStatusFail        PaymentStatus = "FAIL"
)

type Result struct {
	ResultCode    string `json:"resultCode"`
	ResultStatus  string `json:"resultStatus"`
	ResultMessage string `json:"resultMessage"`
}

type ApplyTokenResponse struct {
	Result                 Result    `json:"result"`
	AccessToken            string    `json:"accessToken"`
	AccessTokenExpiryTime  time.Time `json:"accessTokenExpiryTime"`
	RefreshToken           string    `json:"refreshToken"`
	RefreshTokenExpiryTime time.Time `json:"refreshTokenExpiryTime"`
	CustomerID             string    `json:"customerId"`
}

type InquiryUserCardListResponse struct {
	Result   Result `json:"result"`
	CardList []struct {
		MaskedCardNo  string `json:"maskedCardNo"`
		AccountNumber string `json:"accountNumber"`
	} `json:"cardList"`
}

type InquiryUserAccountListResponse struct {
	Result      Result `json:"result"`
	AccountList []struct {
		AccountNumber string `json:"accountNumber"`
	} `json:"accountList"`
}

type InquiryUserInfoResponse struct {
	Result   Result `json:"result"`
	UserInfo struct {
		UserID       string `json:"userId"`
		LoginIDInfos []struct {
			LoginID     string `json:"loginId"`
			HashLoginID string `json:"hashLoginId"`
			MaskLoginID string `json:"maskLoginId"`
			LoginIDType string `json:"loginIdType"`
		} `json:"loginIdInfos"`
		UserName struct {
			FullName   string `json:"fullName"`
			FirstName  string `json:"firstName"`
			SecondName string `json:"secondName"`
			ThirdName  string `json:"thirdName"`
			LastName   string `json:"lastName"`
		} `json:"userName"`
		UserNameInArabic struct {
			FullName   string `json:"fullName"`
			FirstName  string `json:"firstName"`
			SecondName string `json:"secondName"`
			ThirdName  string `json:"thirdName"`
			LastName   string `json:"lastName"`
		} `json:"userNameInArabic"`
		Avatar       string `json:"avatar"`
		Gender       string `json:"gender"`
		BirthDate    string `json:"birthDate"`
		Nationality  string `json:"nationality"`
		ContactInfos []struct {
			ContactType string `json:"contactType"`
			ContactNo   string `json:"contactNo"`
		} `json:"contactInfos"`
	} `json:"userInfo"`
}

type PayResponse struct {
	PaymentId          string `json:"paymentId"`
	Result             Result `json:"result"`
	RedirectActionForm struct {
		Method      string `json:"method"`
		Parameters  string `json:"parameters"`
		RedirectUrl string `json:"redirectUrl"`
	} `json:"redirectActionForm"`
}

type InquiryPaymentResponse struct {
	PaymentId        string `json:"paymentId"`
	PaymentRequestId string `json:"paymentRequestId"`
	PaymentAmount    struct {
		Currency string `json:"currency"`
		Value    string `json:"value"`
	} `json:"paymentAmount"`
	PaymentTime   time.Time `json:"paymentTime"`
	PaymentStatus string    `json:"paymentStatus"`
	Result        struct {
		ResultStatus  string `json:"resultStatus"`
		ResultCode    string `json:"resultCode"`
		ResultMessage string `json:"resultMessage"`
	} `json:"result"`
	ExtendInfo   string `json:"extendInfo"`
	Transactions []struct {
		TransactionId     string    `json:"transactionId"`
		TransactionTime   time.Time `json:"transactionTime"`
		TransactionType   string    `json:"transactionType"`
		TransactionAmount struct {
			Currency string `json:"currency"`
			Value    string `json:"value"`
		} `json:"transactionAmount"`
		TransactionStatus string `json:"transactionStatus"`
	} `json:"transactions"`
}
