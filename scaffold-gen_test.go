package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSetupParseFlags(t *testing.T) {

    testCases := []struct {
        args                   []string
        err                    error
        expectedConf           projectConfig
        expectedOutputContains string
    }{
        {
            args: []string{"-n", "MyProject", "-d", "/path/to/dir", "-r", "github.com/username/myproject"},
            err:  nil,
            expectedConf: projectConfig{
              Name:         "MyProject",
              LocalPath:    "/path/to/dir",
              RepoUrl:      "github.com/username/myproject",
              StaticAssets: false},
			},
        {
            args: []string{"-n", "MyProject", "-d", "/path/to/dir", "-r", "github.com/username/myproject"},
            err:  nil,
            expectedConf: projectConfig{
              Name:         "MyProject",
						},
        },
        // Define other test cases
    }

    byteBuf := new(bytes.Buffer)
    for _, tc := range testCases {
        c, err := setupParseFlags(byteBuf, tc.args)
        if err != nil {
					t.Errorf("Got error %v", err)
				}

				if c != tc.expectedConf {
					t.Errorf("Expected output: %v, Got: %v", c, tc.expectedConf)
				}
      
               // TODO: check if the returned config object matches expected config object
             
               // check expected output
        if len(tc.expectedOutputContains) != 0 {
					gotOutput := byteBuf.String()
					if strings.Index(gotOutput, tc.expectedOutputContains) == -1 {
							t.Errorf("Expected output: %s, Got: %s", tc.expectedOutputContains, gotOutput)
					}
        }
        byteBuf.Reset()
    }
}
