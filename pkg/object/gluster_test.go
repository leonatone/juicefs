//go:build gluster
// +build gluster

/*
 * JuiceFS, Copyright 2023 leonatone, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package object

import (
	"os"
	"testing"
)

func TestGluster(t *testing.T) {
	if os.Getenv("GLUSTER_VOLUME") == "" {
		t.SkipNow()
	}
	b, _ := newGluster(os.Getenv("GLUSTER_VOLUME"), "", "", "")
	testStorage(t, b)

}

func TestGluster2(t *testing.T) {
	if os.Getenv("GLUSTER_VOLUME") == "" {
		t.SkipNow()
	}
	b, _ := newGluster(os.Getenv("GLUSTER_VOLUME"), "", "", "")
	testFileSystem(t, b)
}
