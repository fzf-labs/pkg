// CreateOne 创建一条数据
func (r *{{.upperTableName}}Repo) CreateOne(ctx context.Context, data *{{.lowerDBName}}_model.{{.upperTableName}}) error {
	dao := {{.lowerDBName}}_dao.Use(r.db).{{.upperTableName}}
	err := dao.WithContext(ctx).Create(data)
	if err != nil {
		return err
	}
	return nil
}