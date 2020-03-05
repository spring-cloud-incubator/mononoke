/*
Copyright 2020 the original author or authors.

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

package opinions

import (
	"context"

	corev1 "k8s.io/api/core/v1"
)

type Opinion struct {
	Id         string
	Applicable func(imageMetadata map[string]string) bool
	Apply      func(ctx context.Context, podSpec *corev1.PodTemplateSpec, imageMetadata map[string]string) error
}

type Opinions []Opinion

func (os Opinions) Apply(ctx context.Context, podSpec *corev1.PodTemplateSpec, imageMetadata map[string]string) ([]string, error) {
	applied := []string{}
	for _, o := range os {
		if o.Applicable == nil || o.Applicable(imageMetadata) {
			applied = append(applied, o.Id)
			if err := o.Apply(ctx, podSpec, imageMetadata); err != nil {
				return nil, err
			}
		}
	}
	return applied, nil
}
