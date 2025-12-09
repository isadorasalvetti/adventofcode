#include <cstdlib>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <algorithm>
#include <cassert>
#include <regex>
#include <cstdint>

using namespace std;

int solvep1(vector<pair<long long, long long>> &ranges, vector<long long> &foods){
    int good_food = 0;
    for (auto food: foods){
        for (auto range: ranges){
            if (food >= range.first && food <= range.second){
                good_food ++;
                break;
            }
        }
    }
    return good_food;
}

long long solvep2(vector<pair<long long, long long>> ranges){    
    sort(ranges.begin(), ranges.end());

    bool pass_again = true;
    while (pass_again){
        vector<pair<long long, long long>> newRanges;
        pass_again = false;

        for (auto range:ranges) {
            bool consumed = false;
            for (auto &newRange: newRanges) {
                if (!(range.second < newRange.first || range.first > newRange.second)) {
                    pass_again = true;
                    //cout << "Consuming " << range.first << "," << range.second << endl;
                    if (range.first <= newRange.second && range.second > newRange.second) {
                        newRange.second = range.second;
                        //cout << "case 1 " << newRange.first << "," << newRange.second << endl;
                    } else if (range.second >= newRange.first && range.first < newRange.first) {
                        newRange.first = range.first;
                        //cout << "case 2 " << endl;
                    }
                    consumed = true;
                    break;
                }
            }
            if (!consumed) newRanges.push_back(range);
        }

        ranges = newRanges;
    }
    

    long long total = 0;
    sort(ranges.begin(), ranges.end());
    for(auto range: ranges){
        total += range.second - range.first + 1;
        cout << range.first << ", " << range.second << ": " << range.second - range.first + 1 << endl;
    }

    return total;

}

int main() {
    ifstream input;
    cout << "Day 5:" << "\n";

    input.open("input.txt");
    assert(input.is_open());

    vector<long long> foods;
    vector<pair<long long, long long>> ranges;

    long long num1;
    long long num2;
    while(input >> num1) {
        input >> num2;
        if (num2 < 0) {
            ranges.push_back({num1, -num2});
        } else {
            foods.push_back(num1);
            if (num2) foods.push_back(num2);
        }
    }
    input.close();

    cout << "Foods: " << foods.size() << ", Ranges: " << ranges.size() << endl;
    //cout << solvep1(ranges, foods) << endl;
    cout << solvep2(ranges) << endl;

    return 0;
}