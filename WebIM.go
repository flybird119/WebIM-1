// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// This sample is about using long polling and WebSocket to build a web-based chat room based on beego.
package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"WebIM/models"
	_ "WebIM/routers"
	//导包的时候默认执行该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，
	// 仅仅是是希望它执行init()函数而已。这个时候就可以使用 import _ 引用该包。即使用【import _ 包路径】只是引用该包，
	// 仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。

	"github.com/astaxie/beego/orm"
)

const (
	APP_VER = "0.1.1.0227"
)

func init()  {
	models.RegisterDB()

}
func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)
	orm.RunSyncdb("default", false, true)
	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
