package adapters

import (
	"auth-api/app/config"
	"auth-api/domain/adapters"
	externalErrs "auth-api/external/errors"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
)

type MySQLAdapter struct {
	cfg      config.DBConfig
	pool     *sql.DB
	pqPrefix string
}

func (m MySQLAdapter) Query(ctx context.Context, query string, parameters map[string]interface{}) ([]map[string]interface{}, error) {
	convertedQuery, placeholders := m.convertQuery(query)

	reorderedParameters, err := m.reorderParameters(parameters, placeholders)
	if err != nil {
		return nil, err
	}

	statement, err := m.prepareStatement(ctx, convertedQuery)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(convertedQuery[:1]) == "s" {

		rows, err := statement.Query(reorderedParameters...)
		if err != nil {
			return nil, err
		}

		return m.prepareDataSet(rows)
	}

	result, err := statement.Exec(reorderedParameters...)
	if err != nil {
		return nil, err
	}

	return m.prepareResultSet(result)
}

func (m MySQLAdapter) Destruct() {
	panic("implement me")
}

func NewMySQLAdapter(cfg config.DBConfig) (adapters.DBAdapterInterface, error) {

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	a := &MySQLAdapter{
		cfg:      cfg,
		pool:     db,
		pqPrefix: "?",
	}

	return a, nil
}

func (a *MySQLAdapter) reorderParameters(params map[string]interface{}, namedParams []string) ([]interface{}, error) {

	var reorderedParams []interface{}

	for _, param := range namedParams {

		paramValue, isParamExist := params[param]

		if !isParamExist {
			return nil, externalErrs.NewAdapterError(fmt.Sprintf("parameter '%s' is missing", param), 100, "")
		}

		reorderedParams = append(reorderedParams, paramValue)
	}

	return reorderedParams, nil
}

func (a *MySQLAdapter) prepareDataSet(rows *sql.Rows) ([]map[string]interface{}, error) {

	defer rows.Close()

	var data []map[string]interface{}
	cols, _ := rows.Columns()

	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))

	for i := range columns {
		columnPointers[i] = &columns[i]
	}

	for rows.Next() {
		err := rows.Scan(columnPointers...)
		if err != nil {
			return nil, err
		}

		row := make(map[string]interface{})

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			row[colName] = *val
		}

		data = append(data, row)
	}

	return data, nil
}

func (a *MySQLAdapter) prepareResultSet(result sql.Result) ([]map[string]interface{}, error) {

	var data []map[string]interface{}

	row := make(map[string]interface{})

	row["affected_rows"], _ = result.RowsAffected()
	row["last_insert_id"], _ = result.LastInsertId()

	return append(data, row), nil
}

func (a *MySQLAdapter) prepareStatement(ctx context.Context, query string) (*sql.Stmt, error) {

	tx := ctx.Value("Tx")
	if tx != nil {
		return tx.(*sql.Tx).Prepare(query)
	}

	return a.pool.Prepare(query)
}

func (a *MySQLAdapter) convertQuery(query string) (string, []string) {

	query = strings.TrimSpace(query)
	exp := regexp.MustCompile(`\` + a.pqPrefix + `\w+`)

	namedParams := exp.FindAllString(query, -1)

	for i := 0; i < len(namedParams); i++ {
		namedParams[i] = strings.TrimPrefix(namedParams[i], a.pqPrefix)
	}

	query = exp.ReplaceAllString(query, "?")

	return query, namedParams
}
