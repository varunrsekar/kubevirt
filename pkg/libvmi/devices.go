/*
 * This file is part of the KubeVirt project
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
 *
 * Copyright the KubeVirt Authors.
 *
 */

package libvmi

import v1 "kubevirt.io/api/core/v1"

// WithTablet adds tablet device with given name and bus
func WithTablet(name string, bus v1.InputBus) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		vmi.Spec.Domain.Devices.Inputs = append(vmi.Spec.Domain.Devices.Inputs,
			v1.Input{
				Name: name,
				Bus:  bus,
				Type: v1.InputTypeTablet,
			},
		)
	}
}

// WithPanicDevice adds a panic device with the given model
func WithPanicDevice(model v1.PanicDeviceModel) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		vmi.Spec.Domain.Devices.PanicDevices = append(vmi.Spec.Domain.Devices.PanicDevices, v1.PanicDevice{Model: &model})
	}
}
