# DICE SIMULATOR
# --------------
# Alex Safatli
# June 2013  
# --------------

import sys, re
import random as r
from matplotlib import pyplot as plt
from matplotlib.backends.backend_pdf import PdfPages


class Die:

    def __init__(self, num: int, sides: int):
        self.num = num
        self.sides = sides

    def roll(self):
        result = 0
        for i in range(self.num):
            result += r.randint(1, self.sides)
        return result


class DiceRoller:

    def __init__(self, inputstr):
        self.string = inputstr
        self.dice = []
        self.modifier = 0
        self.__parse__()

    def __parse__(self):
        dic = re.compile('\d*d\d*')
        posmod = re.compile('\+\d*')
        negmod = re.compile('-\d*')
        for d in dic.finditer(self.string):
            d = d.group(0).split('d')
            self.dice.append(Die(int(d[0]), int(d[1])))
        for _ in posmod.finditer(self.string):
            d = d.group(0).strip('+')
            self.modifier += int(d)
        for _ in negmod.finditer(self.string):
            d = d.group(0).strip('-')
            self.modifier -= int(d)

    def roll_all(self, num):
        out = []
        for x in range(num):
            tot = 0
            for di in self.dice:
                tot += di.roll()
            tot += self.modifier
            out.append(tot)
        return out


# Get input.
inp = sys.argv
if len(inp) != 4:
    print('usage: python diceSimulator.py XdY+Z NUM_SAMPLES OUTPUT_NAME')
roll = inp[1]
num_times = int(inp[2])
print('Performing %s multiple times (number samples: %d).' % (roll, num_times))
outname = inp[3]

# Roll (like an 18-wheela').
d = DiceRoller(roll)
arr = d.roll_all(num_times)

# Plot.
plt.hist(arr)
plt.title(roll)
plt.ylabel('Number of Occurences per %d Samples' % num_times)
pp = PdfPages(outname)
plt.savefig(pp, format='pdf')
pp.close()
