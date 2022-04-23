package db

import "github.com/prabhav-keyvalue/order-management-go/config"

func GetTableNameWithSchema(tableName string) string {
	return config.GetDbConfig().DbSchema + "." + tableName
}
