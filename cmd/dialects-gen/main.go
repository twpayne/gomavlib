package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

var tplTest = template.Must(template.New("").Parse(
	`//autogenerated:yes
//nolint:revive
package dialects

import (
    "testing"

    "github.com/stretchr/testify/require"

    "github.com/aler9/gomavlib/pkg/dialect"
{{- range .Dialects }}
	"github.com/aler9/gomavlib/pkg/dialects/{{ . }}"
{{- end }}
)

func TestDialects(t *testing.T) {
{{- range .Dialects }}
	func() {
		_, err := dialect.NewDecEncoder({{ . }}.Dialect)
		require.NoError(t, err)
	}()
{{- end }}
}
`))

func writeTemplate(fpath string, tpl *template.Template, args map[string]interface{}) error {
	f, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tpl.Execute(f, args)
}

func shellCommand(cmdstr string) error {
	fmt.Fprintf(os.Stderr, "%s\n", cmdstr)
	cmd := exec.Command("sh", "-c", cmdstr)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func downloadJSON(addr string, data interface{}) error {
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(data)
}

func processDialect(commit string, name string) error {
	fmt.Fprintf(os.Stderr, "[%s]\n", name)

	pkgName := strings.ReplaceAll(strings.ToLower(name), "_", "")

	os.Mkdir(filepath.Join("pkg", "dialects", pkgName), 0o755)

	err := shellCommand(fmt.Sprintf("go run ./cmd/dialect-import --pwd=%s --package=%s --comment=\"%s\" %s",
		filepath.Join("pkg", "dialects", pkgName),
		pkgName,
		"Package "+pkgName+" contains the "+name+" dialect.",
		"https://raw.githubusercontent.com/mavlink/mavlink/"+commit+"/message_definitions/v1.0/"+name+".xml",
	))
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "\n")
	return nil
}

func run() error {
	err := shellCommand("rm -rf pkg/dialects/*/")
	if err != nil {
		return err
	}

	var res struct {
		Sha string `json:"sha"`
	}
	err = downloadJSON("https://api.github.com/repos/mavlink/mavlink/commits/master", &res)
	if err != nil {
		return err
	}

	var files []struct {
		Name string `json:"name"`
	}
	err = downloadJSON("https://api.github.com/repos/mavlink/mavlink/contents/message_definitions/v1.0?ref="+res.Sha,
		&files)
	if err != nil {
		return err
	}

	var dialects []string //nolint:prealloc

	for _, f := range files {
		if !strings.HasSuffix(f.Name, ".xml") {
			continue
		}
		name := f.Name[:len(f.Name)-len(".xml")]

		err = processDialect(res.Sha, name)
		if err != nil {
			return err
		}

		dialects = append(dialects, strings.ReplaceAll(strings.ToLower(name), "_", ""))
	}

	err = writeTemplate(
		filepath.Join("pkg", "dialects", "package_test.go"),
		tplTest,
		map[string]interface{}{
			"Dialects": dialects,
		})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %s\n", err)
		os.Exit(1)
	}
}
