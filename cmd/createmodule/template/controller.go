package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetControllerTemplate(module string, moduleName string) (string, error) {
	controllerConfig := NewControllerConfig(module)
	controllerConfig.ModuleName = moduleName

	controllerTemplate, err := template.New("controllerTemplate").Parse(controllerTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template controllerTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = controllerTemplate.Execute(&templateBuf, controllerConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type ControllerConfig struct {
	ModulePascalCase string
	ModuleCamelCase  string
	ModuleShort      string
	ModuleName       string
}

func NewControllerConfig(module string) ControllerConfig {
	return ControllerConfig{
		ModulePascalCase: strcase.ToCamel(module),
		ModuleCamelCase:  strcase.ToLowerCamel(module),
		ModuleShort:      getModuleShort(module),
	}
}

var controllerTemplate = `package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/internal/service"
)

type {{.ModulePascalCase}}Controller struct {
	{{.ModuleCamelCase}}Service service.{{.ModulePascalCase}}
}

func New{{.ModulePascalCase}}Controller({{.ModuleCamelCase}}Service service.{{.ModulePascalCase}}) *{{.ModulePascalCase}}Controller {
	return &{{.ModulePascalCase}}Controller{
		{{.ModuleCamelCase}}Service: {{.ModuleCamelCase}}Service,
	}
}

func ({{.ModuleShort}}c *{{.ModulePascalCase}}Controller) Bar(ctx *gin.Context) {
	if err := {{.ModuleShort}}c.{{.ModuleCamelCase}}Service.Bar(ctx); err != nil {
		log.Println(err)
		return
	}
}
`
