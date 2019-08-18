for index in range(0,101):
	if index % 3 == 0 :
		print( str(index) + " " +  "fizz")
	if index % 5 == 0:
		print( str(index) + " " +  "buzz")
	if index % 5 == 0 and index % 3 == 0:
		print( str(index) + " " +  "fizzbuzz")