package schema

import (
	"Myorm/dialect"
	"go/ast"
	"reflect"
)

// Field represents a column of database
//Field 包含 3 个成员变量，字段名 Name、类型 Type、和约束条件 Tag
type Field struct {
	Name string
	Type string
	Tag  string
}

// Schema represents a table of database
//Schema 主要包含被映射的对象 Model、表名 Name 和字段 Fields。
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string          //FieldNames 包含所有的字段名(列名)
	fieldMap   map[string]*Field //fieldMap 记录字段名和 Field 的映射关系，方便之后直接使用，无需遍历 Fields。
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)

		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type)))}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
