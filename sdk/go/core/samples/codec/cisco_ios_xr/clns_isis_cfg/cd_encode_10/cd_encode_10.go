/*
 * ------------------------------------------------------------------
 * YANG Development Kit
 * Copyright 2017 Cisco Systems. All rights reserved
 *
 *----------------------------------------------
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http:www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *----------------------------------------------
 */
package main

import (
	"flag"
	"fmt"
	isisCfg "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr/clns_isis_cfg"
	"github.com/CiscoDevNet/ydk-go/ydk/providers"
	"github.com/CiscoDevNet/ydk-go/ydk/services"
	"github.com/CiscoDevNet/ydk-go/ydk"
	encoding "github.com/CiscoDevNet/ydk-go/ydk/types/encoding_format"
)

// configIsis adds config data to isis object.
func configIsis(isis *isisCfg.Isis) {
}

// main executes main program.
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nError occured!!\n ", r)
		}
	}()

	vPtr := flag.Bool("v", false, "Enable verbose")
	flag.Parse()

	if *vPtr {
		ydk.EnableLogging(ydk.Info)
	}

	// create codec provider
	var provider = providers.CodecServiceProvider{}
	provider.Encoding = encoding.XML

	// create codec service
	var service = services.CodecService{}

	isis := isisCfg.Isis{}		// create an object
	configIsis(&isis)			// add object configuration

	_ = service
}
