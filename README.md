# Interview Cake - Weekly Problem #13 Solution

This is my solution in Go, not the most elegant solution possible but it works
well enough.

## Problems

1. While I am checking for overflow on either Order slice, I have to make the
   combined order slice one spot shorter so that I don't double the lowest
	 number. This could be solved by checking to see if either list is
	 exhausted, which is how the solution on the site works. This way it quits
	 when there are no longer any orders left in a slice. My solution then
	 requires tacking the last, largest order ID into the combined.
1. Setting a fixed size array and assigning to those values would be the
   smarter way of combining orders instead of using the append because
	 you end up with a 0 as the first element in the slice which then needs
	 to be trimmed off

Not the most amazing solution by far but I'm happy since I haven't touched
go in a couple months. This was a nice refresher.
