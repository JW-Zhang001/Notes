"""
该类实例只能创建__slots__中声明的属性，否则报错, 具体作用就是节省内存
"""
from memory_profiler import profile


class Test(object):
    __slots__ = ['a']


def test_01():
    Test.c = 3  # 类属性仍然可以自由添加
    t = Test()
    t.a = 1
    print(t.c)
    # t.b = 2  # AttributeError: 'Test' object has no attribute 'b'


class TestA(object):
    __slots__ = ['a', 'b', 'c']

    def __init__(self, a, b, c):
        self.a = a
        self.b = b
        self.c = c


class TestB(object):
    def __init__(self, a, b, c):
        self.a = a
        self.b = b
        self.c = c


@profile
def test():
    """"""
    temp = [TestA(i, i + 1, i + 2) for i in range(10000)]
    del temp
    temp = [TestB(i, i + 1, i + 2) for i in range(10000)]
    del temp


def test_02():
    test()


if __name__ == '__main__':
    # test_01()
    test_02()
