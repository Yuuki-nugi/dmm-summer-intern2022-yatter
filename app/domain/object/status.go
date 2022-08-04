package object

type (
	StatusID = int64

	// Account account
	Status struct {
		// The internal ID of the account
		ID StatusID `json:"id"`

		// The username of the account
		AccountID int64 `json:"-" db:"account_id"`

		Account Account `json:"account" db:"account"`

		Content *string `json:"content"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
