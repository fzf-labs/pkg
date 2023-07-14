package gen

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/fzf-labs/fpkg/db/gen/repo"
	"golang.org/x/exp/slog"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func Generation(db *gorm.DB, dataMap map[string]func(columnType gorm.ColumnType) (dataType string), outPutPath string) {
	// 路径处理
	dbName := db.Migrator().CurrentDatabase()
	outPutPath = strings.Trim(outPutPath, "/")
	daoPath := fmt.Sprintf("%s/%s_dao", outPutPath, dbName)
	modelPath := fmt.Sprintf("%s/%s_model", outPutPath, dbName)
	repoPath := fmt.Sprintf("%s/%s_repo", outPutPath, dbName)
	// 初始化
	g := gen.NewGenerator(gen.Config{
		OutPath:      daoPath,
		ModelPkgPath: modelPath,
	})
	// 使用数据库
	g.UseDB(db)
	// 自定义字段类型映射
	g.WithDataTypeMap(dataMap)
	// json 小驼峰模型命名
	g.WithJSONTagNameStrategy(func(c string) string {
		return LowerCamelCase(c)
	})
	// 针对PG空字符串特殊处理
	g.WithOpts(gen.FieldGORMTagReg(".*?", func(tag field.GormTag) field.GormTag {
		if strings.Contains(tag.Build(), "default:''::character varying") {
			tag.Set("default", "")
		}
		return tag
	}))
	// 从数据库中生成所有表
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()

	//生成repo
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return
	}
	generationRepo := repo.NewGenerationRepo(db, daoPath, modelPath, repoPath)
	err = generationRepo.MkdirPath()
	if err != nil {
		slog.Error("repo MkdirPath err:", err)
		return
	}
	for _, tableName := range tables {
		columnNameToDataType := make(map[string]string)
		queryStructMeta := g.GenerateModel(tableName)
		for _, v := range queryStructMeta.Fields {
			columnNameToDataType[v.ColumnName] = v.Type
		}
		err = generationRepo.GenerationTable(tableName, columnNameToDataType)
		if err != nil {
			slog.Error("repo GenerationTable err:", err)
			return
		}
	}
}

// DefaultMySqlDataMap 默认mysql字段类型映射
var DefaultMySqlDataMap = map[string]func(columnType gorm.ColumnType) (dataType string){
	"int":     func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	"tinyint": func(columnType gorm.ColumnType) (dataType string) { return "int32" },
	"json":    func(columnType gorm.ColumnType) string { return "datatypes.JSON" },
	"timestamp": func(columnType gorm.ColumnType) string {
		nullable, _ := columnType.Nullable()
		if nullable {
			return "sql.NullTime"
		}
		return "time.Time"
	},
	"datetime": func(columnType gorm.ColumnType) string {
		nullable, _ := columnType.Nullable()
		if nullable {
			return "sql.NullTime"
		}
		return "time.Time"
	},
}

// DefaultPostgresDataMap 默认Postgres字段类型映射
var DefaultPostgresDataMap = map[string]func(columnType gorm.ColumnType) (dataType string){
	"json": func(columnType gorm.ColumnType) string { return "datatypes.JSON" },
	"timestamptz": func(columnType gorm.ColumnType) string {
		nullable, _ := columnType.Nullable()
		if nullable {
			return "sql.NullTime"
		}
		return "time.Time"
	},
}

// ConnectDB 数据库连接
func ConnectDB(dbType string, dsn string) *gorm.DB {
	var db *gorm.DB
	var err error
	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			panic(fmt.Errorf("connect mysql db fail: %s", err))
		}
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn))
		if err != nil {
			panic(fmt.Errorf("connect postgres db fail: %s", err))
		}
	default:
		panic(fmt.Errorf(" db type err"))
	}
	return db
}

// UpperCamelCase 下划线单词转为大写驼峰单词
func UpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	causer := cases.Title(language.English)
	s = causer.String(s)
	return strings.Replace(s, " ", "", -1)
}

// LowerCamelCase 下划线单词转为小写驼峰单词
func LowerCamelCase(s string) string {
	s = UpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
