#include <cstdlib>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <cassert>
#include <regex>

using namespace std;


long long solvep1(vector<pair<long long, long long>> &ids) {
    long long bad_ids = 0;
    for (auto &range: ids) {
        cout << range.first << "-" << range.second << endl;
        for (long long i = range.first; i <= range.second; i++){
            string id_string = to_string(i);
            int pos_repeat = 1;
            if (id_string.length() % 2 == 1) { continue; }
            while (pos_repeat <= id_string.length() / 2) {
                regex re_check("^(" + id_string.substr(0, pos_repeat) + "){2}$");
                if (regex_match(id_string, re_check)) {
                    cout << "Match! " << i << " " << pos_repeat << " " << bad_ids << endl;
                    bad_ids += i;
                    break;
                }
                pos_repeat ++;
            }
            
        }
    }
    return bad_ids;
}


long long solvep2(vector<pair<long long, long long>> &ids) {
    long long bad_ids = 0;
    for (auto &range: ids) {
        cout << range.first << "-" << range.second << endl;
        for (long long i = range.first; i <= range.second; i++){
            string id_string = to_string(i);
            int pos_repeat = 1;
            if (id_string.length() % pos_repeat != 0) { continue; }
            while (pos_repeat <= id_string.length() / 2) {
                regex re_check("^(" + id_string.substr(0, pos_repeat) + "){2,}$");
                if (regex_match(id_string, re_check)) {
                    cout << "Match! " << i << " " << pos_repeat << " " << bad_ids << endl;
                    bad_ids += i;
                    break;
                }
                pos_repeat ++;
            }
            
        }
    }
    return bad_ids;
}

int main() {
    ifstream input;
    cout << "Day 2:" << "\n";

    input.open("input.txt");
    assert(input.is_open());

    vector<pair<long long, long long>> moves;

    long long id1;
    long long id2;
    char dash;
    char comma;
    while(input >> id1){
        input >> dash;
        input >> id2;
        input >> comma;  
        assert(comma == ',' && dash == '-');
        moves.push_back(make_pair(id1, id2));
    }
    input.close();

    cout << solvep2(moves);

    return 0;
}