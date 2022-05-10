package customer

import (
	"fmt"
	"math/rand"
	"testing"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomCustomers(n int) {
	for i := 0; i < n; i++ {
		p := New()
		p.SetFirstName(RandomString(5))
		p.SetLastName(RandomString(5))
		p.SetEmail(RandomString(5))

		result, _ = GetByID(p.ID)
	}
	result = customer{}
}

var result customer

func BenchmarkRestore(b *testing.B) {
	customers := []int{1, 5, 10, 25, 100, 250, 1000}
	for j, c := range customers {
		mapList = true
		RandomCustomers(c)
		var tmp customer
		b.Run(fmt.Sprintf("Map-%d", c), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tmp, _ = restore(j)
			}
			result = tmp
		})

		mapList = false
		RandomCustomers(c)
		b.Run(fmt.Sprintf("List-%d", c), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tmp, _ = restore(j)
			}
			result = tmp
		})
	}
}
