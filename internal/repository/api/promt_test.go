package api

import (
	"agahi-plus-plus/internal/helper"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"testing"
)

func Test_promptApi_Generate(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		config *helper.ServiceConfig
	}
	type args struct {
		ctx      *gin.Context
		imageUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{{name: "test", fields: struct {
		logger *zap.Logger
		config *helper.ServiceConfig
	}{logger: zaptest.NewLogger(t), config: helper.NewServiceConfigMock()}, args: struct {
		ctx      *gin.Context
		imageUrl string
	}{
		ctx:      nil,
		imageUrl: "https://s100.divarcdn.com/static/photo/neda/post/T1sY3uy63452W-qvKBjNhA/445af152-1577-4e15-a044-4dc8ea2e4c3b.jpg"},
		want:    "error",
		wantErr: true}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := promptApi{
				logger: tt.fields.logger,
				config: tt.fields.config,
			}
			got, err := r.Generate(tt.args.ctx, tt.args.imageUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
