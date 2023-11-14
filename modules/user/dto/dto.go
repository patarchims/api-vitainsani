package dto

type (
	RequestHeader struct {
		Username string `header:"x-username" bson:"x-username" binding:"required"`
		Password string `header:"x-password" bson:"x-password" binding:"required"`
	}

	// verifikasi-data
	RequestToken struct {
		Username string
		Token    string
	}
)
