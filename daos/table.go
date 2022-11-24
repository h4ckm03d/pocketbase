package daos

import (
	"github.com/pocketbase/dbx"
)

// HasTable checks if a table with the provided name exists (case insensitive).
func (dao *Dao) HasTable(tableName string) bool {
	var exists bool

	err := dao.DB().Select("count(*)").
		From("information_schema.tables").
		AndWhere(dbx.NewExp("LOWER([[table_name]])=LOWER({:tableName})", dbx.Params{"tableName": tableName})).
		Limit(1).
		Row(&exists)

	return err == nil && exists
}

// GetTableColumns returns all column names of a single table by its name.
func (dao *Dao) GetTableColumns(tableName string) ([]string, error) {
	columns := []string{}

	err := dao.DB().NewQuery("SELECT column_name FROM information_schema.columns WHERE table_name = '{:tableName}'").
		Bind(dbx.Params{"tableName": tableName}).
		Column(&columns)

	return columns, err
}

// DeleteTable drops the specified table.
func (dao *Dao) DeleteTable(tableName string) error {
	_, err := dao.DB().DropTable(tableName).Execute()

	return err
}
