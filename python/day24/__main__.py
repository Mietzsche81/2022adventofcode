from argparse import ArgumentParser

from parse import parse
from process import StateIterator

arg_parser = ArgumentParser()
arg_parser.add_argument("part", type=str)
arg_parser.add_argument("input", type=str)

args = arg_parser.parse_args()

data = parse(args.input)

if str(args.part).lower() in ("1", "a"):
    result = StateIterator(data).solve()
elif str(args.part).lower() in ("2", "b"):
    result = StateIterator(data).solve_backtrack()
else:
    raise ValueError(f"Unknown part {args.part}")

print(result)
