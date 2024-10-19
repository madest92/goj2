// Package render Jinja2
package render

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/noirbizarre/gonja"
	"gopkg.in/yaml.v3"
)

// Template generates result data from jinj2 template
func Template(from, to string, varsFiles []string) {
	// Load variables from YAML file
	var mergedVars map[string]interface{}
	for _, file := range varsFiles {
		var vars map[string]interface{}
		varsData, err := os.ReadFile(filepath.Clean(file))
		if err != nil {
			log.Fatalf("Error reading variables file %q: %v", file, err)
		}
		err = yaml.Unmarshal(varsData, &vars)
		if err != nil {
			log.Fatalf("Error parsing variables file %q: %v", file, err)
		}
		mergedVars = mergeMap(mergedVars, vars)
	}

	// Render the template with the variables
	tpl, err := gonja.FromFile(from)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	rendered, err := tpl.Execute(mergedVars)
	if err != nil {
		log.Fatalf("Error rendering template: %v", err)
	}

	// Write the rendered template to the output file or print to console
	if to == "" {
		fmt.Println(rendered)
	} else {
		err = os.WriteFile(to, []byte(rendered), 0600)
		if err != nil {
			log.Fatalf("Error writing output file %q: %v", to, err)
		}
		fmt.Printf("Template rendered and saved to %s\n", to)
	}
}

// mergeMap merges two maps recursively
func mergeMap(dst, src map[string]interface{}) map[string]interface{} {
	if dst == nil {
		return src
	}
	for k, v := range src {
		switch v.(type) {
		case map[string]interface{}:
			if dst[k] == nil {
				dst[k] = v
				continue
			}
			dst[k] = mergeMap(dst[k].(map[string]interface{}), v.(map[string]interface{}))
		default:
			dst[k] = v
		}
	}
	return dst
}
