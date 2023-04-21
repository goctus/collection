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
	"testing"

	collection "github.com/goctus/collection/pkg"
	"github.com/stretchr/testify/assert"
)

func TestFound(t *testing.T) {
	t.Run("finds", func(t *testing.T) {
		want := "Hello, World!"
		subject := collection.NewFound[string, string](
			collection.NewNativeMap[string](map[any]string{}).With("test", want),
			"test",
		)
		assert.False(t, subject.Empty(), "Should not be empty")
		actual, err := subject.Value(context.TODO())
		assert.Nil(t, err, "Failed to get value")
		assert.Equal(t, want, actual, "Actual value does not match the expeced")
	})
}
