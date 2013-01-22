/*

img_date renames files based on their embeded date information.

The new path will use the same dir and extension as the old path, but the file
name will be replaced by a date representation. A suffix might be added to avoid
duplicate file names.

It performs a basic sanity check on the date, which can be ignored using the
"-insane" flag.

Installation:

	$ go get github.com/mewkiz/cmd/img_date

Usage:

	img_date [OPTION] [FILE]...

The flags are:

	-f (default=false)
		Force rename.
	-insane (default=false)
		Disable sanity checks.

Examples:

1. Force rename.

	$ img_date -f IMG_2818.JPG IMG_2819.JPG

2. Print rename shell script.

	$ img_date IMG_2818.JPG IMG_2819.JPG
	// output:
	//    mv "IMG_2818.JPG" "2012-12-21 00:23:16.jpg"
	//    mv "IMG_2819.JPG" "2012-12-21 00:23:50.jpg"

*/
package documentation
