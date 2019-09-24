/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: manager.go
 * @time: 2019/9/23 15:08
 */
package mgr

import (
	"reflect"
	"testing"
)

func TestGetManagerInst(t *testing.T) {
	tests := []struct {
		name string
		want *Manager
	}{
		// TODO: Add test cases.
		{
			name:"TestGetManagerInst",
			want:mgrInst,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetManagerInst(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetManagerInst() = %v, want %v", got, tt.want)
			}
		})
	}
}
