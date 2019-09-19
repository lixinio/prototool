// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package lint

import (
	"strings"

	"github.com/emicklei/proto"
	"github.com/uber/prototool/internal/text"
)

var serviceNamesServiceSuffixLinter = NewLinter(
	"SERVICE_NAMES_SERVICE_SUFFIX",
	`Verifies that all service names end with "Service".`,
	checkServiceNamesServiceSuffix,
)

func checkServiceNamesServiceSuffix(add func(*text.Failure), dirPath string, descriptors []*FileDescriptor) error {
	return runVisitor(serviceNamesServiceSuffixVisitor{baseAddVisitor: newBaseAddVisitor(add)}, descriptors)
}

type serviceNamesServiceSuffixVisitor struct {
	baseAddVisitor
}

func (v serviceNamesServiceSuffixVisitor) VisitService(service *proto.Service) {
	if !strings.HasSuffix(service.Name, "Service") {
		if strings.HasSuffix(service.Name, "API") {
			v.AddFailuref(service.Position, `Service name %q must end with "Service". Since it currently ends with "API", this likely means replacing "API" with "Service"`, service.Name)
		} else {
			v.AddFailuref(service.Position, `Service name %q must end with "Service".`, service.Name)
		}
	}
}
