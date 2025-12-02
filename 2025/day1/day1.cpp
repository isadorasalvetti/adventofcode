#include <cstdlib>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <cassert>


using namespace std;


int solvep1(vector<pair<char, int>> &moves) {
    int pos = 50;
    int password = 0;
    for (auto &move: moves) {
        if (move.first == 'R') {
            pos += move.second;
        } else {
            pos -= move.second;
        }
        pos = remainder(pos, 100);
        if (pos < 0) {
            pos = 100 + pos;
        }

        if (pos == 0) { password ++; }
    }

    return password;
}

int solvep2(vector<pair<char, int>> &moves) {
    int pos = 50;
    int password = 0;
    int rounds = 0;

    for (auto &move: moves) {
        int prev_pos = pos;
        if (move.first == 'R') {
            pos += move.second;
        } else {
            pos -= move.second;
        }
        
        rounds = abs(pos / 100);
        pos = pos % 100;
        
        if (pos < 0) {
            pos = 100 + pos;
            if (prev_pos != 0 ) rounds ++;
        }
        if (pos == 0 && move.first == 'L') rounds ++;

        password += rounds;
        cout << move.first << move.second << " " << pos << ", " << rounds << "\n";
    }

    if (pos == 0) { password += 1; }

    return password;
}

int main() {
    ifstream input;
    cout << "Day 1:" << "\n";

    input.open("input.txt");
    assert(input.is_open());

    vector<pair<char, int>> moves;

    int x;
    char dir;
    string line;
    while(input >> dir){
        input >> x;
        moves.push_back(make_pair(dir, x));
    }
    input.close();

    cout << solvep2(moves);

    return 0;
}