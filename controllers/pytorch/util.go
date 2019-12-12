/*
Copyright 2019 The Alibaba Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pytorch

import v1 "github.com/alibaba/kubedl/api/pytorch/v1"

func ContainMasterSpec(job *v1.PyTorchJob) bool {
	_, ok := job.Spec.PyTorchReplicaSpecs[v1.PyTorchReplicaTypeMaster]
	return ok
}
