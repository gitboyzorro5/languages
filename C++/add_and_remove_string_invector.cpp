#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

void add_players(std::vector<std::string>& p,std::string player_name)
{

    p.push_back(player_name);
}

void print_players(std::vector<std::string> p)
{
    for(size_t i = 0;i < p.size(); i++)
    {
        std::cout << p[i] << std::endl;
    }
}

void remove_player(std::vector<std::string>& p,std::string player_name)
{

    auto itr = std::find(p.begin(),p.end(),player_name);
    if(itr != p.end())
       {
           p.erase(itr);

       }

}




int main()
{

std::vector<std::string> players = {"Marcus","martial"};

add_players(players,"mctomnay");

print_players(players);

std::cout << std::endl;

remove_player(players,"Marcus");

print_players(players);

return 0;



}
