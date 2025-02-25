import os

def moveZeroesToLeft(inputList):
    zeroesList = []
    nonZeroesList = []
    for element in inputList:
        if element == 0:
            zeroesList.append(element)
        else:
            nonZeroesList.append(element)

    sortedList = zeroesList + nonZeroesList
    return sortedList

def main():
    inputListTest1 = [1, 0, 2, 3, 0, 4, 5]
    outputList = moveZeroesToLeft(inputListTest1)
    print(f"inputListTest1 = {inputListTest1}")
    print(f"outputList = {outputList}") 

if __name__ == "__main__":
    main()