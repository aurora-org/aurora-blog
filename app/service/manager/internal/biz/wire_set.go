package biz

import (
	"github.com/google/wire"
)

// ProvideSet for biz package ...
var ProvideSet = wire.NewSet(
	NewAccountBusiness,
	NewArticleBusiness,
	NewCategoryBusiness,
	NewSiteBusiness,
	NewTagBusiness,
	NewThemeBusiness,
	NewUserBusiness,
)

func NewAccountBusiness(repo AccountRepo) *AccountBusiness {
	return &AccountBusiness{repo: repo}
}

func NewArticleBusiness(repo ArticleRepo) *ArticleBusiness {
	return &ArticleBusiness{repo: repo}
}

func NewCategoryBusiness(repo CategoryRepo) *CategoryBusiness {
	return &CategoryBusiness{repo: repo}
}

func NewSiteBusiness(repo SiteRepo) *SiteBusiness {
	return &SiteBusiness{repo: repo}
}

func NewTagBusiness(repo TagRepo) *TagBusiness {
	return &TagBusiness{repo: repo}
}

func NewThemeBusiness(repo ThemeRepo) *ThemeBusiness {
	return &ThemeBusiness{repo: repo}
}

func NewUserBusiness(repo UserRepo) *UserBusiness {
	return &UserBusiness{repo: repo}
}
