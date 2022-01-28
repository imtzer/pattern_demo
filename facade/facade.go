// file: facade.go
// author: TaoZer
// Date: 2022/1/28
// Description: facade pattern

package facade

import "fmt"

type Client struct {
	Hfi HomeFacadeI
}

func NewClient(hfi HomeFacadeI) Client {
	return Client{
		hfi,
	}
}

type HomeFacadeI interface {
	WatchMovie()
	EndMovie()
}

type HomeFacade struct {
	MI MovieI
	LG LightI
}

func (h HomeFacade) WatchMovie() {
	h.MI.MovieOn()
	h.LG.LightOff()
}

func (h HomeFacade) EndMovie() {
	h.MI.MovieOff()
	h.LG.LightOn()
}

func NewHomeFacade(mi MovieI, lg LightI) HomeFacade {
	return HomeFacade{
		MI: mi,
		LG: lg,
	}
}

// MovieI 电影接口
type MovieI interface {
	MovieOn()
	MovieOff()
}

type MovieImp struct {}

func NewMovieImp() MovieImp {
	return MovieImp{}
}

func (m MovieImp) MovieOn() {
	fmt.Println("turn on movie")
}

func (m MovieImp) MovieOff() {
	fmt.Println("turn off movie")
}

// LightI 灯接口
type LightI interface {
	LightOn()
	LightOff()
}

type LightImp struct {}

func NewLightImp() LightImp {
	return LightImp{}
}

func (l LightImp) LightOn() {
	fmt.Println("turn on light")
}

func (l LightImp) LightOff() {
	fmt.Println("turn off light")
}