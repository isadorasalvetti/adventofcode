#include <map>
#include <set>
#include <cstdlib>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <cassert>
#include <algorithm>

using namespace std;


struct Point {
    long long x;
    long long y;

    bool operator<(const Point& other) const {
        return (x < other.x) || (x == other.x && y < other.y);
    }

    bool operator==(const Point& other) const{
        return (x == other.x && y == other.y);
    }
};

long long solvep1(vector<Point> reds){
    long long biggest = 0;

    for (int i =0; i< reds.size(); i++){
        for (int  j=i+1; j< reds.size(); j++){
            long long size = (abs((reds[j].x - reds[i].x))+1) * (abs((reds[j].y - reds[i].y))+1);
            if (size > biggest) {
                //cout << reds[i].x << ", " << reds[i].y << " " << reds[j].x << ", " << reds[j].y << endl;
                biggest = size;
            }
        }
    }

    return biggest;
}

long long solvep2(vector<Point> reds){
    long long biggest = 0;
    map<int, vector<Point>> isGreenX;
    map<int, vector<Point>> isGreenY;

    for (int i =0; i< reds.size(); i++){
        for (int  j=i+1; j< reds.size(); j++){
            if (reds[i].x == reds[j].x) {
                int maxy = max(reds[i].y, reds[j].y);
                int miny = min(reds[i].y, reds[j].y);
                isGreenX[reds[i].x].push_back({miny, maxy});
            } 
            else {
                int maxx = max(reds[i].x, reds[j].x);
                int minx = min(reds[i].x, reds[j].x);
                isGreenY[reds[i].y].push_back({minx, maxx});
            } 
        }
    }

    for (int i =0; i< reds.size(); i++){
        for (int  j=i+1; j< reds.size(); j++){
            
            // right side
            bool notGreen = true;
            for (auto range: isGreenX[reds[i].x]) {
                int maxy = max(reds[i].y, reds[j].y);
                int miny = min(reds[i].y, reds[j].y);
                if (range.x <= miny && range.y >= maxy){
                    notGreen = false;
                    break;
                }
            }
            if (notGreen) continue;

            // left side
            notGreen = true;
            for (auto range: isGreenX[reds[j].x]) {
                int maxy = max(reds[i].y, reds[j].y);
                int miny = min(reds[i].y, reds[j].y);
                if (range.x <= miny && range.y >= maxy){
                    notGreen = false;
                    break;
                }
            }
            if (notGreen) continue;

            // top side
            notGreen = true;
            for (auto range: isGreenY[reds[i].y]) {
                int maxx = max(reds[i].x, reds[j].x);
                int minx = min(reds[i].x, reds[j].x);
                if (range.x <= minx && range.y >= maxx){
                    notGreen = false;
                    break;
                }
            }
            if (notGreen) continue;

            // bottom side
            notGreen = true;
            for (auto range: isGreenY[reds[j].y]) {
                int maxx = max(reds[i].x, reds[j].x);
                int minx = min(reds[i].x, reds[j].x);
                if (range.x <= minx && range.y >= maxx){
                    notGreen = false;
                    break;
                }
            }
            if (notGreen) continue;

            long long size = (abs((reds[j].x - reds[i].x))+1) * (abs((reds[j].y - reds[i].y))+1);
            if (size > biggest) {
                cout << reds[i].x << ", " << reds[i].y << " " << reds[j].x << ", " << reds[j].y << endl;
                biggest = size;
            }
        }
    }

    return biggest;
}

int main() {
    cout << "Day 8:" << "\n";
    ifstream input;

    input.open("sample.txt");
    assert(input.is_open());

    vector<Point> reds;

    int x;
    int y;
    char comma;
    while(input >> x){
        input >> comma;
        input >> y;
        reds.push_back({x, y});
    }

    cout << solvep2(reds);

    return 0;
}