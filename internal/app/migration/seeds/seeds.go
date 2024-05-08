package seeds

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
)

var Seeds = []*gormigrate.Migration{
	&S20240428123800_users,
	&S20240428123801_companies,
	&S20240428123803_credit_cards,
	&S20240428123804_roles,
	&S20240428123805_user_roles,
}
