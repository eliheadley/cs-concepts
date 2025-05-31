class Circle:
    pi = 3.14159

    def __init__(self, radius:float=1.0):
        self._radius = radius

    def calculate_area(self) -> float:
        return Circle.pi * self._radius * self._radius

    @classmethod
    def set_pi(cls, value:float) -> None:
        cls.pi = value

    @staticmethod
    def is_valid_radius(value:float) -> bool:
        return True if value >= 0 else False
        


if __name__ == "__main__":
    # Input width and height from the user
    radius = -1
    while not Circle.is_valid_radius(radius):
        radius = float(input("Enter the radius of the Circle (decimal): ")) 

    my_circle = Circle(radius)
    print(f"This is a more exact area: {my_circle.calculate_area()}")
    Circle.set_pi(3.14)
    print(f"This is a less exact area: {my_circle.calculate_area()}")