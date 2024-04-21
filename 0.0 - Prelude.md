# Prelude
The idea is to game-ify learning. Rather than develop a game to do this from scratch, I thought modifying other games would serve just as well. So, we're going to be looking at, primarily, RuneScape 3 and Old School RuneScape. Both of these titles lack anticheat and competent developers, so the risk of accounts being banned and the process being slowed is next to nil. Especially with RS3.

## Interacting with the game via Python. How?
You'll be using Lennissa (CB, if you know me like that). I've written a simple API that exposes game functions to Python. You'll be able to call functions that the original developers call as if you had source access (kinda).

Example of this would be something like
```python
class Script(CBScript):
    def __init__(self):
        self.player.walk_to(x, y)
```

> Where'd `self.player` come from?

The bot declares and defines various helpers to make it easier for the would-be script author (you). Any class that inherits from `CBScript` will have access to them. `self.player` being one of many.


# Java -> Python examples
From recollection, most of you are more partial to Java than Python. In Jack's case I know this to be due to the differences in professors at uni. Even so, we'll be using Python. It makes much more sense. If we were to use Java, it requires a large runtime to be mapped to the game's memory. Which doesn't make sense for something that we will be running up to 30,40,50 instances of at once. The JRE (Java Runtime Environment) alone, while optimized, would ask 750MB-1GB of memory per instance. That's shit.

That being said, since you're all (probably) more familiar with Java than you are Python, I'll write some Java examples then translate them to Python for you.


Java, basic **Hello, world**.
```java
package net.ryswick;

public class Main {
    public static void main(String[] args) {
        System.out.println("Hello, world.");
    }
}
```

Python, basic **Hello, world**.
```python
print("Hello, world.")
```

Java requires classes where Python does not. Java also requires strictly defined entry points, where Python, again, does not. An entry point is simply the "main" function.

In Java, it's 
```java
public static void main(String[] args) {    // This bit here being the entry point.
    System.out.println("Hello, world.");    // This bit here being code that's ran in the entry point.
}
```

In C/C++, it's
```c
int main() {
    printf("Hello, world.");
    // Whatever integer is returned from `main` in C-family languages
    // is considered the "exit code". If a program fails, there's a code for the
    // programmer to return. Succeed, same. 
    // This is not required in Java, but it's not absent either.
    return 0;
}
```

Now, I said Python doesn't **require** an entry point, but it does offer the option.
```python
# hello.py
if __name__ == "__main__":
    print("Hello, world.");
```

> But if it's optional, what is it good for?

In short, the entry point in Python will only run if the script that contains that entry point is explicitly ran, like so:

`python hello.py`

Now, consider this scenario...
```python
# hello.py
def addition(i, j):
    return i + j

if __name__ == "__main__":
    print("Hello, world.")
```

```python
# main.py
import hello    # Lets us use the contents of hello.py in main.py

print(hello.addition(2, 6))
```

```
python hello.py
> Hello, world.
python main.py
> 8
```

