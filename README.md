BLAKESUM
========

Utility to calculate BLAKE-224, -256, -384, -512 checksums.


Installation
------------

From source, if you have Go installed:

	$ go get github.com/dchest/blakesum


Usage
-----

	blakesum [-a=224|256|384|512] [filename1] [filename2] ...

By default calculates BLAKE-256 sum. Pass option "-a=xxx" before filenames to
calculate BLAKE-xxx, where xxx is 224, 256, 384, or 512.  If no filenames
specified, reads from stdin.


Examples
--------

	$ echo -n "Hello world" | blakesum
	7ad560fefa2d287892478dccc5c724694fe21a2f8b004486cc87f76c40618575

	$ echo -n "Hello world" | blakesum -a=224
	fde00425968221a451c6f06f008bddc44cfdb9d8190507d0fd063707

	$ blakesum /bin/sh /etc/bashrc
	BLAKE-256 (/bin/sh) = b162509ca9f0920d41cec68cb0a2d4037b8d0a89463d888f8e971be085da6786
	BLAKE-256 (/etc/bashrc) = 1cf1a48d8c7bbd8410f270208bb60276b063fcf23608f6c3198b26391cff4e1e

