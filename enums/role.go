package enums

type Role struct {
	Name string
	Slug string
}

var (
	Admin         = Role{"Admin", "role-admin"}
	User          = Role{"User", "role-user"}
	PharmacyAdmin = Role{"Pharmacy Admin", "role-admin-pharmacy"}
	Doctor        = Role{"Doctor", "role-doctor"}
)
