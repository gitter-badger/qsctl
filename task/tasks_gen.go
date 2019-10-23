// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates types, It can be invoked by running
// go generate
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"
)

type task struct {
	Name           string   `json:"-"`
	Type           string   `json:"type"`
	Path           string   `json:"path"`
	Depend         string   `json:"depend,omitempty"`
	Description    string   `json:"description"`
	InheritedValue []string `json:"inherited_value,omitempty"`
	RuntimeValue   []string `json:"runtime_value,omitempty"`
}

var funcs = template.FuncMap{
	"lowerFirst": func(s string) string {
		if len(s) == 0 {
			return ""
		}
		if s[0] < 'A' || s[0] > 'Z' {
			return s
		}
		return string(s[0]+'a'-'A') + s[1:]
	},
}

//go:generate go run tasks_gen.go
func main() {
	data, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	tasks := make(map[string]*task)
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	// Do sort to all tasks via name.
	taskNames := make([]string, 0)
	for k := range tasks {
		sort.Strings(tasks[k].InheritedValue)
		sort.Strings(tasks[k].RuntimeValue)

		taskNames = append(taskNames, k)
	}
	sort.Strings(taskNames)

	// Set task name and categorized via path.
	pages := make(map[string][]*task)
	for _, v := range taskNames {
		tasks[v].Name = v
		pages[tasks[v].Path] = append(pages[tasks[v].Path], tasks[v])
	}

	// Format input tasks.json
	data, err = json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("tasks.json", data, 0664)
	if err != nil {
		log.Fatal(err)
	}

	for pathName, page := range pages {
		err := os.MkdirAll(path.Dir(pathName), 0664)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Create(pathName + "_generated.go")
		if err != nil {
			log.Fatal(err)
		}

		packageName := "task"
		if strings.Contains(pathName, "/") {
			packageName = path.Dir(pathName)
		}

		// Write page temple firstly.
		err = pageTmpl.Execute(f, struct {
			Package string
		}{
			packageName,
		})
		if err != nil {
			log.Fatal(err)
		}

		// Write task.
		for _, task := range page {
			err = requirementTmpl.Execute(f, task)
			if err != nil {
				log.Fatal(err)
			}
			err = mockTmpl.Execute(f, task)
			if err != nil {
				log.Fatal(err)
			}
			err = taskTmpl[task.Type].Execute(f, task)
			if err != nil {
				log.Fatal(err)
			}
		}

		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	for pathName, page := range pages {
		err := os.MkdirAll(path.Dir(pathName), 0664)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Create(pathName + "_generated_test.go")
		if err != nil {
			log.Fatal(err)
		}

		packageName := "task"
		if strings.Contains(pathName, "/") {
			packageName = path.Dir(pathName)
		}

		// Write page temple firstly.
		err = testPageTmpl.Execute(f, struct {
			Package string
		}{
			packageName,
		})
		if err != nil {
			log.Fatal(err)
		}

		// Write task.
		for _, task := range page {
			err = taskTestTmpl[task.Type].Execute(f, task)
			if err != nil {
				log.Fatal(err)
			}
		}

		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

var taskTmpl = map[string]*template.Template{
	"required":  requiredTaskTmpl,
	"dependent": dependentTaskTmpl,
}

var taskTestTmpl = map[string]*template.Template{
	"required":  requiredTaskTestTmpl,
	"dependent": dependentTaskTestTmpl,
}

var pageTmpl = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package {{ .Package }}

import (
	"fmt"

	"github.com/Xuanwo/navvy"
	"github.com/google/uuid"

	"github.com/yunify/qsctl/v2/pkg/types"
	"github.com/yunify/qsctl/v2/utils"
)

var _ navvy.Pool
var _ types.Pool
var _ = utils.SubmitNextTask
var _ = uuid.New()
`))

var testPageTmpl = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package {{ .Package }}

import (
	"errors"
	"testing"

	"github.com/Xuanwo/navvy"
	"github.com/stretchr/testify/assert"

	"github.com/yunify/qsctl/v2/pkg/types"
	"github.com/yunify/qsctl/v2/utils"
)

var _ navvy.Pool
var _ types.Pool
var _ = utils.SubmitNextTask
`))

var requirementTmpl = template.Must(template.New("").Funcs(funcs).Parse(`
// {{ .Name | lowerFirst }}TaskRequirement is the requirement for execute {{ .Name }}Task.
type {{ .Name | lowerFirst }}TaskRequirement interface {
	navvy.Task
{{- if eq .Type "required" }}
	types.Todoist
	types.PoolGetter
	types.FaultSetter
	types.FaultValidator
	types.IDGetter
{{ else }}
{{- if .Depend }}
	types.PoolGetter
{{- end }}
{{- end }}

	// Inherited value
{{- range $k, $v := .InheritedValue }}
	types.{{$v}}Getter
{{- end }}

{{- if eq .Type "required" }}
	// Runtime value
{{- range $k, $v := .RuntimeValue }}
	types.{{$v}}Setter
{{- end }}
{{- end }}
}
`))

var mockTmpl = template.Must(template.New("").Funcs(funcs).Parse(`
// mock{{ .Name }}Task is the mock task for {{ .Name }}Task.
type mock{{ .Name }}Task struct {
	types.Todo
	types.Pool
	types.Fault
	types.ID

	// Inherited value
{{- range $k, $v := .InheritedValue }}
	types.{{$v}}
{{- end }}

{{- if eq .Type "required" }}
	// Runtime value
{{- range $k, $v := .RuntimeValue }}
	types.{{$v}}
{{- end }}
{{- end }}
}

func (t *mock{{ .Name }}Task) Run() {
	panic("mock{{ .Name }}Task should not be run.")
}
`))

var requiredTaskTmpl = template.Must(template.New("").Funcs(funcs).Parse(`
// {{ .Name }}Task will {{ .Description }}.
type {{ .Name }}Task struct {
	{{ .Name | lowerFirst }}TaskRequirement
}

// Run implement navvy.Task.
func (t *{{ .Name }}Task) Run() {
	t.run()
	if t.ValidateFault() {
		return
	}
	utils.SubmitNextTask(t.{{ .Name | lowerFirst }}TaskRequirement)
}

func (t *{{ .Name }}Task) TriggerFault(err error) {
	t.SetFault(fmt.Errorf("Task {{ .Name }} failed: {%w}", err))
}

// New{{ .Name }}Task will create a new {{ .Name }}Task.
func New{{ .Name }}Task(task types.Todoist) navvy.Task {
	return &{{ .Name }}Task{task.({{ .Name | lowerFirst }}TaskRequirement)}
}
`))

var requiredTaskTestTmpl = template.Must(template.New("").Funcs(funcs).Parse(`
func TestNew{{ .Name }}Task(t *testing.T) {
	m := &mock{{ .Name }}Task{}
	task := New{{ .Name }}Task(m)
	assert.NotNil(t, task)
}

func Test{{ .Name }}Task_TriggerFault(t *testing.T) {
	m := &mock{{ .Name }}Task{}
	task := &{{ .Name }}Task{m}
	err := errors.New("test error")
	task.TriggerFault(err)
	assert.True(t, task.{{ .Name | lowerFirst }}TaskRequirement.ValidateFault())
}

func TestMock{{ .Name }}Task_Run(t *testing.T) {
	task := &mock{{ .Name }}Task{}
	assert.Panics(t, func() {
		task.Run()
	})
}
`))

var dependentTaskTmpl = template.Must(template.New("").Funcs(funcs).Parse(`
// {{ .Name }}Task will {{ .Description }}.
type {{ .Name }}Task struct {
	{{ .Name | lowerFirst }}TaskRequirement

	// Predefined runtime value
	types.Fault
	types.ID
	types.Todo

	// Runtime value
{{- range $k, $v := .RuntimeValue }}
	types.{{$v}}
{{- end }}
}

// Run implement navvy.Task
func (t *{{ .Name }}Task) Run() {
	if t.ValidateFault() {
		return
	}
	utils.SubmitNextTask(t)
}

func (t *{{ .Name }}Task) TriggerFault(err error) {
	t.SetFault(fmt.Errorf("Task {{ .Name }} failed: {%w}", err))
}

{{- if .Depend }}
// New{{ .Name }}Task will create a {{ .Name }}Task and fetch inherited data from {{ .Depend }}Task.
func New{{ .Name }}Task(task types.Todoist) navvy.Task {
	t := &{{ .Name }}Task{
		{{ .Name | lowerFirst }}TaskRequirement: task.({{ .Name | lowerFirst }}TaskRequirement),
	}
	t.SetID(uuid.New().String())
	t.new()
	return t
}
{{- else }}
// Wait will wait until {{ .Name }}Task has been finished
func (t *{{ .Name }}Task) Wait() {
	t.GetPool().Wait()
}
{{- end }}
`))

var dependentTaskTestTmpl = template.Must(template.New("").Funcs(funcs).Parse(`
func Test{{ .Name }}Task_GeneratedRun(t *testing.T) {
	cases := []struct {
		name     string
		hasFault bool
		hasCall  bool
		gotCall  bool
	}{
		{
			"has fault",
			true,
			false,
			false,
		},
		{
			"no fault",
			false,
			true,
			false,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			pool := navvy.NewPool(10)

			{{- if .Depend }}
			m := &mock{{ .Name }}Task{}
			m.SetPool(pool)
			task := &{{ .Name }}Task{ {{ .Name | lowerFirst }}TaskRequirement: m}
			{{- else }}
			task := &{{ .Name }}Task{}
			task.SetPool(pool)
			{{- end }}

			err := errors.New("test error")
			if v.hasFault {
				task.SetFault(err)
			}
			task.AddTODOs(func(todoist types.Todoist) navvy.Task {
				x := utils.NewCallbackTask(func() {
					v.gotCall = true
				})
				return x
			})

			task.Run()
			pool.Wait()

			assert.Equal(t, v.hasCall, v.gotCall)
		})
	}
}

func Test{{ .Name }}Task_TriggerFault(t *testing.T) {
	err := errors.New("trigger fault")
	x := &{{ .Name }}Task{}
	x.TriggerFault(err)

	assert.Equal(t, true, x.ValidateFault())
	assert.Equal(t, true, errors.Is(x.GetFault(), err))
}

func TestMock{{ .Name }}Task_Run(t *testing.T) {
	task := &mock{{ .Name }}Task{}
	assert.Panics(t, func() {
		task.Run()
	})
}
`))
