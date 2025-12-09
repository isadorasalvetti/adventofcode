    #include <map>
    #include <cstdlib>
    #include <iostream>
    #include <fstream>
    #include <string>
    #include <vector>
    #include <cassert>
    #include <regex>
    #include <cstdint>

    using namespace std;

    struct Point {
        int x;
        int y;

        bool operator<(const Point& other) const {
            return (x < other.x) || (x == other.x && y < other.y);
        }
    };

    int neighboors(Point point, map<Point, char> &room){
        int count = 0;
        for (int y = -1; y <= 1; y ++){
            for (int x = -1; x <= 1; x ++){
                if (y==0 && x==0) continue;
                if (room[{point.x + x, point.y + y}] == '@') count ++;
            }
        }
        return count;
    }

    int solvep1(map<Point, char> room, int maxx, int maxy){
        //cout << neighboors({1, 2}, room) << endl;
        int count = 0;
        for (int y = 0; y <= maxy; y ++){
            for (int x = 0; x <= maxx; x ++){
                if (room[{x, y}] == '@' && neighboors({x, y}, room) < 4) count++;
            }
        }
        
        return count;
    }

    int removeRolls(map<Point, char> &room, int maxx, int maxy){
        int removed = 0;
        for (int y = 0; y <= maxy; y ++){
            for (int x = 0; x <= maxx; x ++){
                if (room[{x, y}] == '@' && neighboors({x, y}, room) < 4) {
                    room[{x, y}] = '.';
                    removed++;
                }
            }
        }
        return removed;
    }

    int solvep2(map<Point, char> &room, int maxx, int maxy){
        //cout << neighboors({1, 2}, room) << endl;
        int total = 0;
        int removed = 999;
        while (removed != 0) {
            removed = removeRolls(room, maxx, maxy);
            cout << "Removed: " << removed << endl;
            total += removed;
        }
        
        return total;
    }


    int main() {
        ifstream input;
        cout << "Day 4:" << "\n";

        input.open("input.txt");
        assert(input.is_open());

        map<Point, char> room;

        int y = 0;
        int x = 0;
        for (string line; getline( input, line ); ){
            x = 0;
            for (char &paper: line) {
                room[{x, y}] = paper;
                x ++;
            }
            y++;
        }
        input.close();

        cout << "parsed room " << x << " by " << y << endl;
        cout << solvep2(room, x, y) << endl;

        return 0;
    }