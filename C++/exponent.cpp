//a pow fucntion to return the exponent

#include <iostream>
//#include <cmath>
using namespace std;

double pow(int,int); // this is a fucntion declaration


int main(){

	cout << "The value of 2^5 is :" << pow(2,5);

	return 0;
}

 double pow(int base,int exponent) // this is a function definition
{

	int result = 1;
	int i;

	for (i = 1; i <= exponent; i++) {

		result = result * base;
	}

return result;

}
