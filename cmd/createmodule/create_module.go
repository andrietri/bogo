package createmodule

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/andrietri/bogo/cmd/createmodule/template"
	"github.com/iancoleman/strcase"
)

// CreateModule is main func to create new Module.
func CreateModule(module string) {
	log.Println("create module start")

	log.Println("get module name")
	moduleName, err := getModuleName()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("generate controller template")
	if err := generateControllerTemplate(module, moduleName); err != nil {
		log.Fatal(err)
	}

	log.Println("generate service template")
	if err := generateServiceTemplate(module, moduleName); err != nil {
		log.Fatal(err)
	}

	log.Println("generate repository template")
	if err := generateRepositoryTemplate(module); err != nil {
		log.Fatal(err)
	}

	log.Println("generate mock repository using mockgen")
	if err := generateMockRepository(module); err != nil {
		log.Println("err generate mock repository using mockgen, did you able to run `mockgen`?")
		log.Println("try to install mockgen:\n\tgo install go.uber.org/mock/mockgen@latest")
		log.Println("create module finish without mock repository and service test template")
		return
	}

	log.Println("generate service test template")
	if err := generateServiceTestTemplate(module, moduleName); err != nil {
		log.Fatal(err)
	}

	log.Println("run go mod tidy")
	if err := runGoModTidy(); err != nil {
		log.Fatal(fmt.Errorf("err run go mod tidy: %v", err))
	}

	log.Println("create module finish")
}

func getModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", fmt.Errorf("err open go.mod: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		isLineStartsWithModule := strings.HasPrefix(line, "module ")
		if !isLineStartsWithModule {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) >= 2 {
			moduleName := parts[1]
			return moduleName, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner err: %v", err)
	}

	return "", errors.New("module name not found in go.mod file")
}

func generateControllerTemplate(module string, moduleName string) error {
	controllerTemplate, err := template.GetControllerTemplate(module, moduleName)
	if err != nil {
		return fmt.Errorf("err get controller template: %v", err)
	}

	path := filepath.Join("internal", "controller", "http", strcase.ToSnake(module)+".go")
	err = os.WriteFile(path, []byte(controllerTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write controller template: %v", err)
	}

	return nil
}

func generateServiceTemplate(module string, moduleName string) error {
	serviceTemplate, err := template.GetServiceTemplate(module, moduleName)
	if err != nil {
		return fmt.Errorf("err get service template: %v", err)
	}

	path := filepath.Join("internal", "service", strcase.ToSnake(module)+".go")
	err = os.WriteFile(path, []byte(serviceTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write service template: %v", err)
	}

	return nil
}

func generateRepositoryTemplate(module string) error {
	repositoryTemplate, err := template.GetRepositoryTemplate(module)
	if err != nil {
		return fmt.Errorf("err get repository template: %v", err)
	}

	path := filepath.Join("internal", "repository", strcase.ToSnake(module)+".go")
	err = os.WriteFile(path, []byte(repositoryTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write repository template: %v", err)
	}

	return nil
}

func runGoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout))
	}

	return nil
}

func generateMockRepository(module string) error {
	moduleFileName := strcase.ToSnake(module)

	source := fmt.Sprintf("-source=%s.go", moduleFileName)
	destination := fmt.Sprintf("-destination=%s", filepath.Join("mock", moduleFileName+".go"))

	cmd := exec.Command("mockgen", source, destination, "-package=repository")
	cmd.Dir = filepath.Join("internal", "repository")
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout))
	}

	return nil
}

func generateServiceTestTemplate(module string, moduleName string) error {
	serviceTestTemplate, err := template.GetServiceTestTemplate(module, moduleName)
	if err != nil {
		return fmt.Errorf("err get service test template: %v", err)
	}

	path := filepath.Join("internal", "service", strcase.ToSnake(module)+"_test.go")
	err = os.WriteFile(path, []byte(serviceTestTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write service test template: %v", err)
	}

	return nil
}
