str_input = [1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,19,5,23,1,9,23,27,2,27,6,31,1,5,31,35,2,9,35,39,2,6,39,43,2,
            43,13,47,2,13,47,51,1,10,51,55,1,9,55,59,1,6,59,63,2,63,9,67,1,67,6,71,1,71,13,75,1,6,75,79,1,9,79,83,
            2,9,83,87,1,87,6,91,1,91,13,95,2,6,95,99,1,10,99,103,2,103,9,107,1,6,107,111,1,10,111,115,2,6,115,119,
            1,5,119,123,1,123,13,127,1,127,5,131,1,6,131,135,2,135,13,139,1,139,2,143,1,143,10,0,99,2,0,14,0]


def parse_input(str):
    num_lst = []
    for t in str.split(","):
        try:
            num_lst.append(int(t))
        except ValueError:
            pass
    return num_lst


def execute(inpt=None):
    if inpt == None:
        inpt = str_inpt[:]
    i = 0
    while True:
        if i > len(inpt): return

        if inpt[i] == 1:
            pos1 = inpt[i+1]
            pos2 = inpt[i+2]
            pos3 = inpt[i+3]
            sum_ = inpt[pos1] + inpt[pos2]
            inpt[pos3] = sum_
            i += 4
        elif inpt[i] == 2:
            pos1 = inpt[i + 1]
            pos2 = inpt[i + 2]
            pos3 = inpt[i + 3]
            mult = inpt[pos1] * inpt[pos2]
            inpt[pos3] = mult
            i += 4
        elif inpt[i] == 99:
            #print(inpt)
            return 0
        else:
            print("something went very wrong.", i, inpt[i])
            return -1

def find_input(output = 19690720):
    for i in range(99):
        for j in range(99):
            program = str_input[:]
            program[1] = i
            program[2] = j

            if execute(program) == -1: return
            if program[0] == output:
                print(100 * i + j)
                print(i, j)
                return

