#include <iostream>
#include <vector>

void print_vector(std::vector<int> v)
{

for (int i = 0; i < v.size(); i ++)
{

std::cout << v[i] << "\t";

}

}

int sum_vector(std::vector<int> v)
{

int sum = 0;

for (int i = 0; i < v.size(); i ++)
{

sum += v[i];

}
return sum;
}

int main(){

std::vector<int> data = {1,2,3,234,434,534,54,54,54,};

print_vector(data);
std::cout << "The sum of the vector is :" << sum_vector(data) << std::endl;

return 0;

}
