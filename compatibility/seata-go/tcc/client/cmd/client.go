/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"

	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"github.com/dubbogo/gost/log/logger"

	cli "github.com/seata/seata-go/pkg/client"
	"github.com/seata/seata-go/pkg/tm"
)

// need to setup environment variable "DUBBO_GO_CONFIG_PATH" to "seata-go/tcc/client/conf/dubbogo.yml"
// and run "seata-go/tcc/server/cmd/server.go" before run
func main() {
	cli.InitPath("../../../conf/seatago.yml")
	ins, err := dubbo.NewInstance(
		dubbo.WithName("dubbo_seata_client"),
	)
	if err != nil {
		panic(err)
	}
	cli, err := ins.NewClient(
		client.WithClientURL("127.0.0.1:20000"),
		client.WithClientProtocolDubbo(),
	)
	if err != nil {
		panic(err)
	}
	conn, err := cli.Dial("UserProvider", client.WithSerialization(constant.Hessian2Serialization))
	if err != nil {
		panic(err)
	}
	test(conn)
}

type connection struct {
	conn string
}

func test(conn *client.Connection) {
	ctx := context.WithValue(context.Background(), connection{"conn"}, conn)
	tm.WithGlobalTx(ctx, &tm.GtxConfig{
		Name: "DubboGoSeataSampleLocalTx",
	}, business)
	<-make(chan struct{})
}

func business(ctx context.Context) (re error) {
	var resp bool
	conn := ctx.Value(connection{"conn"}).(*client.Connection)
	if re := conn.CallUnary(context.Background(), []interface{}{1}, &resp, "Prepare"); re != nil {
		logger.Infof("response prepare: %v", re)
	} else {
		logger.Infof("get resp %#v", resp)
	}
	return
}
