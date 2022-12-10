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
        lines = fin.readlines()
        count = 0
        elves = [ [] ]
        for line in lines:
            count += 1
            # If reaching the end of an elf, create a new elf
            if line.isspace():
                elves.append([])
            # Otherwise, we have another entry to add to our elf
            else:
                # Convert the line to an int and add it to the current elf
                # The current elf is the last elf in our list, so use -1
                elves[-1].append(int(line))
    
    # Now that we have the data, output that data
    print(f"We read {count} lines from {fileName}")
    return elves

def process_data(data):
    # For each elf, need to sum items
    output = []
    # For each elf:
    for elf in data:
        # Start a running sum with initial value of nothing
        elf_total = 0
        # For each entry in the elfs bag:
        for entry in elf:
            # add the entry to the running sum
            elf_total += entry
        # After completing running sum, add that elf's total as
        # an entry to our output
        output.append(elf_total)        

    # Now that we have the total for each elf, 
    # we need to sort them so that we can find the total
    output = sorted(output, reverse=True)

    return output[0]

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

