# DICE SIMULATOR
# --------------
# Alex Safatli
# June 2013  
# --------------

import sys, re
import random as r
from matplotlib import pyplot as plt
from matplotlib.backends.backend_pdf import PdfPages

class die:
    def __init__(self,num,sides):
        self.num = num
        self.sides = sides
    def roll(self):
        result = 0
        for i in xrange(self.num): result += r.randint(1,self.sides)
        return result

class diceroller:
    def __init__(self,inputstr):
        self.string = inputstr
        self.dice = []
        self.modifier = 0
        self.__parse__()
    def __parse__(self):
        dic = re.compile('[0-9]*d[0-9]*')
        posmod = re.compile('\+[0-9]*')
        negmod = re.compile('\-[0-9]*')
        for d in dic.finditer(self.string):
            d = d.group(0).split('d')
            self.dice.append(die(int(d[0]),int(d[1])))
        for mod in posmod.finditer(self.string):
            d = d.group(0).strip('+')
            self.modifier += int(d)
        for mod in negmod.finditer(self.string):
            d = d.group(0).strip('-')
            self.modifier -= int(d)
    def rollAll(self,num):
        out = []
        for x in xrange(num):
            tot = 0
            for d in self.dice: tot += d.roll()
            tot += self.modifier
            out.append(tot)
        return out

# Get input.

inp = sys.argv
if len(inp) != 4: print 'usage: python diceSimulator.py XdY+Z NUM_SAMPLES OUTPUT_NAME'
roll = inp[1]
numtimes = int(inp[2])
print 'Performing %s multiple times (number samples: %d).' % (roll,numtimes)
outname = inp[3]

# Roll (like an 18-wheela').

d = diceroller(roll)
arr = d.rollAll(numtimes)

# Plot.

plt.hist(arr)
plt.title(roll)
plt.ylabel('Number of Occurences per %d Samples' % (numtimes))
pp = PdfPages(outname)
plt.savefig(pp,format='pdf')
pp.close()
