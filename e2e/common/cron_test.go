// +build integration

// To enable compilation of this file in Goland, go to "Settings -> Go -> Vendoring & Build Tags -> Custom Tags" and add "integration"

/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"testing"

	. "github.com/apache/camel-k/e2e/support"
	camelv1 "github.com/apache/camel-k/pkg/apis/camel/v1"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
)

func TestRunCronExample(t *testing.T) {
	WithNewTestNamespace(t, func(ns string) {
		Expect(Kamel("install", "-n", ns).Execute()).Should(BeNil())

		t.Run("cron", func(t *testing.T) {
			RegisterTestingT(t)

			Expect(Kamel("run", "-n", ns, "files/cron.groovy").Execute()).Should(BeNil())
			Eventually(IntegrationCronJob(ns, "cron"), TestTimeoutMedium).ShouldNot(BeNil())
			Eventually(IntegrationCondition(ns, "cron", camelv1.IntegrationConditionReady), TestTimeoutMedium).Should(Equal(v1.ConditionTrue))
			Eventually(IntegrationLogs(ns, "cron"), TestTimeoutMedium).Should(ContainSubstring("Magicstring!"))
			Expect(Kamel("delete", "--all", "-n", ns).Execute()).Should(BeNil())
		})

		t.Run("cron-quarkus", func(t *testing.T) {
			RegisterTestingT(t)

			Expect(Kamel("run", "--name", "cron-quarkus", "-n", ns, "files/cron.groovy", "-t", "quarkus.enabled=true").Execute()).Should(BeNil())
			Eventually(IntegrationCronJob(ns, "cron-quarkus"), TestTimeoutMedium).ShouldNot(BeNil())
			Eventually(IntegrationCondition(ns, "cron-quarkus", camelv1.IntegrationConditionReady), TestTimeoutMedium).Should(Equal(v1.ConditionTrue))
			Eventually(IntegrationLogs(ns, "cron-quarkus"), TestTimeoutMedium).Should(ContainSubstring("Magicstring!"))
			Expect(Kamel("delete", "--all", "-n", ns).Execute()).Should(BeNil())
		})

		t.Run("cron-timer", func(t *testing.T) {
			RegisterTestingT(t)

			Expect(Kamel("run", "-n", ns, "files/cron-timer.groovy").Execute()).Should(BeNil())
			Eventually(IntegrationCronJob(ns, "cron-timer"), TestTimeoutMedium).ShouldNot(BeNil())
			Eventually(IntegrationCondition(ns, "cron-timer", camelv1.IntegrationConditionReady), TestTimeoutMedium).Should(Equal(v1.ConditionTrue))
			Eventually(IntegrationLogs(ns, "cron-timer"), TestTimeoutMedium).Should(ContainSubstring("Magicstring!"))
			Expect(Kamel("delete", "--all", "-n", ns).Execute()).Should(BeNil())
		})

		t.Run("cron-fallback", func(t *testing.T) {
			RegisterTestingT(t)

			Expect(Kamel("run", "-n", ns, "files/cron-fallback.groovy").Execute()).Should(BeNil())
			Eventually(IntegrationPodPhase(ns, "cron-fallback"), TestTimeoutMedium).Should(Equal(v1.PodRunning))
			Eventually(IntegrationCondition(ns, "cron-fallback", camelv1.IntegrationConditionReady), TestTimeoutShort).Should(Equal(v1.ConditionTrue))
			Eventually(IntegrationLogs(ns, "cron-fallback"), TestTimeoutShort).Should(ContainSubstring("Magicstring!"))
			Expect(Kamel("delete", "--all", "-n", ns).Execute()).Should(BeNil())
		})

		t.Run("cron-fallback-quarkus", func(t *testing.T) {
			RegisterTestingT(t)

			Expect(Kamel("run", "--name", "cron-fallback-quarkus", "-n", ns, "files/cron-fallback.groovy", "-t", "quarkus.enabled=true").Execute()).Should(BeNil())
			Eventually(IntegrationPodPhase(ns, "cron-fallback-quarkus"), TestTimeoutMedium).Should(Equal(v1.PodRunning))
			Eventually(IntegrationCondition(ns, "cron-fallback-quarkus", camelv1.IntegrationConditionReady), TestTimeoutShort).Should(Equal(v1.ConditionTrue))
			Eventually(IntegrationLogs(ns, "cron-fallback-quarkus"), TestTimeoutShort).Should(ContainSubstring("Magicstring!"))
			Expect(Kamel("delete", "--all", "-n", ns).Execute()).Should(BeNil())
		})

		// Enable this test when https://github.com/apache/camel-k-runtime/issues/346 is fixed (Camel K Runtime 1.3.x)
		/*
			t.Run("cron-quartz-quarkus", func(t *testing.T) {
				RegisterTestingT(t)

				Expect(Kamel("run", "--name", "cron-quartz-quarkus", "-n", ns, "files/cron-quartz.groovy", "-t", "quarkus.enabled=true").Execute()).Should(BeNil())
				Eventually(IntegrationPodPhase(ns, "cron-quartz-quarkus"), TestTimeoutMedium).Should(Equal(v1.PodRunning))
				Eventually(IntegrationCondition(ns, "cron-quartz-quarkus", camelv1.IntegrationConditionReady), TestTimeoutShort).Should(Equal(v1.ConditionTrue))
				Eventually(IntegrationLogs(ns, "cron-quartz-quarkus"), TestTimeoutShort).Should(ContainSubstring("Magicstring!"))
				Expect(Kamel("delete", "--all", "-n", ns).Execute()).Should(BeNil())
			})
		*/
	})
}
