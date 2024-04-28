package migrations

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
)

var Migrations = []*gormigrate.Migration{
	&M20240428123800_users,
	&M20240428123801_companies,
	&M20240428123803_credit_cards,
	&M20240428123804_roles,
	&M20240428123805_user_roles,
}
