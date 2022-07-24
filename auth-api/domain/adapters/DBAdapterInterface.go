package adapters

import (
	"context"
)

type DBAdapterInterface interface {
	Query(ctx context.Context, query string, parameters map[string]interface{}) ([]map[string]interface{}, error)

	Destruct()
}
