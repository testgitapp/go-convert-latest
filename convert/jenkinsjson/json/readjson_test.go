package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvertReadJson(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	filePath := filepath.Join(workingDir, "../convertTestFiles/readjson/readjsonSnippet.json")
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read JSON file: %v", err)
	}

	var node1 Node
	if err := json.Unmarshal(jsonData, &node1); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	tests := []struct {
		json Node
		want Node
	}{
		{
			json: node1,
			want: Node{
				AttributesMap: map[string]string{
					"ci.pipeline.run.user":                 "SYSTEM",
					"jenkins.pipeline.step.id":             "9",
					"jenkins.pipeline.step.name":           "Read JSON from files in the workspace.",
					"jenkins.pipeline.step.plugin.name":    "pipeline-utility-steps",
					"jenkins.pipeline.step.plugin.version": "2.17.0",
					"jenkins.pipeline.step.type":           "readJSON",
					"harness-attribute":                    "{\n  \"file\" : \"input.json\"\n}",
					"harness-others":                       "",
				},
				Name:         "ag-readJSON #18",
				Parent:       "ag-readJSON",
				ParentSpanId: "f07a51df3222232e",
				SpanId:       "b7793f5483304f58",
				SpanName:     "readJSON",
				TraceId:      "97a084b8b9fbac2ccadc2e31e648c395",
				Type:         "Run Phase Span",
				ParameterMap: map[string]any{"file": "input.json"},
			},
		},
	}

	for i, test := range tests {
		got := test.json
		if diff := cmp.Diff(got, test.want); diff != "" {
			t.Errorf("Unexpected parsing results for test %v", i)
			t.Log(diff)
		}
	}
}
