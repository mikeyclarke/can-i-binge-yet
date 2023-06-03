package show

import (
    "context"
)

type ShowLoaderInterface interface {
    Load(ctx context.Context, id int) *ShowPageResult
}
