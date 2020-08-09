package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/willabides/octo-go/generator/internal/model"
	"github.com/willabides/octo-go/generator/internal/model/openapi"
)

func main() {
	var schemaPath string
	var outputPath string
	var pkgPath string
	var pkgName string
	flag.StringVar(&schemaPath, "schema", "", "path to openapi schema")
	flag.StringVar(&outputPath, "out", "", "directory to write all these files")
	flag.StringVar(&pkgPath, "pkgpath", "", "path for output package")
	flag.StringVar(&pkgName, "pkg", "", "name for output package")
	flag.Parse()
	err := run(schemaPath, outputPath, pkgPath, pkgName)
	if err != nil {
		log.Fatal(err)
	}
}

func mkDirs(outputPath string) error {
	for _, dirName := range []string{"components/examples", "internal"} {
		target := filepath.Join(outputPath, filepath.FromSlash(dirName))
		err := os.MkdirAll(target, 0o750)
		if err != nil {
			return err
		}
	}
	return nil
}

func run(schemaPath, outputPath, pkgPath, pkgName string) error {
	err := mkDirs(outputPath)
	if err != nil {
		return err
	}
	err = componentExampleFiles(schemaPath, filepath.Join(outputPath, "components", "examples"))
	if err != nil {
		return err
	}
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		return err
	}
	mdl, err := openapi.Openapi2Model(schemaFile)
	if err != nil {
		return err
	}
	endpoints := mdl.Endpoints

	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].ID < endpoints[j].ID
	})

	concernFiles := map[string]*jen.File{}

	for _, endpoint := range endpoints {
		err = addEndpointToConcernFiles(endpoint, concernFiles, pkgPath, pkgName)
		if err != nil {
			return err
		}
	}

	err = renderConcernFiles(concernFiles, outputPath)
	if err != nil {
		return err
	}

	pq := pkgQual(pkgPath)
	componentSchemas := mdl.ComponentSchemas
	if len(componentSchemas) > 0 {
		filename := "schemas_gen.go"
		var f *os.File
		f, err = os.Create(filepath.Join(outputPath, "components", filename))
		if err != nil {
			return err
		}
		cf := jen.NewFilePath(path.Join(pkgPath, "components"))
		cf.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
		addComponentSchemas(cf, componentSchemas, pq)
		err = cf.Render(f)
		if err != nil {
			return err
		}
	}

	internalPkgName := "internal"
	internalPkgPath := path.Join(pkgPath, "internal")
	internalOutputPath := filepath.Join(outputPath, "internal")

	err = writeIDMapFile(endpoints, internalPkgPath, internalPkgName, internalOutputPath)
	if err != nil {
		return err
	}

	err = writeAttrsFile(endpoints, internalPkgPath, internalPkgName, internalOutputPath)
	if err != nil {
		return err
	}

	err = generateUnmarshalTests(outputPath, pq, endpoints)
	if err != nil {
		return err
	}

	return nil
}

func writeAttrsFile(endpoints []*model.Endpoint, pkgPath, pkgName, outputPath string) error {
	file := jen.NewFilePathName(pkgPath, pkgName)
	file.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
	file.Const().Parens(jen.Do(func(statement *jen.Statement) {
		for i := endpointAttribute(0); i < attrInvalid; i++ {
			statement.Id(toExportedName(attrNames[i]))
			if i == 0 {
				statement.Id("EndpointAttribute").Op("=").Iota()
			}
			statement.Line()
		}
	}))
	file.Func().Id("init()").Block(
		jen.Id("operationAttributes = operationAttributesGen"),
	)
	file.Id("var operationAttributesGen = map[string][]EndpointAttribute").Block(opIdAttributesDict(endpoints))
	out, err := os.Create(filepath.Join(outputPath, "attributes_gen.go"))
	if err != nil {
		return err
	}
	return file.Render(out)
}

func writeIDMapFile(endpoints []*model.Endpoint, pkgPath, pkgName, outputPath string) error {
	file := jen.NewFilePathName(pkgPath, pkgName)
	file.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
	file.Func().Id("init()").Block(
		jen.Id("reqOperationIDs = reqOperationIDsGen"),
	)
	file.Id("var reqOperationIDsGen = map[string]string").Block(
		jen.DictFunc(func(dict jen.Dict) {
			for _, endpoint := range endpoints {
				dict[jen.Lit(reqStructName(endpoint))] = jen.Lit(endpoint.ID)
			}
		}),
	)
	out, err := os.Create(filepath.Join(outputPath, "idmaps_gen.go"))
	if err != nil {
		return err
	}
	return file.Render(out)
}

func renderConcernFiles(concernFiles map[string]*jen.File, outputPath string) error {
	for concern, concernFile := range concernFiles {
		// zz_licenses_gen.go causes issues with license file detection on GitHub
		if concern == "licenses" {
			concern = "lic"
		}
		name := fmt.Sprintf("zz_%s_gen.go", strings.ReplaceAll(concern, "-", "_"))
		var f *os.File
		f, err := os.Create(filepath.Join(outputPath, name))
		if err != nil {
			return err
		}
		err = concernFile.Render(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func addEndpointToConcernFiles(endpoint *model.Endpoint, concernFiles map[string]*jen.File, pkgPath, pkgName string) error {
	pq := pkgQual(pkgPath)
	err := applyEndpointOverrides(endpoint)
	if err != nil {
		return err
	}
	if endpoint.Legacy {
		return nil
	}
	if concernFiles[endpoint.Concern] == nil {
		cf := jen.NewFilePathName(pkgPath, pkgName)
		cf.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
		concernFiles[endpoint.Concern] = cf
	}
	file := concernFiles[endpoint.Concern]
	addRequestFunc(file, pq, endpoint)
	addClientMethod(file, pq, endpoint)
	addRequestStruct(file, pq, endpoint)
	addRequestBody(file, pq, endpoint)
	addResponseBody(file, pq, endpoint)
	addResponse(file, pq, endpoint)
	return nil
}
