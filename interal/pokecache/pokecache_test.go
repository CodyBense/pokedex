package pokecache

import "testing"

func TestAddGet(t *testing.T) {
	cache := NewCache()
	cache.Add("first", []byte("original"))
	cache.Add("second", []byte("other"))
	cache.Add("first", []byte("updated"))

	for _, tc := range []struct {
		key   string
		want  string
		found bool
	}{
		{key: "first", want: "updated", found: true},
		{key: "second", want: "other", found: true},
		{key: "missing", found: false},
	} {
		t.Run(tc.key, func(t *testing.T) {
			got, ok := cache.Get(tc.key)
			if ok != tc.found || string(got) != tc.want {
				t.Fatalf("Get(%q) = %q, %v; want %q, %v", tc.key, got, ok, tc.want, tc.found)
			}
			if !ok && got != nil {
				t.Fatal("missing key should return nil")
			}
		})
	}
}
