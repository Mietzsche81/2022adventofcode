# Import packages
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
    with open(file=fileName, mode='r') as fin:
        data = fin.read()
    
    # Now that we have the data, output that data
    return data

def process_data(data):
    # Do stuff here
    output = data

    return output

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
    
    """
    l = [1,2,3,4,5,'a',"adsfasd",1.234234]
    l.append("4")
    t = (1,2,3)
    d = {"asdf": 0, "COCO": "butt"}
    d["COCO"] # vbutt
    """
    


# ---END MAIN FUNCTION HERE---
