package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"io/ioutil"
	"os"
	"shequn1/foundation/app"
	"shequn1/foundation/database/orm"
	"shequn1/foundation/server"
	"shequn1/foundation/util"
	"shequn1/internal/model"
)

var (
	cfgPath = "runtime/cfg"
	cfgFile = "system_cfg.json"
)

func GetSystemCfg() g.Map {
	setting := model.SystemSetting{}
	systemCfg := model.SystemCfg{Label: "app"}
	orm.Master().Table("system_cfgs").Where("label = ?", "app").First(&systemCfg)
	err := json.Unmarshal([]byte(systemCfg.Cfg), &setting)
	if err != nil {
		app.Logger().Error(err.Error())
	}
	systemCfgMap := gconv.Map(setting)
	if setting.DefaultWechatQrCodeUrl != "" {
		systemCfgMap["thumb_url"] = fmt.Sprintf("%s://%s%s", server.Config.Schema, server.Config.Domain, setting.DefaultWechatQrCodeUrl)
	}
	return systemCfgMap
}

func SaveSystemCfg(bizAttr g.Map) error {
	systemCfg := model.SystemCfg{}
	bizAttr["label"] = "app"
	tmpJsonByte, _ := json.Marshal(bizAttr)
	settingCfg := string(tmpJsonByte)
	systemCfg.Label, _ = bizAttr["label"].(string)
	orm.Master().Table("system_cfgs").FirstOrCreate(&systemCfg, "label = ?", bizAttr["label"])
	systemCfg.Cfg = settingCfg
	err := orm.Master().Table("system_cfgs").Save(&systemCfg).Error
	if dirExits := util.Exists(cfgPath); !dirExits {
		os.MkdirAll(cfgPath, 0777)
	}
	os.WriteFile(cfgPath+"/"+cfgFile, tmpJsonByte, 0777)
	if err != nil {
		app.Logger().Error(err.Error())
		return errors.New("保存配置失败")
	}
	return nil
}

func GetSystemCfgJson() g.Map {
	setting := model.SystemSetting{}
	fp, err := os.OpenFile(cfgPath+"/"+cfgFile, os.O_RDWR, 0777)
	if err != nil {
		app.Logger().Error(err.Error())
	}
	cfgJsonbyte, _ := ioutil.ReadAll(fp)
	cfgJson := string(cfgJsonbyte)
	if cfgJson == "" {
		systemCfg := model.SystemCfg{}
		orm.Master().Table("system_cfgs").Where("label = ?", "app").First(&systemCfg)
		cfgJsonbyte = []byte(systemCfg.Cfg)
		os.WriteFile(cfgPath+"/"+cfgFile, []byte(systemCfg.Cfg), 0777)
	}

	json.Unmarshal(cfgJsonbyte, &setting)
	return gconv.Map(setting)
}
