// FindMultiByPaginator 查询分页数据(通用)
FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*{{.lowerDBName}}_model.{{.upperTableName}}, *orm.PaginatorReply, error)