def format_input(file_name):
    file = file_name + ".txt"
    open_file = open(file)
    text = open_file.read().split('\n')
    num_input = []
    for t in text:
        try: num_input.append(int(t))
        except ValueError: pass
    return num_input

