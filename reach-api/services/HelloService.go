package services

import "rsc.io/quote"

func HelloService() string {
	return quote.Hello()
}