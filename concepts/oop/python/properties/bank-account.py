class BankAccount:

    def __init__(self, balance=0):
        self._balance = balance

    @property
    def balance(self):
        return self._balance
    
    @balance.setter
    def balance(self, value):
        if value < 0:
            print("Balance can't be negative.")
        else:
            self._balance = value
    
    def deposit(self, ammount):
        new_balance = self.balance + ammount
        self.balance = new_balance

    def withdraw(self, ammount):
        new_balance = self.balance - ammount
        self.balance = new_balance


if __name__ == "__main__":
    account = BankAccount()
    account.deposit(100)
    print(account.balance)
    account.withdraw(120)
    print(account.balance)
