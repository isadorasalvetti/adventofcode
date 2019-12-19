password_min = 245182
password_max = 790572


def get_amout_of_passwords2():
    passwords = 0
    for i in range(245182, 790572):
        if not find_adjacent_num(i) or not check_digits(i):
            continue
        else:
            passwords += 1
    return passwords


def get_amout_of_passwords():
    passwords = 0
    for i in range(2, 8):
        for j in range(3, 10):
            if j >= i:
                for k in range(3, 10):
                    if k >= j:
                        for l in range(3, 10):
                            if l >= k:
                                for m in range(3, 10):
                                    if m >= l:
                                        for n in range(3, 10):
                                            if n >= m:
                                                num = 100000*i + 10000*j + 1000*k + 100*l + 10*m + n
                                                if password_min < num < password_max:
                                                    if find_adjacent([i, j, k, l, m, n]):
                                                        passwords += 1
    return passwords


def find_adjacent(list):
    for i in range(len(list)-1):
        if list[i] == list[i+1] and \
                (i - 1 < 0 or not list[i] == list[i - 1]) and \
                (i + 2 >= len(list) or not list[i] == list[i+2]):
            try: print(list, list[i], list[i+1], list[i-1], list[i+2])
            except IndexError: pass

            return True
    return False


def find_adjacent_num(num):
    num_str = str(num)
    for i in range(len(num_str)-1):
        if num_str[i] == num_str[i+1]: return True
    return False


def check_digits(num):
    num_str = str(num)
    for i in range(len(num_str) - 1):
        if int(num_str[i]) > int(num_str[i+1]): return False
    return True


#p1 = get_amout_of_passwords2()
p2 = get_amout_of_passwords()
print(p2)
