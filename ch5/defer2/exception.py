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
        try:
            print(f"f({x+0//x})")  # raises if x == 0
            f(x - 1)
        finally:
            print(f"defer {x}")
            i = 1
    except Exception as e:
        print(f"error instead of Exception: {e}, {i}, {x}")
    else:
        try:
            i = 0 // (i - 1)
        except Exception as e:
            print(f"new error instead of new Exception: {e}, {i}, {x}")


if __name__ == '__main__':
    main()
