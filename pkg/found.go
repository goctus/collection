/*
 * MIT License
 *
 * Copyright (c) 2023-present goctus
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package collection

import (
	"context"

	option "github.com/goctus/option/pkg"
)

// Found is an option of an entry of a collection.
type Found[K any, V any] struct {
	collection Collecton[K, V]
	entry      K

	cache option.Option[V]
}

// NewFound creates a new Found.
func NewFound[K any, V any](collection Collecton[K, V], entry K) Found[K, V] {
	return Found[K, V]{collection, entry, nil}
}

func (f Found[K, V]) Empty() bool {
	return f.origin().Empty()
}

func (f Found[K, V]) Value(ctx context.Context) (V, error) {
	return f.origin().Value(ctx)
}

func (f Found[K, V]) origin() option.Option[V] {
	if f.cache == nil {
		f.cache = f.collection.Found(f.entry)
	}
	return f.cache
}
