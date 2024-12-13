package repositories

import (
	"github.com/Xurliman/auth-service/internal/server/app/interfaces"
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/Xurliman/auth-service/internal/server/app/responses"
	"github.com/Xurliman/auth-service/pkg/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) EmailExists(email string) bool {
	var user models.User
	if err := r.db.Unscoped().Select("id").First(&user, "email = ?", email).Error; err != nil {
		return false
	}
	return true
}

func (r *UserRepository) FindByEmail(email string) (user models.User, err error) {
	if err = r.db.Model(&user).First(&user, "email = ?", email).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindById(id string) (user models.User, err error) {
	if err = r.db.Model(&user).First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) GetAll(pagination pagination.Pagination) (*pagination.Pagination, error) {
	var users []*models.User
	clauses := r.filterClauses(&pagination)
	filter := filterPaginate(users, &pagination, clauses)
	if err := r.db.Scopes(filter).Find(&users).Error; err != nil {
		return nil, err
	}

	pagination.Rows = responses.UserListTransformer(users)
	return &pagination, nil
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) UpdateById(id string, storeData models.User) (models.User, error) {
	var user models.User
	if err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&storeData).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) Delete(id string) error {
	var user models.User
	if err := r.db.Delete(&user, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) filterClauses(pagination *pagination.Pagination) (clauses []clause.Expression) {
	if pagination.Search != "" {
		vars := setKeywordVarsByTotalExpr(pagination.Search, 2)
		query := lowerLikeQuery("username") + " OR " + lowerLikeQuery("email")
		clauses = append(clauses, clause.Expr{SQL: query, Vars: vars})
	}

	return clauses
}
