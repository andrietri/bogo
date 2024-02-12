package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetServiceTestTemplate(module string, moduleName string) (string, error) {
	serviceTestConfig := NewServiceTestConfig(module)
	serviceTestConfig.ModuleName = moduleName

	serviceTestTemplate, err := template.New("serviceTestTemplate").Parse(serviceTestTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template serviceTestTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = serviceTestTemplate.Execute(&templateBuf, serviceTestConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type ServiceTestConfig struct {
	ModulePascalCase string
	ModuleCamelCase  string
	ModuleShort      string
	ModuleName       string
}

func NewServiceTestConfig(module string) ServiceTestConfig {
	return ServiceTestConfig{
		ModulePascalCase: strcase.ToCamel(module),
		ModuleCamelCase:  strcase.ToLowerCamel(module),
		ModuleShort:      getModuleShort(module),
	}
}

var serviceTestTemplate = `package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	repository "{{.ModuleName}}/internal/repository/mock"
)

func Test_{{.ModuleCamelCase}}Service_Bar(t *testing.T) {
	type fields struct {
		{{.ModuleCamelCase}}Repository *repository.Mock{{.ModulePascalCase}}
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		wantErr bool
	}{
		{
			name: "bar",
			mock: func(f fields) {
				f.{{.ModuleCamelCase}}Repository.EXPECT().
					Foo(nil).Return(assert.AnError)
			},
			args: args{
				ctx: nil,
			},
			wantErr: true,
		},
		{
			name: "bar err baz",
			mock: func(f fields) {
				f.{{.ModuleCamelCase}}Repository.EXPECT().
					Foo(nil).Return(nil)

				f.{{.ModuleCamelCase}}Repository.EXPECT().
					Baz(nil).Return(assert.AnError)
			},
			args: args{
				ctx: nil,
			},
			wantErr: true,
		},
		{
			name: "bar success",
			mock: func(f fields) {
				f.{{.ModuleCamelCase}}Repository.EXPECT().
					Foo(nil).Return(nil)

				f.{{.ModuleCamelCase}}Repository.EXPECT().
					Baz(nil).Return(nil)
			},
			args: args{
				ctx: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				{{.ModuleCamelCase}}Repository: repository.NewMock{{.ModulePascalCase}}(ctrl),
			}
			tt.mock(f)

			{{.ModuleShort}}s := &{{.ModuleCamelCase}}Service{
				{{.ModuleCamelCase}}Repository: f.{{.ModuleCamelCase}}Repository,
			}

			err := {{.ModuleShort}}s.Bar(tt.args.ctx)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
`
