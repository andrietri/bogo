package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetRepositoryTemplate(module string) (string, error) {
	repositoryConfig := NewRepositoryConfig(module)

	repositoryTemplate, err := template.New("repositoryTemplate").Parse(repositoryTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template repositoryTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err := repositoryTemplate.Execute(&templateBuf, repositoryConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type RepositoryConfig struct {
	ModulePascalCase string
	ModuleCamelCase  string
	ModuleShort      string
	ModuleSnakeCase  string
}

func NewRepositoryConfig(module string) RepositoryConfig {
	return RepositoryConfig{
		ModulePascalCase: strcase.ToCamel(module),
		ModuleCamelCase:  strcase.ToLowerCamel(module),
		ModuleShort:      getModuleShort(module),
		ModuleSnakeCase:  strcase.ToSnake(module),
	}
}

var repositoryTemplate = `package repository

import (
	"context"

	"gorm.io/gorm"
)

//go:generate mockgen -source={{.ModuleSnakeCase}}.go -destination=mock/{{.ModuleSnakeCase}}.go -package=repository

type {{.ModulePascalCase}} interface {
	Foo(ctx context.Context) error
	Bar(ctx context.Context) error
}

type {{.ModuleCamelCase}}Repository struct {
	db *gorm.DB
}

func New{{.ModulePascalCase}}Repository(db *gorm.DB) {{.ModulePascalCase}} {
	return &{{.ModuleCamelCase}}Repository{
		db: db,
	}
}

func ({{.ModuleShort}}r *{{.ModuleCamelCase}}Repository) Foo(ctx context.Context) error {
	return gorm.ErrRecordNotFound
}

func ({{.ModuleShort}}r *{{.ModuleCamelCase}}Repository) Bar(ctx context.Context) error {
	return gorm.ErrRecordNotFound
}
`
