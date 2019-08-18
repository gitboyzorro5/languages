num1 = float(input("Enter a number"))
op = input("Enter an operation ")
num2 = float(input("Enter the second number"))

if op == "+":
	print(num1 + num2)
elif op == "-":
	print(num1 - num2)
elif op == "/":
	print(num1 / num2)
elif op == "*":
	print(num1 * num2)
else:
	print("invalid operator")