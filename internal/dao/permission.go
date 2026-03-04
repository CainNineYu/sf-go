package dao

import (
	"sf-go/internal/dao/db"
	"sf-go/internal/dao/models"
)

type PermissionDAO struct {
	db *db.DB
}

func NewPermissionDAO(db *db.DB) *PermissionDAO {
	return &PermissionDAO{db: db}
}

func (p *PermissionDAO) ListPermissionCodesByRole(role string) ([]string, error) {
	var codes []string
	if err := p.db.ReadDB.Model(&models.RolePermission{}).
		Where("role = ?", role).
		Distinct().
		Pluck("code", &codes).Error; err != nil {
		return nil, err
	}
	return codes, nil
}

func (p *PermissionDAO) ListPermissionPathsByRole(role string) ([]string, error) {
	var paths []string
	err := p.db.ReadDB.
		Table("role_permissions rp").
		Select("DISTINCT p.path").
		Joins("JOIN permissions p ON p.code = rp.code").
		Where("rp.role = ?", role).
		Pluck("p.path", &paths).Error
	if err != nil {
		return nil, err
	}
	return paths, nil
}

func (p *PermissionDAO) ListPermissionsByCodes(codes []string) ([]models.Permission, error) {
	var perms []models.Permission
	if len(codes) == 0 {
		return perms, nil
	}
	if err := p.db.ReadDB.Model(&models.Permission{}).
		Where("code IN ?", codes).
		Find(&perms).Error; err != nil {
		return nil, err
	}
	return perms, nil
}
