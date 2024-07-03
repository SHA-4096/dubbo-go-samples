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

package config

import (
	"fmt"
)

type Config struct {
	Port  int
	Redis RedisConfig
}

type RedisConfig struct {
	Addr string
}

var cfg *Config = &Config{
	Port: 7070,
}

func Address() string {
	return fmt.Sprintf(":%d", cfg.Port)
}

func Redis() RedisConfig {
	return cfg.Redis
}

// TODO:Config
func Load() error {
	return nil
}
