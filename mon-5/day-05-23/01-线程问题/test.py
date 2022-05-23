"""
# 很久没写 python了，都忘记怎么写了~
print( 'Hello world!' )
print( 'is me' )
"""

# from json.tool import main
import _thread
import threading
import time

a = 1
b = 1


def func_for():
    for i in range(1, 10000000):
        global a, b
        a += 1
        b += 1

        #print("func run  ", i)


def nums():
    for i in range(1,10000000):
        #print('ok')

        if a < b:
            print("a = ", a, "\tb = ", b)


def main():
    print("begin")
    thread1 = threading.Thread(target=func_for, name='T1')
    thread2 = threading.Thread(target=nums, name='T2')

    thread2.start()
    thread1.start()

    print("end")


if __name__ == "__main__":
    main()
