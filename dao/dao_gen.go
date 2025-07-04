package dao

import (
	"database/sql"

	"gitee.com/go-framework_llllllh/go-backnormal-gen/util"
	_ "github.com/go-sql-driver/mysql" // 或其他数据库驱动
)

func GenDao_Mysql(dsn string, tables []string, DaoDir string, boPath string) error {
	// 1. 连接数据库
	//dsn: "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// 2. 查询表名
	// 查询所有表名
	// tables, err := getTables(db)
	// if err != nil {
	// 	return err
	// }

	// 指定生成的表名
	// tables := []string{
	// 	// "user",          // 只生成 user 表
	// }

	// 3. 为每个表生成模型
	for _, table := range tables {
		model := util.Model{
			ModelName: util.ToCamelCase(table), // 表名转结构体名（如 "user_info" -> "UserInfo"）
			Tablename: table,                   // 表名
			BoPath:    boPath,
		}

		// 查询表字段信息
		model.Fields, err = util.GetTableFields(db, table)
		if err != nil {
			return err
		}

		// 4. 生成 Dao 文件
		tplPath := "./dao.tpl"
		if err := util.Generate(model, tplPath, DaoDir); err != nil {
			return err
		}
	}

	return nil
}
