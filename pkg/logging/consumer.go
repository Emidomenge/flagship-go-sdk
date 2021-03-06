/****************************************************************************
 * Copyright 2020, AB Tasty, Inc. and contributors                        *
 *                                                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");          *
 * you may not use this file except in compliance with the License.         *
 * You may obtain a copy of the License at                                  *
 *                                                                          *
 *    http://www.apache.org/licenses/LICENSE-2.0                            *
 *                                                                          *
 * Unless required by applicable law or agreed to in writing, software      *
 * distributed under the License is distributed on an "AS IS" BASIS,        *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. *
 * See the License for the specific language governing permissions and      *
 * limitations under the License.                                           *
 ***************************************************************************/

// Package logging //
package logging

import (
	"fmt"
	"io"
	"log"
)

// FilteredLevelLogConsumer is an implementation of the FlagshipLogConsumer that filters by log level
type FilteredLevelLogConsumer struct {
	level  LogLevel
	logger *log.Logger
}

// Log logs the message depengind on log level
func (l *FilteredLevelLogConsumer) Log(level LogLevel, message string, name string) {
	if l.level <= level {
		// prepends the name and log level to the message
		message = fmt.Sprintf("[%s][%s] %s", level, name, message)
		l.logger.Println(message)
	}
}

// SetLogLevel changes the log level to the given level
func (l *FilteredLevelLogConsumer) SetLogLevel(level LogLevel) {
	l.level = level
}

// NewFilteredLevelLogConsumer returns a new logger that logs to stdout
func NewFilteredLevelLogConsumer(level LogLevel, out io.Writer) *FilteredLevelLogConsumer {
	return &FilteredLevelLogConsumer{
		level:  level,
		logger: log.New(out, "[Flagship]", log.LstdFlags),
	}
}
