package genutil

import (
	"strings"

	"github.com/dave/jennifer/jen" // Assuming you're using the jennifer library for Go code generation
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

type PredefinedType struct {
	Name          string
	SQLName       string
	NeedToSpecify bool
}

var dbTypeMap = map[string]PredefinedType{
	"TINYINT":          {"int8", "TINYINT", false},
	"SMALLINT":         {"int16", "SMALLINT", false},
	"MEDIUMINT":        {"int32", "MEDIUMINT", false},
	"INT":              {"int", "INT", false},
	"BIGINT":           {"int64", "BIGINT", false},
	"FLOAT":            {"float32", "FLOAT", false},
	"DOUBLE":           {"float64", "DOUBLE", false},
	"DECIMAL":          {"float64", "DECIMAL", true},
	"DEC":              {"float64", "DEC", true},
	"BIT":              {"bool", "BIT", false},
	"BOOL":             {"bool", "BOOL", false},
	"REAL":             {"float64", "REAL", false},
	"MONEY":            {"float64", "MONEY", true},
	"BINARY_FLOAT":     {"float32", "BINARY_FLOAT", false},
	"BINARY_DOUBLE":    {"float64", "BINARY_DOUBLE", false},
	"SMALLMONEY":       {"float64", "SMALLMONEY", true},
	"ENUM":             {"string", "ENUM", true},
	"CHAR":             {"string", "CHAR", false},
	"BINARY":           {"[]byte", "BINARY", false},
	"VARCHAR":          {"string", "VARCHAR", false},
	"VARBINARY":        {"[]byte", "VARBINARY", false},
	"TINYBLOB":         {"[]byte", "TINYBLOB", false},
	"TINYTEXT":         {"string", "TINYTEXT", false},
	"BLOB":             {"[]byte", "BLOB", false},
	"TEXT":             {"string", "TEXT", false},
	"MEDIUMBLOB":       {"[]byte", "MEDIUMBLOB", false},
	"MEDIUMTEXT":       {"string", "MEDIUMTEXT", false},
	"LONGBLOB":         {"[]byte", "LONGBLOB", false},
	"LONGTEXT":         {"string", "LONGTEXT", false},
	"SET":              {"string", "SET", true},
	"INET6":            {"string", "INET6", false}, // Assuming string for IP address types
	"UUID":             {"string", "UUID", false},
	"NVARCHAR":         {"string", "NVARCHAR", false},
	"NCHAR":            {"string", "NCHAR", false},
	"NTEXT":            {"string", "NTEXT", false},
	"IMAGE":            {"[]byte", "IMAGE", false},
	"VARCHAR2":         {"string", "VARCHAR2", false},
	"NVARCHAR2":        {"string", "NVARCHAR2", false},
	"DATE":             {"time.Time", "DATE", false},
	"TIME":             {"time.Duration", "TIME", false},
	"DATETIME":         {"time.Time", "DATETIME", true},
	"DATETIME2":        {"time.Time", "DATETIME2", false},
	"TIMESTAMP":        {"time.Time", "TIMESTAMP", true},
	"YEAR":             {"int", "YEAR", false},
	"SMALLDATETIME":    {"time.Time", "SMALLDATETIME", false},
	"DATETIMEOFFSET":   {"time.Time", "DATETIMEOFFSET", false},
	"XML":              {"string", "XML", true},
	"SQL_VARIANT":      {"interface{}", "SQL_VARIANT", true},
	"UNIQUEIDENTIFIER": {"string", "UNIQUEIDENTIFIER", false},
	"CURSOR":           {"interface{}", "CURSOR", true},
	"BFILE":            {"[]byte", "BFILE", false},
	"CLOB":             {"string", "CLOB", false},
	"NCLOB":            {"string", "NCLOB", false},
	"RAW":              {"[]byte", "RAW", false},
}

func getGORMTypeForName(typeName string) (string, bool) {
	predefined, ok := dbTypeMap[strings.ToUpper(typeName)]
	if !ok {
		return strutil.TryUnquote(typeName), true
	}
	if !predefined.NeedToSpecify {
		return "", false
	}
	return predefined.SQLName, true
}

func MapDBTypeToGoType(statement *jen.Statement, dbType string) {
	predefinedType, exists := dbTypeMap[strings.ToUpper(dbType)]
	if !exists {
		statement.Any() // Fallback for unknown types
		return
	}

	// Add type to statement based on the Name
	switch predefinedType.Name {
	case "int":
		statement.Int()
	case "int8":
		statement.Int8()
	case "int16":
		statement.Int16()
	case "int32":
		statement.Int32()
	case "int64":
		statement.Int64()
	case "float32":
		statement.Float32()
	case "float64":
		statement.Float64()
	case "bool":
		statement.Bool()
	case "string":
		statement.String()
	case "[]byte":
		statement.Index().Byte()
	case "time.Time":
		statement.Qual("time", "Time")
	case "net.IP":
		statement.Qual("net", "IP") // Assumed for net.IP if using net package
	default:
		statement.Any() // Fallback for any other types
	}
}
