package cloud

import (
	"fmt"
	"testing"
)

func TestNewProvider(t *testing.T) {
	config := Config{
		Provider:     ALI,
		Region:       "cn-hongkong",
		AccessKey:    "LTAxxxxxxxzr0uuT",
		AccessSecret: "FN3xxxxxxxxxxxWnHs99bcYlbVZuIm",
	}
	r := Request{
		Num:        2,
		Image:      "m-j6c7cmqwpqwn8onaey27",
		Flavor:     "2C2G",
		NamePrefix: "kube",
		Passwd:     "Fanux#123",
		ZoneID:     "cn-hongkong-b",
		FIP:        true,
	}

	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		r    Request
	}{
		{
			"test provider",
			args{config},
			r,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewProvider(tt.args.config)
			res, err := got.Create(tt.r)
			if err != nil {
				t.Errorf("create vm failed:%s", err)
			}
			fmt.Println(res)
		})
	}
}
