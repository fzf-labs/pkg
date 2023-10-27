// FindMultiByPaginator 查询分页数据(通用)
func ({{.firstTableChar}} *{{.upperTableName}}Repo) FindMultiByPaginator(ctx context.Context, paginatorReq *orm.PaginatorReq) ([]*{{.dbName}}_model.{{.upperTableName}}, *orm.PaginatorReply, error) {
	result := make([]*{{.dbName}}_model.{{.upperTableName}}, 0)
	var total int64
	whereExpressions, orderExpressions, err := paginatorReq.ConvertToGormExpression({{.dbName}}_model.{{.upperTableName}}{})
	if err != nil {
		return result, nil, err
	}
	err = {{.firstTableChar}}.db.WithContext(ctx).Model(&{{.dbName}}_model.{{.upperTableName}}{}).Select([]string{"*"}).Clauses(whereExpressions...).Count(&total).Error
	if err != nil {
		return result, nil, err
	}
	if total == 0 {
		return result, nil, nil
	}
	paginatorReply := paginatorReq.ConvertToPage(int(total))
	err = {{.firstTableChar}}.db.WithContext(ctx).Model(&{{.dbName}}_model.{{.upperTableName}}{}).Limit(paginatorReply.Limit).Offset(paginatorReply.Offset).Clauses(whereExpressions...).Clauses(orderExpressions...).Find(&result).Error
	if err != nil {
		return result, nil, err
	}
	return result, paginatorReply, err
}
