package facade

import (
	"fmt"
	"testing"
)

func TestHomeFacade_EndMovie(t *testing.T) {
	c := NewClient(NewHomeFacade(NewMovieImp(), NewLightImp()))
	fmt.Println("XiaoAi, watching movie")
	c.Hfi.WatchMovie()
	fmt.Println("XiaoAi, off movie")
	c.Hfi.EndMovie()
}
