package gorm

import "github.com/jinzhu/gorm/clause"

// Model specify the model you would like to run db operations
//    // update all users's name to `hello`
//    db.Model(&User{}).Update("name", "hello")
//    // if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
//    db.Model(&user).Update("name", "hello")
func (db *DB) Model(value interface{}) (tx *DB) {
	tx = db.getInstance()
	tx.Statement.Model = value
	return
}

// Clauses Add clauses
func (db *DB) Clauses(conds ...clause.Expression) (tx *DB) {
	tx = db.getInstance()
	var whereConds []interface{}

	for _, cond := range conds {
		if c, ok := cond.(clause.Interface); ok {
			tx.Statement.AddClause(c)
		} else {
			whereConds = append(whereConds, cond)
		}
	}

	if len(whereConds) > 0 {
		tx.Statement.AddClause(clause.Where{
			AndConditions: tx.Statement.BuildCondtion(whereConds[0], whereConds[1:]...),
		})
	}
	return
}

// Table specify the table you would like to run db operations
func (db *DB) Table(name string) (tx *DB) {
	tx = db.getInstance()
	tx.Statement.Table = name
	return
}

// Select specify fields that you want when querying, creating, updating
func (db *DB) Select(query interface{}, args ...interface{}) (tx *DB) {
	tx = db.getInstance()
	return
}

// Omit specify fields that you want to ignore when creating, updating and querying
func (db *DB) Omit(columns ...string) (tx *DB) {
	tx = db.getInstance()
	return
}

func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB) {
	tx = db.getInstance()
	tx.Statement.AddClause(clause.Where{AndConditions: tx.Statement.BuildCondtion(query, args...)})
	return
}

// Not add NOT condition
func (db *DB) Not(query interface{}, args ...interface{}) (tx *DB) {
	tx = db.getInstance()
	tx.Statement.AddClause(clause.Where{
		AndConditions: []clause.Expression{clause.NotConditions(tx.Statement.BuildCondtion(query, args...))},
	})
	return
}

// Or add OR conditions
func (db *DB) Or(query interface{}, args ...interface{}) (tx *DB) {
	tx = db.getInstance()
	tx.Statement.AddClause(clause.Where{
		ORConditions: []clause.ORConditions{tx.Statement.BuildCondtion(query, args...)},
	})
	return
}

// Joins specify Joins conditions
//     db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
func (db *DB) Joins(query string, args ...interface{}) (tx *DB) {
	tx = db.getInstance()
	return
}

// Group specify the group method on the find
func (db *DB) Group(column string) (tx *DB) {
	tx = db.getInstance()
	return
}

// Having specify HAVING conditions for GROUP BY
func (db *DB) Having(query interface{}, args ...interface{}) (tx *DB) {
	tx = db.getInstance()
	return
}

// Order specify order when retrieve records from database
//     db.Order("name DESC")
//     db.Order(gorm.Expr("name = ? DESC", "first")) // sql expression
func (db *DB) Order(value interface{}) (tx *DB) {
	tx = db.getInstance()
	return
}

// Limit specify the number of records to be retrieved
func (db *DB) Limit(limit int64) (tx *DB) {
	tx = db.getInstance()
	return
}

// Offset specify the number of records to skip before starting to return the records
func (db *DB) Offset(offset int64) (tx *DB) {
	tx = db.getInstance()
	return
}

// Scopes pass current database connection to arguments `func(*DB) *DB`, which could be used to add conditions dynamically
//     func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
//         return db.Where("amount > ?", 1000)
//     }
//
//     func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
//         return func (db *gorm.DB) *gorm.DB {
//             return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
//         }
//     }
//
//     db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
func (db *DB) Scopes(funcs ...func(*DB) *DB) *DB {
	for _, f := range funcs {
		db = f(db)
	}
	return db
}

// Preload preload associations with given conditions
//    db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func (db *DB) Preload(column string, conditions ...interface{}) (tx *DB) {
	tx = db.getInstance()
	return
}

func (db *DB) Assign(attrs ...interface{}) (tx *DB) {
	tx = db.getInstance()
	return
}

func (db *DB) Attrs(attrs ...interface{}) (tx *DB) {
	tx = db.getInstance()
	return
}

func (db *DB) Unscoped() (tx *DB) {
	tx = db.getInstance()
	return
}

func (db *DB) Raw(sql string, values ...interface{}) (tx *DB) {
	tx = db.getInstance()
	return
}
