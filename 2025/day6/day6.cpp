#include <map>
#include <cstdlib>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <cassert>
#include <regex>
#include <cstdint>
#include <numeric>

using namespace std;

long long doHomework(vector<vector<int>>problems, vector<char>ops){
    long long total = 0;
    long long sol = 0;
    int idx = 0;
    for (vector<int> problem: problems){
        if (ops[idx] == '+') sol = std::reduce(problem.begin(), problem.end());
        else sol = reduce(problem.begin(), problem.end(), (long long)1, multiplies<long long>());
        total += sol;
        idx ++;
    }

    return total;
}

long long solvep1(string input_file) {
    ifstream input;

    input.open(input_file);
    assert(input.is_open());

    int num;
    char op;
    vector<char> ops;
    vector<int> allnums;
    while(input >> num){
        allnums.push_back(num);
    }
    input.clear();
    while(input >> op){
        ops.push_back(op);
    }
    input.close();

    cout << "parsed nums " << allnums.size() << " ops " << ops.size() << endl;

    int problems_cout = ops.size();
    vector<vector<int>> problems;
    problems.resize(problems_cout);
    for (int i=0; i < allnums.size(); i++) {
        problems[i%ops.size()].push_back(allnums[i]);
    }    

    cout << "Problem 3: ";
    for (int num: problems[2]) {
            cout << num << " ";
    }
    cout << endl;

    return doHomework(problems, ops);
}

long long solvep2(string input_file){
    ifstream input;

    input.open(input_file);
    assert(input.is_open());

    vector<string> lines;
    vector<string> trans_lines;

    for (string line; getline( input, line ); ){
        lines.push_back(line);
    }

    string ops_line = lines[lines.size()-1];
    lines.erase(lines.end()-1);

    trans_lines.resize(lines[0].size());

    for (int i = 0; i < lines.size(); i++){
        for (int j = 0; j < lines[i].size(); j++){
            trans_lines[j].push_back(lines[i][j]);
        }
    }

    vector<vector<int>> problems;
    problems.resize(1);
    int indx = 0;
    for (string line: trans_lines){
        try {
            int num = stoi(line);
            problems[indx].push_back(num);
        }
        catch (const std::invalid_argument& e){
            problems.push_back({});
            indx ++;
        }
    }

    vector<char> ops;
    for (char el: ops_line){
        if (el!=' ') ops.push_back(el);
    }

    return doHomework(problems, ops);
}

int main() {
    cout << "Day 6:" << "\n";
    string input_file = "input.txt"; 
    cout << solvep2(input_file);
    return 0;
}