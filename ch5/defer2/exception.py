def main():
    try:
        f(3)
        print("2niki2")
    finally:
        print_niki()


def print_niki():
    print("!!!niki!!!")


def f(x):
    try:
        print(f"f({x+0//x})")  # raises if x == 0
        f(x - 1)
    finally:
        print(f"defer {x}")


if __name__ == '__main__':
    main()
