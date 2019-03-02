package component

import (
	"testing"

	"github.com/lotteryjs/wechat-oauth/model"
)

func TestGetAuthorizeURL(t *testing.T) {
	type args struct {
		authorize *model.Authorize
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAuthorizeURL(tt.args.authorize); got != tt.want {
				t.Errorf("GetAuthorizeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
