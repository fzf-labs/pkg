// FindMultiBy{{.upperFieldPlural}} 根据{{.lowerFieldPlural}}查询多条数据
func (r *{{.upperTableName}}Repo) FindMultiBy{{.upperFieldPlural}}(ctx context.Context, {{.lowerFieldPlural}} []{{.dataType}}) ([]*{{.lowerDBName}}_model.{{.upperTableName}}, error) {
	dao := {{.lowerDBName}}_dao.Use(r.db).{{.upperTableName}}
	result, err := dao.WithContext(ctx).Where(dao.{{.upperField}}.In({{.lowerFieldPlural}}...)).Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}