package utility

import (
	"database/sql"
	"fmt"
)

func SelectData(db *sql.DB, selectList string, tableName string, condition string, joinTable string, joinKey string, orderBy string, offset int, limit int) *sql.Rows {
	if selectList == "" && tableName == "" {
		return nil
	}

	sql :=  fmt.Sprintf("SELECT %s FROM %s", selectList, tableName)
	if joinTable != "" && joinKey != "" {
		sqlJoin := fmt.Sprintf(" INNER JOIN %s ON %s", joinTable, joinKey)
		sql = sql + sqlJoin
	}
	if condition != "" {
		sqlOrder := fmt.Sprintf(" WHERE %s", condition)
		sql = sql + sqlOrder
	}
	if orderBy != "" {
		sqlOrder := fmt.Sprintf(" ORDER BY %s", orderBy)
		sql = sql + sqlOrder
	}
	if offset > 0 {
		sqlOffset := fmt.Sprintf(" OFFSET %d", offset)
		sql = sql + sqlOffset
	}
	if limit > 0 {
		sqlLimit := fmt.Sprintf(" LIMIT %d", limit)
		sql = sql + sqlLimit
	}
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	return rows
}

func SelectDataManual(db *sql.DB, sql string) *sql.Rows {
	if sql == "" {
		return nil
	}
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	return rows
}