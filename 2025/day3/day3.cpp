#include <cstdlib>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <cassert>
#include <regex>
#include <cstdint>

using namespace std;


int solvep1(vector<vector<int>> banks){
    int res = 0;
    for (auto bank: banks){
        int dec = 0;
        int unit = 0;

        int maxpos1 = distance(bank.begin(), max_element(bank.begin(), bank.end()));
        if (maxpos1 == bank.size()-1){
            unit = bank[maxpos1];
            dec = *max_element(bank.begin(), bank.end()-1);
        } else {
            dec = bank[maxpos1];
            unit = *max_element(bank.begin()+maxpos1+1, bank.end());
        }
        //cout << dec << unit << endl;
        res += dec * 10;
        res += unit;
    }

    return res;
}

long long solvep2(vector<vector<int>> banks){
    unsigned long long res = 0;
    for (auto bank: banks){
        int sgap = 0;
        long long batery = 0;

        for (int i = 11; i >= 0; i--) {
            int maxpos = distance(bank.begin(), max_element(bank.begin()+sgap, bank.end()-i));
            sgap = maxpos+1;
            batery += bank[maxpos] * pow(10,i);
        }
        //cout << " " << batery << " " << res << endl; 
        res += batery;
    }

    return res;
}

int main() {
    ifstream input;
    cout << "Day 3:" << "\n";

    input.open("input.txt");
    assert(input.is_open());

    vector<vector<int>> banks;

    for (string line; getline( input, line ); ){
        vector<int> nums;
        for (char &num: line) {
            nums.push_back(num - '0');
        }
        banks.push_back(nums);
    }
    input.close();

    cout << solvep2(banks) << endl;

    return 0;
}