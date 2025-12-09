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
    int x;
    int y;
    int z;

    bool operator<(const Point& other) const {
        return (x < other.x) || (x == other.x && y < other.y) || (x == other.x && y == other.y && z < other.z);
    }

    bool operator==(const Point& other) const{
        return (x == other.x && y == other.y && z==other.z);
    }
};

struct Dist {
    Point p1;
    Point p2;
    long long dist;
    
    bool operator<(const Dist& other) const {
        return (dist < other.dist);
    }
};

long long square(long long x){
    return x*x;
}

void printdist(Dist d){
    cout << d.dist << ": " << d.p1.x << ", " << d.p1.y << ", " << d.p1.z << "-" << d.p2.x << ", " << d.p2.y << ", " << d.p2.z << endl;

}

void printpoint(Point p1, int c){
    cout << p1.x << ", " << p1.y << ", " << p1.z << " - " << c << endl;
}


int solvep1(vector<Point> boxes, int times){
    int m_circuit = 0;
    map<Point, int> connections;
    vector<Dist> distances;

    for (int i=0; i<boxes.size(); i++){
        for (int j=i+1; j<boxes.size(); j++) {
            Point p1 = boxes[i]; Point p2 = boxes[j];
            long long dist = square(p1.x - p2.x) + square(p1.y - p2.y) + square(p1.z - p2.z);
            distances.push_back({p1, p2, dist});
        }
    }

    sort(distances.begin(), distances.end());
    cout << boxes.size() << ", " << distances.size() << endl;
    
    cout << distances[0].dist << endl;
    for (int i=0; i<times; i++){
        //printdist(distances[i]);
        
        Point p1 = distances[i].p1;
        Point p2 = distances[i].p2;

        int c1 = connections[p1];
        int c2 = connections[p2];

        if (c1 && c2) {
            vector<Point> change;
            for (auto const& [point, c] : connections) if (c == c2) change.push_back(point);
            for (auto point: change) connections[point] = c1;
        }
        else if (c1) connections[p2] = c1;
        else if (c2) connections[p1] = c2;
        else {
            m_circuit ++;
            connections[p1] = m_circuit;
            connections[p2] = m_circuit;
        }
    }

    for (auto const& [point, c] : connections){
        //printpoint(point, c);
    }
    
    vector<vector<Point>> circuits;
    vector<int> c_sizes;
    circuits.resize(m_circuit);
    for (auto const& [point, c] : connections){
        circuits[c-1].push_back(point);
    }
    
    for(auto v: circuits) c_sizes.push_back(v.size());

    sort(c_sizes.begin(), c_sizes.end(), greater<int>());

    return c_sizes[0] * c_sizes[1] * c_sizes[2];
}

int solvep2(vector<Point> boxes){
    int m_circuit = 0;
    map<Point, int> connections;
    vector<Dist> distances;

    for (int i=0; i<boxes.size(); i++){
        for (int j=i+1; j<boxes.size(); j++) {
            Point p1 = boxes[i]; Point p2 = boxes[j];
            long long dist = square(p1.x - p2.x) + square(p1.y - p2.y) + square(p1.z - p2.z);
            distances.push_back({p1, p2, dist});
        }
    }

    sort(distances.begin(), distances.end());
    cout << boxes.size() << ", " << distances.size() << endl;
    
    Point p1;
    Point p2;
    for (int i=0; i<distances.size(); i++){        
        p1 = distances[i].p1;
        p2 = distances[i].p2;

        int c1 = connections[p1];
        int c2 = connections[p2];

        if (c1 && c2) {
            vector<Point> change;
            for (auto const& [point, c] : connections) if (c == c2) change.push_back(point);
            for (auto point: change) connections[point] = c1;
        }
        else if (c1) connections[p2] = c1;
        else if (c2) connections[p1] = c2;
        else {
            m_circuit ++;
            connections[p1] = m_circuit;
            connections[p2] = m_circuit;
        }

        if (connections.size() == boxes.size()){
            int mc=connections[p1];
            bool unique = true;
            for (auto const& [point, c] : connections){
                if (c != mc) {
                    unique = false;
                    break;
                }
            }
            if (unique) {
                break;
            }
        }

    }

    for (auto const& [point, c] : connections){
        //printpoint(point, c);
    }

    printpoint(p1, 0);
    printpoint(p2, 0);
    
    return p1.x * p2.x;
}

int main() {
    cout << "Day 8:" << "\n";
    ifstream input;

    input.open("input.txt");
    assert(input.is_open());

    vector<Point> boxes;

    int x;
    int y;
    int z;
    char comma;
    while(input >> x){
        input >> comma;
        input >> y;
        input >> comma;
        input >> z;
        boxes.push_back({x, y, z});
    }

    cout << solvep2(boxes);

    return 0;
}