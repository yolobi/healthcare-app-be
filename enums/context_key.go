package enums

type KeyContext struct {
	Key string
}

var (
	CertificateKey KeyContext = KeyContext{
		Key: "X-Certificate",
	}

	UserIdKey KeyContext = KeyContext{
		Key: "user-id",
	}

	RequestIdKey KeyContext = KeyContext{
		Key: "request-id",
	}

	CategoryIconKey KeyContext = KeyContext{
		Key: "category-icon",
	}

	ProductImageKey KeyContext = KeyContext{
		Key: "product-image",
	}

	PaymentFileKey KeyContext = KeyContext{
		Key: "payment-file",
	}

	UserPhotoKey KeyContext = KeyContext{
		Key: "user-photo",
	}
)
