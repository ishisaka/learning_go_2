package internal_package_example

// ここからもinternalパッケージにはアクセスできない
import "github.com/learning-go-book-2e/internal_example/foo/internal"

func FailUseDoubler(i int) int {
	return internal.Doubler(i)
}
