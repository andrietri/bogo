package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetServiceTemplate(module string, moduleName string) (string, error) {
	serviceConfig := NewServiceConfig(module)
	serviceConfig.ModuleName = moduleName

	serviceTemplate, err := template.New("serviceTemplate").Parse(serviceTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template serviceTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = serviceTemplate.Execute(&templateBuf, serviceConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}
	return templateBuf.String(), nil
}

type ServiceConfig struct {
	ModulePascalCase string
	ModuleCamelCase  string
	ModuleShort      string
	ModuleName       string
}

func NewServiceConfig(module string) ServiceConfig {
	return ServiceConfig{
		ModulePascalCase: strcase.ToCamel(module),
		ModuleCamelCase:  strcase.ToLowerCamel(module),
		ModuleShort:      getModuleShort(module),
	}
}

var serviceTemplate = `package service

import (
	"context"
	"fmt"

	"{{.ModuleName}}/internal/repository"
)

type {{.ModulePascalCase}} interface {
	Bar(ctx context.Context) error
}

type {{.ModuleCamelCase}}Service struct {
	{{.ModuleCamelCase}}Repository repository.{{.ModulePascalCase}}
}

func New{{.ModulePascalCase}}Service({{.ModuleCamelCase}}Repository repository.{{.ModulePascalCase}}) {{.ModulePascalCase}} {
	return &{{.ModuleCamelCase}}Service{
		{{.ModuleCamelCase}}Repository: {{.ModuleCamelCase}}Repository,
	}
}

func ({{.ModuleShort}}s *{{.ModuleCamelCase}}Service) Bar(ctx context.Context) error {
	if err := {{.ModuleShort}}s.{{.ModuleCamelCase}}Repository.Foo(ctx); err != nil {
		return fmt.Errorf("err : %v", err)
	}

	if err := {{.ModuleShort}}s.{{.ModuleCamelCase}}Repository.Bar(ctx); err != nil {
		return fmt.Errorf("err : %v", err)
	}

	return nil
}
`
