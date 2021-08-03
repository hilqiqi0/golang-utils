package mysqls

/*
有关mysql数据库连接的封装
*/
import (
	"fmt"
	"time"

	"github.com/hilqiqi0/golang-utils/tools/errs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hilqiqi0/golang-utils/common/consts"
	"github.com/hilqiqi0/golang-utils/config"
	"github.com/jmoiron/sqlx"
)

type MysqlMethod interface {
	GetDbConnFromConf(c *config.ConfigEngine, name string)
	SelectStrList(sql string, para ...interface{}) (dest []string, err error)
	SelectIntList(sql string, para ...interface{}) (dest []int, err error)
	QueryStruct(sql string, dest ...interface{}) (err error)
	QueryIdMap(sql string, para ...interface{}) (dest map[string]string, err error)
	GetDb() *sqlx.DB
	GetTableName(key string) string
}

type MysqlDbInfo struct {
	sqlDataDb *sqlx.DB
	tableName map[string]string
	maxConns  int
	dbTimeOut int
}

func getMySqlLoginStr(data *config.MysqlDbData) string {
	return data.User + ":" +
		data.Password + "@tcp(" + data.Host + ":" +
		data.Port + ")/" + data.Db_name + "?charset=utf8"
}

func (that *MysqlDbInfo) createDatabaseConns(login *config.MysqlDbData) {
	db, err := sqlx.Open("mysql", getMySqlLoginStr(login))
	errs.CheckFatalErr(err)
	db.SetConnMaxLifetime(time.Duration(login.Time_out) * time.Second)
	db.SetMaxOpenConns(login.Max_conns)
	db.SetMaxIdleConns(login.Max_conns)
	errs.CheckFatalErr(db.Ping())
	that.sqlDataDb = db
	that.tableName = login.Table_name
	that.maxConns = login.Max_conns
	that.dbTimeOut = login.Time_out
}

func (that *MysqlDbInfo) GetDbConnFromConf(c *config.ConfigEngine, name string) {
	that.createDatabaseConns(c.GetMySqlFromConf(name))
}

func (that *MysqlDbInfo) GetDb() *sqlx.DB {
	return that.sqlDataDb
}

func (that *MysqlDbInfo) GetTableName(key string) string {
	if val, ok := that.tableName[key]; ok {
		return val
	} else {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("key %s not in tablenames.", key)))
		return consts.BLANK
	}
}
