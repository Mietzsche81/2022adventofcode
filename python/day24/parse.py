def parse(input_file):
    with open(input_file, "r") as fin:
        lines = fin.readlines()

    return [stripped for line in lines if (stripped := line.strip())]
