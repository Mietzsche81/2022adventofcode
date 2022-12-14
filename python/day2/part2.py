# Import packages
import csv
import os
import pathlib
import sys

# Utility functions
def parse_input(fileName: str = ""):
    # Make sure we have the right file
    if fileName.isspace():
        fileName = os.path.join(
            pathlib.Path(__file__).parent.absolute(),
            'puzzle.input',
        )

    # parse data from that file
    data = []
    with open(file=fileName, mode='r') as fin:
        csvin = csv.reader(fin, delimiter=" ")

        for row in csvin:
            data.append(row)

    # Now that we have the data, output that data
    return data

def process_data(data):
    win_score = {
        "X": 0,
        "Y": 3,
        "Z": 6,
    }
    choice_score = {
        "AX": 3,
        "BX": 1,
        "CX": 2,
        "AY": 1,
        "BY": 2,
        "CY": 3,
        "AZ": 2,
        "BZ": 3,
        "CZ": 1,
    }

    # Start running sum
    total_score = 0
    # For each row in our data
    for row in data:
        # Add the choice score
        total_score += choice_score[row[0]+row[1]]
        # Add the outcome score
        total_score += win_score[row[1]]
        

    return total_score

def display_output(output):
    
    # display logic here
    print(output)

    return

# ---Main function: THIS IS WHAT GETS EXECUTED---
if __name__ == "__main__":
    # Read input
    fileName = sys.argv[1]
    data = parse_input(fileName)

    # Process Data
    output = process_data(data)

    # Report
    display_output(output)
    
# ---END MAIN FUNCTION HERE---