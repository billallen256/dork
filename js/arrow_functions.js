'use strict'

function adder(x) {
	let f = (p,n) => p+n
	return x.reduce(f)
}
