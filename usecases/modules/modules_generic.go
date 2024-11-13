//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package modules

import "github.com/weaviate/weaviate/entities/modulecapabilities"

func (p *Provider) implementsVectorizer(mod modulecapabilities.Module) bool {
	switch mod.(type) {
	case modulecapabilities.Vectorizer[[]float32], modulecapabilities.Vectorizer[[][]float32]:
		return true
	default:
		return false
	}
}
