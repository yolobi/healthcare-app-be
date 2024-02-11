package databases

import (
	"context"
	"healthcare-capt-america/services"

	auth "github.com/harranali/authority"
)

var ctx = context.Background()

func (r *Repositories) CategorySeeding() {
	for _, category := range categories {
		r.CategoryRepository.Save(ctx, &category)
	}
}

func (r *Repositories) RoleSeeding() {
	for _, role := range roles {
		services.Authority.CreateRole(auth.Role{
			Name: role.Name,
			Slug: role.Slug,
		})
	}
}

func (r *Repositories) PermissionSeeding() {
	for _, permission := range permissions {
		services.Authority.CreatePermission(auth.Permission{
			Name: permission.Name,
			Slug: permission.Slug,
		})
	}
}

func (r *Repositories) RolePermissionSeeding() {
	//need to be filled
}

func (r *Repositories) ShipmentSeeding() {
	for _, shipment := range shipments {
		r.ShipmentRepository.Save(ctx, &shipment)
	}
}

func (r *Repositories) RoomStatusSeeding() {
	r.db.Create(room_status)
}
