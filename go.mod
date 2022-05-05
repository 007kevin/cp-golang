module cp

go 1.17

replace lib/io => ./lib/io

replace lib/math => ./lib/math

require lib/io v0.0.0-00010101000000-000000000000

require lib/math v0.0.0-00010101000000-000000000000

require (
	github.com/ktateish/gottani v0.0.0-20201015003731-5e954ca908ae // indirect
	golang.org/x/tools v0.0.0-20200925191224-5d1fdd8fa346 // indirect
)
