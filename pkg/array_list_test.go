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

package collection_test

import (
	"context"
	"math/rand"
	"testing"

	collection "github.com/goctus/collection/pkg"
	"github.com/stretchr/testify/assert"
)

func TestArrayList(t *testing.T) {
	t.Run("Size returns correct value with no element", func(t *testing.T) {
		assert.Equal(t, 0, collection.NewArrayList([]any{}).Size())
	})
	t.Run("Size returns correct value with one element", func(t *testing.T) {
		assert.Equal(t, 1, collection.NewArrayList([]string{"Hello, World!"}).Size())
	})
	t.Run("Size returns correct value with many elements", func(t *testing.T) {
		assert.Equal(t, 3, collection.NewArrayList([]int{0, 0, 0}).Size())
	})
	t.Run("With returns the same list without arguments", func(t *testing.T) {
		subject := collection.NewArrayList([]int{0})
		assert.Equal(
			t,
			subject.Size(),
			subject.With().Size(),
			"ArrayList should not change its size when calling With without any argument",
		)
	})
	t.Run("With does not modify the original list", func(t *testing.T) {
		a := collection.NewArrayList([]int{})
		b := a.With(0)
		assert.NotEqual(
			t,
			a.Size(),
			b.Size(),
			"With should not modify the origin list",
		)
	})
	t.Run("Without doesn't break when called on an empty list", func(t *testing.T) {
		assert.NotPanics(t, func() {
			collection.NewArrayList([]int{}).Without(0)
		})
	})
	t.Run("Without removes the first element", func(t *testing.T) {
		assert.Equal(t, 1, collection.NewArrayList([]int{0, 0}).Without(0).Size())
	})
	t.Run("Without removes the last element", func(t *testing.T) {
		assert.Equal(t, 1, collection.NewArrayList([]int{0, 0}).Without(1).Size())
	})
	t.Run("Without removes the only element", func(t *testing.T) {
		assert.Equal(t, 0, collection.NewArrayList([]int{0}).Without(0).Size())
	})
	t.Run("Found returns the only element", func(t *testing.T) {
		want := rand.Int()
		subject := collection.NewArrayList([]int{want}).Found(0)
		actual, err := subject.Value(context.TODO())
		if err != nil {
			t.Errorf("Didn't expect to receive an error: %v", err)
		}
		assert.Equal(t, want, actual)
	})
	t.Run("Found returns an empty Option on wrong index", func(t *testing.T) {
		_, err := collection.NewArrayList([]int{}).Found(0).Value(context.TODO())
		assert.ErrorIs(t, collection.ErrNoSuchElement, err)
	})
}
