package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TemplateConf holds the schema definition for the TemplateConf entity.
type TemplateConf struct {
	ent.Schema
}

// Annotations 创建 test 表
func (TemplateConf) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "test"},
	}
}

// Indexes 建立索引
func (TemplateConf) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("remarks"),
		// 唯一约束索引
		index.Fields("user_name").Unique(),
	}
}

// Fields 创建字段 选填：Optional；Comment：注释；Unique：唯一
// Fields of the TemplateConf.
func (TemplateConf) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").Comment("账户名称").Unique(),
		field.String("pass_word").Comment("密码"),
		field.Int("status").Comment("状态（-1：删除；0：禁用；1：冻结；2：启用）"),
		field.String("remarks").Optional().Comment("备注"),
	}
}

// Edges of the TemplateConf.
func (TemplateConf) Edges() []ent.Edge {
	return nil
}
