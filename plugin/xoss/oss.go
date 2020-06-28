package xoss

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/geekymedic/x-lite/plugin"
)

var ossClient = map[string]*oss.Bucket{}
var opts = map[string]struct {
	EndPoint    string
	AccessKeyId string
	KeySecret   string
	BucketName  string
}{}

func init() {
	plugin.AddPlugin("oos", func(status plugin.Status, viper *viper.Viper) error {
		err := viper.UnmarshalKey("oss", &opts)
		if err != nil {
			return err
		}
		for key, opt := range opts {
			key = strings.ToLower(key)
			client, err := oss.New(opt.EndPoint, opt.AccessKeyId, opt.KeySecret)
			if err != nil {
				return err
			}
			ossClient[key], err = client.Bucket(opt.BucketName)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func Use(s string) *oss.Bucket {
	return ossClient[strings.ToLower(s)]
}
