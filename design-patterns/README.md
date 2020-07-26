https://github.com/svett/golang-design-patterns

Structural Patterns

Decorator: Add responsibilities to objects dynamically.  Attach additional responsibilities to an object dynamically. Decorators provide a             flexible alternative to subclassing for extending functionality.
Facade: A single class that represents an entire subsystem. Provide a unified interface to a set of interfaces in a system. Facade defines a higher-level interface that makes the subsystem easier to use.
Flyweight: A fine-grained instance used for efficient sharing. Use sharing to support large numbers of fine-grained objects efficiently. A flyweight is a shared object that can be used in multiple contexts simultaneously. The flyweight acts as an independent object in each context — it’s indistinguishable from an instance of the object that’s not shared.
Proxy: An object representing another object. Provide a surrogate or placeholder for another object to control access to it.


Behavioral Patterns
Chain of Responsibility : A way of passing a request between a chain of objects. Avoid coupling the sender of a request to its receiver by giving more than one object a  chance to handle the request. Chain the receiving objects and pass the request along the chain until an object handles it.
Command: Encapsulate a command request as an object. Encapsulate a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.
Interpreter: A way to include language elements in a program. Given a language, define a representation for its grammar along with an interpreter that uses the representation to interpret sentences in the language.
Iterator: Sequentially access the elements of a collection. Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
Mediator: Defines simplified communication between classes. Define an object that encapsulates how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each other explicitly, and it lets you vary their interaction independently.
Memento: Capture and restore an object's internal state. Without violating encapsulation, capture and externalize an object’s internal state so that the object can be restored to this state later.
Observer: A way of notifying change to a number of classes. Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
State: Alter an object's behavior when its state changes. Allow an object to alter its behavior when its internal state changes. The object will appear to change its class.
Strategy: Encapsulates an algorithm inside a class. Define a family of algorithms, encapsulate each one, and make them interchangeable.            Strategy lets the algorithm vary independently from clients that use it.
Template: Defer the exact steps of an algorithm to a subclass. Define the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm’s structure.
Visitor: Defines a new operation to a class without change. Represent an operation to be performed on the elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates.
I must also mention and credit goes to http://www.dofactory.com/Patterns/Patterns.aspx for the most of the above information.