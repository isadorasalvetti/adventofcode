import sys

from timeit import default_timer as timer
from pprint import pprint
from collections import namedtuple

DIRS = {'<': (-1,  0),
        '>': ( 1,  0),
        '^': ( 0, -1),
        'v': ( 0,  1),
       }

def show(walls, boxes, robot, size):
    for y in range(size[1]):
        for x in range(size[0]):
            pos = (x, y)
            if pos in walls:
                print('#', end='')
            elif pos in boxes:
                print('O', end='')
            elif pos == robot:
                print('@', end='')
            else:
                print('.', end='')
        print()
    print()

def show_pt2(walls, boxes, robot, size):
    for y in range(size[1]):
        x = 0
        while x < size[0]:
            pos = (x, y)
            if pos in walls:
                print('#', end='')
                x += 1
            elif (x, x + 1, y) in boxes:
                print('[]', end='')
                x += 2
            elif pos == robot:
                print('@', end='')
                x += 1
            else:
                print('.', end='')
                x += 1
        print()
    print(flush=True)

def part1(filename):
    walls = set()
    boxes = set()
    robot = None
    with open(filename, 'r') as fp:
        field, directions = fp.read().split('\n\n')
        for y, line in enumerate(field.split('\n')):
            for x, char in enumerate(line):
                if char == '#':
                    walls.add((x, y))
                elif char == 'O':
                    boxes.add((x, y))
                elif char == '@':
                    robot = (x, y)
        size = (x + 1, y + 1)
        assert robot is not None, "Didn't find robot"
        moves = directions.replace('\n', '')
    
    # Basically Sokoban
    for move in moves:
        d = DIRS[move]
        pos = robot[0] + d[0], robot[1] + d[1]
        
        if pos in boxes:
            involved = []
            while pos in boxes:
                involved.append(pos)
                pos = pos[0] + d[0], pos[1] + d[1]
            if pos not in walls:    # Pushed to empty space
                robot = robot[0] + d[0], robot[1] + d[1]
                for ix, box in enumerate(involved):
                    boxes.remove(box)
                    involved[ix] = box[0] + d[0], box[1] + d[1]
                boxes.update(involved)
        elif pos not in walls:
            robot = pos
        
    answer = 0
    for box in boxes:
        answer += (100 * box[1]) + box[0]
    
    print(f"Part 1: {answer}")
    return answer

Box = namedtuple('Box', 'min_x max_x y')

def part2(filename):
    verbose = False
    walls = set()
    boxes = set()
    robot = None
    with open(filename, 'r') as fp:
        field, directions = fp.read().split('\n\n')
        for y, line in enumerate(field.split('\n')):
            for x, char in enumerate(line):
                xx = x * 2
                if char == '#':
                    walls.add((xx, y))
                    walls.add((xx + 1, y))
                elif char == 'O':
                    boxes.add(Box(xx, xx + 1, y))
                elif char == '@':
                    robot = (xx, y)
        size = (xx + 2, y + 1)
        assert robot is not None, "Didn't find robot"
        moves = directions.replace('\n', '')

    for ix, move in enumerate(moves):
        #verbose = ix >= 8162
        d = DIRS[move]
        pos = robot[0] + d[0], robot[1] + d[1]
        if verbose:
            print(f"Move {move}: {ix}")
            print(f"{robot=}, {pos=}, {pos not in walls=}")
            print("BOXES:")
            pprint(boxes)
        
        if pos not in walls:
            involved = set()
            hit_wall = False
            if move in '<>':
                if move == '>':
                    search = (pos[0], pos[0] + 1, pos[1])
                    while search in boxes:
                        involved.add(search)
                        search = (search[0] + 2, search[1] + 2, search[2])
                    hit_wall = (search[0], search[2]) in walls
                else:
                    search = (pos[0] - 1, pos[0], pos[1])
                    while search in boxes:
                        involved.add(search)
                        search = (search[0] - 2, search[1] - 2, search[2])
                    hit_wall = (search[0] + 1, search[2]) in walls
            else:   # move in '^v':
                active = None
                for initial in ((pos[0], pos[0] + 1, pos[1]), (pos[0] - 1, pos[0], pos[1])):
                    if initial in boxes:
                        active = [Box(*initial)]
                        break
                
                while active:
                    box = active.pop()
                    involved.add(box)
                    y = box[2] + d[1]
                    if ((box[0], y) in walls) or ((box[1], y) in walls):
                        hit_wall = True
                        break

                    for check in boxes:
                        if check == box:
                            continue
                        if check[2] == y:
                            if (box[0] in check[:2]) or (box[1] in check[:2]):
                                active.append(check)
            if verbose:
                print(f"{hit_wall=}")
                print("INVOLVED:")
                pprint(involved)
            if not hit_wall:
                if involved:
                    for box in involved:
                        boxes.remove(box)
                    for box in involved:
                        boxes.add(Box(box[0] + d[0], box[1] + d[0], box[2] + d[1]))
                robot = pos
        if ix > 1000 and ix < 2000:
            answer = 0
            for box in boxes:
                answer += (100 * box[2]) + box[0]
            print(f"{ix}: {answer}")
            
    answer = 0
    for box in boxes:
        answer += (100 * box[2]) + box[0]

    print(f"Part 2: {answer}")
    print("Final Position:")
    return answer

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print(f"USAGE: {sys.argv[0]} <input file>", file=sys.stderr)
        sys.exit(1)

    before = timer()
    part1(sys.argv[1])
    after = timer()
    print(f"Part 1 Elapsed Time: {(after - before) * 1000.0:.6f} ms.")

    before = timer()
    part2(sys.argv[1])
    after = timer()
    print(f"Part 2 Elapsed Time: {(after - before) * 1000.0:.6f} ms.")

