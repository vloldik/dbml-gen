package opt

import gorm "gorm.io/gorm"

type PreloadOption string
type WhereOption []any

func Where(cond any, conds ...any) WhereOption {
	return append([]any{cond}, conds...)
}

func ApplyOptions(db *gorm.DB, options ...any) *gorm.DB {
	for _, option := range options {
		db = applyOption(db, option)
	}
	return db
}

func applyOption(db *gorm.DB, option any) *gorm.DB {
	switch o := option.(type) {
	case PreloadOption:
		return db.Preload(string(o))
	case WhereOption:
		return db.Where(o[0], o[1:]...)
	default:
		return db
	}
}