// Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package usertools_test

import (
	"bpm/usertools"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

var _ = Describe("UserIDFinder", func() {
	var userFinder *usertools.UserFinder

	BeforeEach(func() {
		userFinder = usertools.NewUserFinder()
	})

	Context("Lookup", func() {
		It("returns a runc spec User", func() {
			user, err := userFinder.Lookup("vcap")
			Expect(err).NotTo(HaveOccurred())
			Expect(user).To(Equal(specs.User{
				UID:      2000,
				GID:      3000,
				Username: "vcap",
			}))
		})

		Context("when the user lookup fails", func() {
			It("returns an error", func() {
				_, err := userFinder.Lookup("")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})