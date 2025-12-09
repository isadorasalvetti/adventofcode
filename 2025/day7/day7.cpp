#include <map>
#include <set>
#include <cstdlib>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <cassert>

using namespace std;

struct Point {
    int x;
    int y;

    bool operator<(const Point& other) const {
        return (x < other.x) || (x == other.x && y < other.y);
    }
};

void print(int currY, set<int> activeX) {
    cout << currY << ": ";
    for (auto &currX: activeX){
        cout << currX << " ";
    }
    cout << endl;
}

int solvep1(Point start, set<Point> splitter, int bottom) {
    int splits = 0;
    int currY = start.y;
    set<int> activeX = {start.x};
    while (currY < bottom) {
        currY ++;
        set<int> newX; 
        for (auto &currX: activeX){
            if (splitter.find({currX, currY}) != splitter.end()) {
                newX.insert(currX+1);
                newX.insert(currX-1);
                splits ++;
            } else newX.insert(currX);
        }

        activeX = newX;
    }

    return splits;
}


void printv(vector<int> v){
    cout << "v: ";
    for (auto el: v){
        cout << el << ", "; 
    }
    cout << endl;
}

long long solvep2(Point start, set<Point> splitter, int bottom) {
    int currY = start.y;
    map<int, long long> timeline_heads = {{start.x, 1}};

    while (currY < bottom) {
        currY ++;
        map<int, long long> new_heads;        
        for (auto &timeline_head: timeline_heads){
            if (splitter.find({timeline_head.first, currY}) != splitter.end()) {
                new_heads[timeline_head.first+1] += timeline_head.second;
                new_heads[timeline_head.first-1] += timeline_head.second;            
            } else {
                new_heads[timeline_head.first] += timeline_head.second; 
            };
        }
        timeline_heads = new_heads;
    }

    long long acc = 0;
    for (auto &timeline_head: timeline_heads){
        acc += timeline_head.second;
    }
    return acc;
}


int main() {
    cout << "Day 7:" << "\n";
    ifstream input;

    input.open("input.txt");
    assert(input.is_open());

    set<Point> splitter;

    int x;
    int y = 0;
    Point start;
    for (string line; getline( input, line ); ){
        x = 0;
        for (char &pos: line) {
            if (pos == '^') splitter.insert({x, y});
            if (pos == 'S') start = {x, y};
            x ++;
        }
        y++;
    }

    cout << solvep2(start, splitter, y) << endl;

    return 0;
}