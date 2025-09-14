class Person:
    def __init__(self):
        self.name = 'eli'
        self.age = 4


    def run(self):
        print("Hello I'm a human and I'm running double")

if __name__ == "__main__":
    person = Person()
    print("Hello world!")
    person.run()