Because our "Hello, world." is only printed within the entry point of `hello.py`, it doesn't execute if we import `hello.py` in another file. All that happens is we get access to whatever isn't within the [scope](https://en.wikipedia.org/wiki/Scope_(computer_science)) of the entry point.


# A brief refresher on OOP
OOP, or [Object Oriented Programming](https://en.wikipedia.org/wiki/Object-oriented_programming), is a framework of thinking when developing and designing software. It's, in practice, largely shit. OOP itself is not inherently bad by any means. Assuming all rules of OOP are followed, the resulting codebase should be fairly consumable. However, following every rule to a T is not easy, and such a feat is rarely achieved.

Rather than attempting to follow the rules to a T, which never happens in real-world scenarios, I prefer to "tread with caution". I use classes, I sometimes even *like* using classes, but I don't **__default__** to using classes.

Java will have you believe that classes should be the default, the go-to. So much so that classes are forced onto you. 

Most languages don't give a shit.

Classes are good when
- Will have many of the same thing (Think a Player class in a multiplayer game)
- Will have many of the same-ish thing (inheritance)
- that's it? I'm sure there's more but cba

So, let's write a simple Player class in Java, then translate that to Python.

```java
// Player.java
package net.ryswick;

public class Player {
    private String username;
    private float x, y, z;
    private int health;
    private int maxHealth;
    private int movementSpeed;

    public Player(String username)
    {
        this.username = username;
        this.x = 1.0f;
        this.y = 1.0f;
        this.z = 1.0f;
        this.maxHealth = 150;
        this.health = this.maxHealth;
    }

    // In Java, you're meant to "encapsulate" things. 
    // Meaning, roughly: private fields with public accessors/writers (getters/setters).
    // If a property on an object is not meant to be directly subjected to the will of outside code,
    // it will have no setter. Meaning the value can only be read from an instance of the class,
    // but not written to. e.g getHealth() etc.
    // Since there's no `setHealth(amount)`, that means that `health` can only be changed
    // by code within this class. Like `takeDamage(amount)`.
    public int getHealth() {
        return this.health;
    }

    public int getMaxHealth() {
        return this.maxHealth;
    }

    public String getUsername() {
        return this.username;
    }

    public int getMovementSpeed() {
        return this.movementSpeed;
    }
    // - End getters/setters.

    public void takeDamage(int amount) {
        this.health -= amount;
    }

    public void heal(int amount) {
        this.health += amount;
    }

    public boolean isAlive() {
        return this.health > 0;
    }

    public String toString() {
        return this.username + " HP: " + this.health;
    }
}
```

```java
// Main.java
package net.ryswick;

public class Main {
    public static void main(String[] args) {
        Player reyzr = new Player("Reyzr");
        Player peter = new Player("Peter");
        Player michelle = new Player("Michelle");

        // If we didn't include the `toString()` method on Player,
        // this wouldn't print anything that we want, but instead a memory address within the JVM.
        System.out.println(reyzr);
        System.out.println(peter);
        System.out.println(michelle);
        /* stdout
            Reyzr HP: 150
            Peter HP: 150
            Michelle HP: 150
        */

        reyzr.takeDamage(69);
        peter.takeDamage(10);
        michelle.takeDamage(5);

        System.out.println(reyzr);
        System.out.println(peter);
        System.out.println(michelle);
        /* stdout
            Reyzr HP: 81
            Peter HP: 140
            Michelle HP: 145
        */
    }
}
```

Now, the same thing in Python.

```python
# main.py (one file is fine here)
class Player:
    def __init__(self, username):
        # Python doesn't require that we declare fields of a class, we can just assign shit as we please.
        self._username = username
        self._x = 1.0
        self._y = 1.0
        self._z = 1.0
        self._max_health = 150
        self._health = self._max_health
    
    # In Python, methods require "self" to be the first parameter.
    # Interestingly, in x64 assembly, when methods of classes are called,
    # a pointer to the instance of the class the method is being called on is stored
    # in the RCX register, or in other words, is given as the first parameter automatically.
    def take_damage(self, amount):
        self._health = self._health - amount
    
    def heal(self, amount):
        self._health = self._health + amount
    
    # In Python, inclusion of a `__str__` method means you're able to 
    #   reyzr = Player("Reyzr")
    #   print(reyzr)
    #   > Reyzr HP: 150
    def __str__(self):
        return f"{self._username} HP: {self._health}"


if __name__ == "__main__":
    reyzr = Player("Reyzr")
    peter = Player("Peter")
    michelle = Player("Michelle")

    print(reyzr)
    print(peter)
    print(michelle)
    # stdout
    #   Reyzr HP: 150
    #   Peter HP: 150
    #   Michelle HP: 150

    reyzr.take_damage(69)
    peter.take_damage(10)
    michelle.take_damage(5)
    # stdout
    #   Reyzr HP: 81
    #   Peter HP: 140
    #   Michelle HP: 145
```