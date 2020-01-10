#! /opt/local/bin/python

# gmTracker.py
# -------------------------
# Winter 2013; Alex Safatli
# -------------------------
# Simple interface to track
# health and other variables
# for units.

import random as r

# Hit Table

hitTable = {3 : 'ear', \
            4 : 'eye', \
            5 : 'neck', \
            6 : 'head', \
            7 : 'leg', \
            8 : 'chest', \
            9 : 'chest', \
            10 : 'chest', \
            11 : 'arm', \
            12 : 'arm', \
            13 : 'leg', \
            14 : 'hand', \
            15 : 'hand', \
            16 : 'head', \
            17 : 'vitals', \
            18 : 'heart'}

# Data Structure

values = {}

# Dice Structure

class dice():
    def __init__(self, num, sides):
        self.num = num
        self.sides = sides
    def roll(self):
        result = 0
        for i in xrange(self.num):
            result += r.randint(1,self.sides)
        return result

# Functions

def define(v):
    if v in values:
        return False
    values[v] = {}
    return True

def setv(v,k,data):
    if v not in values:
        return False
    values[v][k] = data
    return True

def toInt(v,k):
    if v not in values or k not in values[v]:
        return False
    values[v][k] = int(values[v][k])
    return True

def getdict(v):
    if v not in values:
        return None
    return values[v]

def get(v,k):
    if v not in values or \
       k not in values[v]:
        return None
    return values[v][k]

def decr(v,k,data):
    try:
        toInt(v,k)
        values[v][k] -= data
        return True
    except:
        return False

def add(v,k,data):
    try:
        toInt(v,k)
        values[v][k] += data
        return True
    except:
        return False

# Main

inp = ""

while (inp != "!exit"):
    inp = raw_input(">> ")
    sp = inp.split()
    cmd, args = sp[0], sp[1:]
    if (cmd == "!define" or cmd == "!def"):
        for arg in args:
            define(arg)
    elif (cmd == "!set"):
        if len(args) != 3:
            print 'Not enough or too many arguments.'
        else:
            t = setv(args[0],args[1],args[2])
            if not t:
                print 'Value %s not yet set.' % (args[0])
    elif (cmd == "!int"):
        if len(args) != 2:
            print 'Not enough or too many arguments.'
        else:
            t = toInt(args[0],args[1])
            if not t:
                print 'Value %s not yet set.' % \
                      (args[0])
    elif (cmd == "!get"):
        if len(args) < 0 or len(args) > 2:
            print 'Not enough or too many arguments.'
        elif len(args) == 0:
            print '%s' % (values)
        elif len(args) == 1:
            t = getdict(args[0])
            print '%s = %s' % (args[0],t)
        else:
            t = get(args[0],args[1])
            print '%s.%s = %s' % (args[0],args[1],t)
    elif (cmd == "!decr" or cmd == "!add"):
        if len(args) != 3:
            print 'Not enough or too many arguments.'
        else:
            if cmd == "!decr":
                t = decr(args[0],args[1],int(args[2]))
            else:
                t = add(args[0],args[1],int(args[2]))
            print '%s.%s = %s' % \
                  (args[0],args[1],get(args[0],args[1]))
    elif (cmd == "!delete" or cmd == "!del"):
        for arg in args:
            try:
                del values[arg]
    elif (cmd == "!init"):
        init = {}
        if len(args) == 0:
            args = values.keys()
        for arg in args:
            t = toInt(arg,'init')
            if not t:
                values[arg]['init'] = 0
            init[arg] = values[arg]
        order = reversed(sorted(init,key=lambda d: init[d]['init']))
        for i in order:
            print '%s %d' % (i,init[i]['init'])
    elif (cmd == "!roll"):
        out = ''
        for arg in args:
            if 'd' not in arg:
                continue
            spl = arg.split('d')
            if len(spl) != 2:
                continue
            try:
                num = int(spl[0])
                dic = int(spl[1])
            except:
                continue
            dic = dice(num,dic)
            out += str(dic.roll()) + ' '
        if out:
            print out
    elif (cmd == "!hit"):
        dic = dice(3,6) # 3d6
        print hitTable[dic.roll()]
    elif (cmd == "!exit"):
        pass
    else:
        try:
            val = eval(inp)
            if 'print' not in inp:
                print val
        except Exception,e:
            print "%s" % (e)
