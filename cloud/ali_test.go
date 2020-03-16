package cloud

import (
	"fmt"
	"testing"
)

func TestAliProvider_QueryFlavor(t *testing.T) {
	config := Config{
		Provider:     ALI,
		Region:       "cn-hongkong",
		AccessKey:    "LTBOOzr0uuT",
		AccessSecret: "FNxxxxxxxxxxxxxxxxxxbcYlbVZuIm",
	}
	configBeijing := config
	configBeijing.Region = "cn-beijing"

	type fields struct {
		Config Config
	}
	type args struct {
		flavor string
	}
	tests := []struct {
		name   string
		fields fields
		zone   string
	}{
		{
			"test query flavor",
			fields{config},
			"cn-hongkong-b",
		},
		{
			"test query beijing flavor",
			fields{configBeijing},
			"cn-beijing-b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			a := NewProvider(tt.fields.Config)
			for _, f := range []string{"1C1G", "1C2G", "2C4G", "4C8G", "8C16G", "8C32G"} {
				fl := a.QueryFlavor(f, tt.zone, "PostPaid", "SpotAsPriceGo")
				fmt.Println("flavor is : ", fl)
				fmt.Println()
			}
		})
	}
}
