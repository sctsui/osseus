package gencalls

import (
	"archive/tar"
	"bytes"
	"encoding/base64"
	"html/template"
	"log"

	"github.com/ligato/osseus/plugins/generator/model"
)

const tpl = `
package main
	
import "fmt"
func main() {
    fmt.Println({{.Title}})
}`

type fileEntry struct{
	Name string
	Body string
}

// GenAddProj creates a new generated template under the /template prefix
func (d *ProjectHandler) GenAddProj(key string, val *model.Project) error {
	// Init buf writer
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	var genCode bytes.Buffer
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, er := template.New("webpage").Parse(tpl)
	check(er)

	data := struct {
		Title string
	}{
		Title: "Hello world!",
	}
	er = t.Execute(&genCode, data)
	check(er)

	d.log.Debug("contents of genCode buffer: ", genCode.String())

	// Create tar structure
	var files = []fileEntry{
		{"/cmd/agent/main.go", genCode.String()},
	}
	//append a struc of name/body for evy new plugin in project
	for i := 0; i < len(val.Plugin); i++ {
		pluginDirectoryName := 	val.Plugin[i].PluginName
		pluginDocEntry := fileEntry{
			"/plugins/" +  pluginDirectoryName + "/doc.go",
			"Doc file for package description",
		}
		pluginOptionsEntry := fileEntry{
			"/plugins/" +  pluginDirectoryName + "/options.go",
			"Config file for plugin",
		}
		pluginImplEntry := fileEntry{
			"/plugins/" +  pluginDirectoryName + "/plugin_impl_test.go",
			"Plugin file that holds main functions",
		}

			files = append(files, pluginDocEntry, pluginOptionsEntry,pluginImplEntry)
	}

		// Loop through files & write to tar
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	// Close once done & turn into []byte
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Encode to base64 string
	encodedTar := base64.StdEncoding.EncodeToString([]byte(buf.String()))

	// Create template
	template := &model.Template{
		Name:     val.GetProjectName(),
		Id:       1,
		Version:  2.4,
		Category: "health",
		Dependencies: []string{
			"grpc",
			"kafka",
			"Logrus",
		},
		TarFile: encodedTar,
	}

	// Put new value in etcd
	err := d.broker.Put(val.GetProjectName(), template)
	if err != nil {
		d.log.Errorf("Could not create template")
		return err
	}
	d.log.Infof("Return data, Key: %q Value: %+v", val.GetProjectName(), template)

	return nil
}

// GenDelProj removes a generated project in /template prefix
func (d *ProjectHandler) GenDelProj(val *model.Project) error {
	existed, err := d.broker.Delete(val.GetProjectName())
	if err != nil {
		d.log.Errorf("Could not delete template")
		return err
	}
	d.log.Infof("Delete project successful: ", existed)

	return nil
}


