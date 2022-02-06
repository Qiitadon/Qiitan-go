<?php

function fib($x){
	if ($x === 0) {
		return 0;
	}

	if ($x === 1) {
		return 1;
	}

	return fib($x-1) + fib($x-2);
}

echo fib(35) . PHP_EOL;

// Output: 9227465
