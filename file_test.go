package simplefileserver

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUploadFiles(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UploadFiles(tt.args.c)
		})
	}
}
