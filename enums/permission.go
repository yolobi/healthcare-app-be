package enums

type Permission struct {
	Name        string
	Slug        string
	Description string
}

var (
	// it's just a sample permission
	// this admin dashboard better handled with role, not permission
	AdminDashboard = Permission{"Admin Dashboard", "admin-dashboard", "permission to access admin dashboard"}
)
